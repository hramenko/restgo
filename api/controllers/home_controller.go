package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	_ "net/url"
	"os"
	"strings"
	_ "strings"
	"time"
)

func redirectFunc(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}

	lastReq := via[len(via)-1]
	if req.Response.StatusCode == 301 && lastReq.Method == http.MethodPost {
		req.Method = http.MethodPost

		// Get the body of the original request, set here, since req.Body will be nil if a 302 was returned
		if via[0].GetBody != nil {
			var err error
			req.Body, err = via[0].GetBody()
			if err != nil {
				return err
			}
			req.ContentLength = via[0].ContentLength
		}
	}
	return nil
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	//responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

	//jsonBody := []byte(`{"client_message": "hello, server!"}`)
	/*	jsonData := []byte(`{
		"command":"Gelikon",
		"arguments":
		  { "date" : "2022-10-01T00:00:00",
							"data2" : "2021-11-30T00:00:00",
							"TypeRequest":"PVPFinanceResult",
							"PVP":"1"
			}

	}`)*/

	//req.Header.Set("Authorization", "Basic R2VsaUtvbjpVQ2JKVzBuOTJHNDFDNWM=")
	//req.Header.Set("Accept", "application/json")
	//req.Header.Set("Encoding", "application/json")

	//jsonData := []byte(`{"command":"service_pvp.get_pvp", "arguments":{"pvp_id":1000}}`)
	//jsonData := `{"command":"service_pvp.get_pvp", "arguments":{"pvp_id":1000}}`
	//requestUrl := "https://Cons24.tentorium.ru/Trade_11-5/hs/src/src"
	//requestUrl := "http://prk-654/TradeNew2/hs/src/src"
	requestUrl := "http://prk-654/tradenew2/hs/src/src"

	//req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	//req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(jsonData))

	//jsonData := url.QueryEscape(`{"command":"service_pvp.get_pvp", "arguments":{"pvp_id":1000}}`)
	//jsonData := `{"command":"service_pvp.get_pvp", "arguments":{"pvp_id":1000}}`

	//json.NewEncoder(w).Encode(jsonData)
	//strings.NewReader(jsonData)
	//req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer([]byte(jsonData)))

	//payload := strings.NewReader("{\"command\":\"service_pvp.get_pvp\", \"arguments\":{\"pvp_id\":1000}}\r\n")
	payload := strings.NewReader(`{
	"command":"service_pvp.get_pvp",
		"arguments":{
			"pvp_id":1000
	}
}`)
	req, err := http.NewRequest(http.MethodPost, requestUrl, payload)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	//req.Header.Set("Encoding", "application/json")
	req.Header.Set("Authorization", "Basic aHR0cDpodHRw")
	//req.Header.Add("Authorization", "Basic R2VsaUtvbjpVQ2JKVzBuOTJHNDFDNWM=")
	req.Header.Add("Accept-Charset", "utf-8")

	//req.SetBasicAuth("http", "http")
	//req.Body.Read()
	b, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(b))
	//req.
	//aa, aaa := http.Client.Do()
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{
		Timeout:       30 * time.Second,
		CheckRedirect: redirectFunc,
	}
	//response, err := http.DefaultClient.Do(req)
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
	//json.marshal(body, &objs)
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
	//json.NewEncoder(w).Encode(objs[`data`])

	//w.Write([]byte(jsonData))
	json.NewEncoder(w).Encode(objs)
	//json.NewEncoder(w).Encode(jsonData)
	//json.NewEncoder(w).Encode(string(jsonData))

	//responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
