package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"traefik/service/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gojektech/heimdall/httpclient"
)

func main() {
	itemArr := db.GetDb()

	var url string
	var elasticHostname string = "10.8.13.61:9200"
	var elasticsearchIndexName string = "products"

	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	for i := 0; i < len(itemArr); i++ {
		payloadByte, _ := json.Marshal(itemArr[i])
		payload := bytes.NewReader(payloadByte)

		ID := itemArr[i].ID
		IDstring := strconv.FormatInt(ID, 10)
		url = "http://" + elasticHostname + "/" + elasticsearchIndexName + "/_doc/" + IDstring

		res, _ := client.Put(url, payload, headers)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}

}
