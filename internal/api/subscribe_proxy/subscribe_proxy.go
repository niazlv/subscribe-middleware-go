package subscribeproxy

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niazlv/subscribe-middleware-go/internal/database"
)

func Setup(app *gin.RouterGroup) {
	api := app.Group("noncesub")
	api.GET(":id", getSubscribe)
}

func getSubscribe(c *gin.Context) {

	db := database.InitDB()

	defer db.Close()
	// get id and find it on DB
	db_subscribe, err := database.ReadSubscribeByID(db, c.Param("id"))
	if err != nil {
		if err == sql.ErrNoRows {
			// http.StatusBadRequest was changed to http.StatusNotFound because nginx redirect 404 error to original server
			c.Data(http.StatusNotFound, "text/plain; charset=utf-8",
				[]byte("Error!"))
			return
		}
		// TODO: implement handle "error"
		c.Data(http.StatusInternalServerError, "text/plain; charset=utf-8", []byte(""))
		log.Panic(err)
		return
	}

	encodedCombini, err := MergeSubscribes(db, db_subscribe)
	if err != nil {
		// TODO: implement handle "error"
		c.Data(http.StatusInternalServerError, "text/plain; charset=utf-8", []byte(""))
		log.Panic(err)
		return
	}

	c.Data(http.StatusOK,
		"text/plain; charset=utf-8",
		[]byte(encodedCombini))
}

func MergeSubscribe(subscribe1 string, subscribe2 string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(string(subscribe1))
	if err != nil {
		return "", err
	}

	var str1 = string(data)

	data, err = base64.StdEncoding.DecodeString(string(subscribe2))
	if err != nil {
		return "", err
	}

	var str2 = string(data)

	combini := str1 + str2

	encodedCombini := base64.StdEncoding.EncodeToString([]byte(combini))
	return encodedCombini, nil
}

func externalGetHttp(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func MergeSubscribes(db *sql.DB, sub database.Subscibe, lastMerged ...string) (string, error) {
	// crutch
	if len(lastMerged) == 0 {
		lastMerged = append(lastMerged, "")
	} else if len(lastMerged) > 1 {
		return "", fmt.Errorf("ожидается только один параметр lastMerged")
	}
	var data1 string
	var err error
	data1, err = "", nil
	if sub.Subscribe1 != "" {
		data1, err = externalGetHttp(sub.Subscribe1)
		if err != nil {
			return "", err
		}
	}

	var data2 string
	data2, err = "", nil
	if sub.Subscribe2 != "" {
		data2, err = externalGetHttp(sub.Subscribe2)
		if err != nil {
			return "", err
		}
	}

	merged, err := MergeSubscribe(data1, data2)
	if err != nil {
		return merged, err
	}

	if len(lastMerged) == 1 {
		bufMerged := merged
		merged, err = MergeSubscribe(lastMerged[0], bufMerged)
		if err != nil {
			return merged, err
		}
	}

	if sub.Next != "" {
		sub2, err := database.ReadSubscribeByID(db, sub.Next)
		if err != nil {
			return merged, err
		}
		return MergeSubscribes(db, sub2, merged)
	}
	return merged, nil
}
