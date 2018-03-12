package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	
	url := "http://freegeoip.net/xml/"                  // url to parse data from

	response, _ := http.Get(url)                        // return the response (response) and any error (_)

	fmt.Println(response.Body)                          // response status i.e. 404 (not found), 200 (OK) .... etc.

	bytesdata, _ := ioutil.ReadAll(response.Body)       // content of response which is in byte form, read using ioutil

	stringdata := string(bytesdata)                    // typecasted to string

	fmt.Println(stringdata)

	response.Body.Close()                            // close response.Body to make sure we free up the resources
}