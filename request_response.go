package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"                                      // library for parsing data
	"encoding/json"
)

type sitemapindex struct {
	Sites[] string `xml:"sitemap>loc"`               // NOTE : variables must be NON-LOWER CASE 
	             // NOTE : Because Unmarshal uses the reflect package, it can only assign to exported (UPPER CASE) fields       
	            // use ">" (without any space) to access tags within tags 
}

type geoip struct {
	Ip string `json:"ip"`
	Country string `json:"country_name"`
}

func main() {
	
	url := "https://www.washingtonpost.com/news-sitemap-index.xml"                  // url to parse data from

	response, _ := http.Get(url)                        // return the response (response) and any error (_)

	fmt.Println(response.Body)                          // response status i.e. 404 (not found), 200 (OK) .... etc.

	bytesdata, _ := ioutil.ReadAll(response.Body)       // content of response which is in byte form, read using ioutil

	stringdata := string(bytesdata)                    // typecasted to string

	fmt.Println(stringdata)

	var s = new(sitemapindex)                         // var for struct

	err := xml.Unmarshal(bytesdata, &s)         // Unmarshal parses the XML-encoded data and stores the result in the value pointed to by "s", 
	                                           // which must be an arbitrary struct, slice, or string. 
	                                           // Well-formed data that does not fit into "s" is discarded.
	if err != nil {
		fmt.Println("Something went wrong",err)
		return
	}

	fmt.Println(len(s.Sites))
	for _ , value := range s.Sites {
		fmt.Println(value)
	}
                                               
                                               /* Parsing JSON data */
	                                          /* More on :  http://blog.josephmisiti.com/parsing-json-responses-in-golang */
	url = "http://freegeoip.net/json"
	response , _ = http.Get(url)
	bytesdata , _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(bytesdata))

	var g geoip
	err = json.Unmarshal(bytesdata,&g)
	if err != nil {
		fmt.Println("Something went wrong",err)
		return
	}

	fmt.Println(g.Ip,g.Country)

	response.Body.Close()                            // close response.Body to make sure we free up the resources
}