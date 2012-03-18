package handlers

import (
	"fmt"
	"net/http"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Foo handler!")
}

func BarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Bar handler!!")
}
