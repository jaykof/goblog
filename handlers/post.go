package handlers

import (
	"io"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}
