package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		parsedURL := regexp.MustCompile(`^\/(.+:\/)?\/*(.*)`).FindStringSubmatch(r.URL.Path)
		if parsedURL[1] != "" {
			parsedURL[2] = parsedURL[1] + "/" + parsedURL[2]
		}
		resp, err := http.Get(parsedURL[2])
		if err != nil {
			fmt.Fprint(w, err)
		} else {
			defer resp.Body.Close()
			responseData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprint(w, err)
			} else {
				fmt.Fprint(w, string(responseData))
			}
		}
	})
	http.ListenAndServe(":80", nil)
}
