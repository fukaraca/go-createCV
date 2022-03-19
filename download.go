package main

import "net/http"

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./dump/cv.pdf")
}
