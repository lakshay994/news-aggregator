package handlers

import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200 Ok")
}
