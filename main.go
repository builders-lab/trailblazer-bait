package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", 500)
		return
	}
	defer r.Body.Close()

	fmt.Println("=== Incoming POST ===")
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Headers:", r.Header)
	fmt.Println("Body:", string(body))
	fmt.Println("=====================")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on :8081")
	http.ListenAndServe(":8081", nil)
}
