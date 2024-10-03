package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetServerInfo(serverName string) MCServerInfo {
	response, err := http.Get("http://api.mcsrvstat.us/3/" + serverName)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject MCServerInfo
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
