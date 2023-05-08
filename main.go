// Julius Shol
// Project
// My Codes.
// In the video
package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	 type Cookie struct{
		Name 	string
		Value 	string

		Path 	string 		- indicates a URL path that must exist in req URL inorder to send Cookie Header
		Domain 	string  	- specifies which hosts are allowed to recieve the Cookie
		Expires	time.Time 	- deletes at specified data
		RawExpires string   // for reading cookies only


		// MaxAge=0 means no 'Max-Age' attribute specified
		// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
		// MaxAge>0 means Max-Age attribute present and given in seconds
		MaxAge		int 		-deletes cookie after specified amount of the time in seconds
		Secure 		bool 		- sent to the server on encrypted req, never unsecured (HTTP)
		HttpOnly	bool		- not accessible by Javascript, only sent to server
		SameSite	SameSite 	- servers require that a cookie shouldn't be sent with cross-origin req
								- attribute lets servers specify whether/when cookies are sent with cross-site requests
		Raw 		string
		Unparsed	[]string 	- raw text of unparsed attribute-value pairs
	}
*/
func main() {

	//start a new web server with the two endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	log.Print("Listening")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// looks for cookie on our machine
	// func (r *request) Cookie(name string) (*Cookie, error)
	cookie, err := r.Cookie("1st-cookie")
	fmt.Println("cookie:", cookie, "err:", err)

	if err != nil {
		// no cookie so we make our cookie
		fmt.Println("cookie was not found")
		cookie = &http.Cookie{
			Name:     "1st-cookie",
			Value:    "my first cookie value",
			HttpOnly: true,
		}
		// func SetCookie(w ResponseWriter, cookie *Cookie)
		http.SetCookie(w, cookie)
	}
	w.Write([]byte("Dealing with Cookies :)"))
}
