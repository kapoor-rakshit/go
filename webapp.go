package main

import(
	"fmt"
	"net/http"                                     // Package http provides HTTP client and server implementations.
)
                                                  /* w parameter is of the http.ResponseWriter type , to write to respone page, 
                                                     r is *http.Request type , we're reading through the value of the request */
func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Home page")                 /* Fprintf formats by what is specified, and writes to the first parameter, w */
}

func result(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Results page")
}


func main() {
	http.HandleFunc("/", home)                       /* handles functions that will coorespond to paths */
	http.HandleFunc("/res/", result)

	http.ListenAndServe(":5000",nil)                /* listen on port 5000, and the nil (param) is for server stuff for now */
}