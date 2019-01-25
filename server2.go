package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Text(re *http.Response) string {
	body, _ := ioutil.ReadAll(re.Body)
	fmt.Println(body)
	return string(body)
}

func main() {

	session := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	response, err := http.Get("http://localhost:3754/date")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	post_data := []byte(`{"text": 5}`)

	req, err := http.NewRequest("POST", "http://localhost:3754/postdate", bytes.NewBuffer(post_data))
	if err != nil {
		return
	}
	r, err := session.Do(req)
	if err != nil {
		return
	}

	fmt.Println(Text(r))
	i, err := strconv.ParseFloat(Text(r), 64)
	if err == nil {
		fmt.Println("a =", i)
	}
	fmt.Println(i)
	fmt.Println(i)
}
