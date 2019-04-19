package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://127.0.0.1:81/api/leads"

	payload := strings.NewReader("{\n\t\"search\":\"pid:1;lead_platform:;lead_appid:\",\n\t\"page_no\":1,\n\t\"page_size\":10,\n\t\"order\":\"asc\",\n\t\"sort\":\"id\"\n}")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJodHRwOlwvXC8xMjcuMC4wLjE6ODFcL2FwaVwvYXV0aFwvdG9rZW4iLCJpYXQiOjE1NTQ5NjM4NTgsImV4cCI6MTU1NDk3MTA1OCwibmJmIjoxNTU0OTYzODU4LCJqdGkiOiIzWVJaY0RhRXpQS2lWRmxEIiwic3ViIjo1MDc1LCJwcnYiOiI4YzJmZjY2ZjRjYjJlMTM4Zjc2ZjJhMzAxMDcxYjZjMDVjMmFhZjQ0IiwicHJvIjoibWthIiwidGVsIjoiMTEwIiwid3JkIjoiVXNpbmcgYWNjZXNzLXRva2VuIG1hbGljaW91c2x5IHdpbGwgYmUgaW4gamFpbC4ifQ.w02S8daiJzSrVsQKkaN4i-p9HjyLVXRQbTWIdcb_QlykFinNoUeObjhUmjgeCXvlPKpLhdyiyDVp3aWOx7tWhw")
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "319989c2-3a8c-4b1e-999f-7821a02af274")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))
	//fmt.Println(body)

}