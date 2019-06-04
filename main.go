package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")

	// read the environment if we can
	b, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace") // just pass the file name
	if err != nil {
		fmt.Print(w, err)
	}
	str := string(b)

	fmt.Fprintf(w, "The environent should be:\n")
	fmt.Fprintf(w, str)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
