package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Id          string   `json:"_id"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
}

type Uuid struct {
	Ids []string `json:"uuids"`
}

var DMS_HOST string = ""
var DMS_DB string = ""
var DMS_STORAGE string = ""

var http_client http.Client = http.Client{Timeout: time.Duration(5) * time.Second}

func main() {
	if len(os.Args) < 2 {
		print_help()
		os.Exit(1)
	}

	DMS_HOST, _ := os.LookupEnv("DMS_HOST")
	DMS_DB, _ := os.LookupEnv("DMS_DB")
	DMS_STORAGE, _ := os.LookupEnv("DMS_STORAGE")

	if DMS_HOST == "" || DMS_DB == "" || DMS_STORAGE == "" {
		fmt.Println("Environment variables not set.")
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addFilename := addCmd.String("f", "", "file")
	addDescription := addCmd.String("d", "", "description")

	switch os.Args[1] {
	case "filter":
		couchdb_filter()
	case "add":
		addCmd.Parse(os.Args[2:])
		couchdb_add(addFilename, addDescription, addCmd.Args())
	default:
		print_help()
	}
}

func couchdb_filter() {
	// {
	// 		"selector": {
	//			"labels": {
	//				"$all": [
	//					"2023", "invoice"
	//				]
	//			}
	//		}
	// }

}

func couchdb_add(file *string, desc *string, labels []string) {
	uuid := couchdb_get_uuid()

	m := Message{uuid, *desc, labels}
	//todo: check if file exists

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", DMS_HOST+"/"+DMS_DB, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := http_client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status: ", resp.Status)
	//todo: copy file to dms storage
}

func couchdb_get_uuid() string {
	req, err := http.NewRequest("GET", DMS_HOST+"/_uuids", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := http_client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ids := Uuid{}
	if err := json.Unmarshal(body, &ids); err != nil {
		panic(err)
	}

	return ids.Ids[0]
}

func print_help() {

}
