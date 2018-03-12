package main

import(
	"fmt"
	"net/http"                                     // Package http provides HTTP client and server implementations.
)

var naam string

                                                  /* w parameter is of the http.ResponseWriter type , to write to respone page, 
                                                     r is *http.Request type , we're reading through the value of the request */
func home(w http.ResponseWriter, r *http.Request){

	                                            /* Fprintf formats by what is specified, and writes to the first parameter, w */

	fmt.Fprintf(w, "<h1>Welcome to Golang world</h1><p>or Go world</p>")    /* HTML code passed as param to w */

	fmt.Fprintf(w, "My %s is %s.<br>" , "<b>Naam</b>" , naam)   /* Variables passed using %s , as HTML code or value */

	fmt.Fprintf(w, `        
		            <h1>Ready to move %s?</h1>
		            <p>It's a WW %s</p>
		            `,naam, "(Wonderful World)")            /* ` code here ` allows a multi-line code to be written in script */

	fmt.Fprintf(w, "<br>Home page")                           // passed a string
}

func result(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Results page")
}


func main() {

	http.HandleFunc("/", home)                       /* handles functions that will coorespond to paths */
	http.HandleFunc("/res/", result)

	fmt.Scanf("%s",&naam)

	http.ListenAndServe(":5000",nil)                /* listen on port 5000, and the nil (param) is for server stuff for now */
}