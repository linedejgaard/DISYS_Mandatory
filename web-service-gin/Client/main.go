package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	ID            string  `json:"id"`
	Studentrating float64 `json:"studentrating"`
	Name          string  `json:"name"`
	Teacher       string  `json:"teacher"`
	Ects          float64 `json:"ects"`
}

func main() {
	GETAllCourses()
	POSTRANDOMCourse()
	GETBPAKCourses()
	DELETEBPAKCourse()
	GETBPAKCourses()
	GETAllCourses()
}

func GETAllCourses() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl http://localhost:8080/courses

	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	printOutput(resp)
}

func GETBPAKCourses() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl http://localhost:8080/courses

	resp, err := http.Get("http://localhost:8080/courses/2")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	printOutput(resp)
}

func POSTRANDOMCourse() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl http://localhost:8080/courses --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","studentrating": 4.2,"name": "BPAK","teacher": "henriette", "ects": 7.5}'
	//

	data := Payload{
		ID:            "10",
		Studentrating: 4.2,
		Name:          "RAND",
		Teacher:       "Randomette",
		Ects:          7.5,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/courses", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}

func DELETEBPAKCourse(){
// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

req, err := http.NewRequest("DELETE", "http://localhost:8080/courses/2", nil)
if err != nil {
	// handle err
}
req.Header.Set("Accept", "application/json")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()
}

func printOutput(resp *http.Response) {
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}