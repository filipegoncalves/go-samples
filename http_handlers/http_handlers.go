
package main

import (
	"fmt"
	"os"
	"net/http"
)

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

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	if e := http.ListenAndServe("localhost:4000", nil); e != nil {
		fmt.Println("Error on http.ListenAndServe:", e)
		os.Exit(1)
	}
}
