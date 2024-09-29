package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func GetServerInfo(serverName string) MCServerInfo {
	response, err := http.Get("http://api.mcsrvstat.us/3/" + serverName)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject MCServerInfo
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
