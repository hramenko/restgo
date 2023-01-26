package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	_ "strings"
	"time"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	//responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

	//jsonBody := []byte(`{"client_message": "hello, server!"}`)
	jsonData := []byte(`{
		"command":"Gelikon",
		"arguments":
		  { "date" : "2022-10-01T00:00:00",
							"data2" : "2021-11-30T00:00:00",
							"TypeRequest":"PVPFinanceResult",
							"PVP":"1" 
			}
	
	}`)

	requestUrl := "https://Cons24.tentorium.ru/Trade_11-5/hs/src/src"

	//req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Encoding", "application/json")
	req.Header.Set("Authorization", "Basic R2VsaUtvbjpVQ2JKVzBuOTJHNDFDNWM=")

	client := &http.Client{Timeout: 30 * time.Second}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	//json.NewEncoder(w).Encode(data)
	//	fmt.Println("response Status:", response.Status)
	//	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	//	fmt.Println("response Body:", string(body))
	//strdata := string(body)
	//	var objs []map[string]*json.RawMessage

	// Create a map to hold the JSON data
	//var objs map[string]interface{}
	var objs map[string]interface{}
	// Decode the JSON data into the map
	err = json.NewDecoder(response.Body).Decode(&objs)

	//json.Unmarshal(body, &objs)
	json.Unmarshal(body, &objs)

	//	for i := range objs[`data`] {
	//		objs[i][`PVP`] += 1
	//	}
	//data := objs[`data`]
	//	for rec := range objs[`data`] {
	//		fmt.Println(rec)
	//	}

	//responses.JSON(w, http.StatusOK, string(body))
	//jsdata := json.NewDecoder(body)
	//js := json.NewEncoder(w).Encode(string(body))
	//fmt.Println("response Body:", js)
	//	fmt.Println(string(body))

	//w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	//responses.JSON(w, http.StatusOK, objs)
	//fmt.Println(`{afasdf: sdfsdf}`)
	//fmt.Println(objs)
	json.NewEncoder(w).Encode(objs[`data`])

	//responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
