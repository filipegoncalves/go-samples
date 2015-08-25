
package main

import (
	"fmt"
	"os"
	"net/http"
)

type MyServer struct{}

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s Struct) String() string {
	return fmt.Sprintf("Greeting: \"%v\"; Punct: \"%v\"; Who: \"%v\"",
		s.Greeting, s.Punct, s.Who)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", s)
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const index = "<h1>It works!</h1>" +
		"(Blatantly stolen from the Apache default page)<br /><br />" +
		"<p>See <a href=\"/string\">/string</a> to invoke the String HTTP Handler</p>" +
		"<p>See <a href=\"/struct\">/struct</a> to invoke the Struct HTTP Handler</p>"
	fmt.Fprint(w, index)
}

func main() {
	var server MyServer
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.Handle("/", server)
	if e := http.ListenAndServe("localhost:4000", nil); e != nil {
		fmt.Println("Error on http.ListenAndServe:", e)
		os.Exit(1)
	}
}
