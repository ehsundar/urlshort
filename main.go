package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort/composer"
	"urlshort/storage/mem"
)

func main() {
	strg := mem.NewMemStorage()
	generalComposer := composer.NewMd5Base64()

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		short := ""

		for nonce := 0; ; nonce++ {
			short = generalComposer.Compose(string(body), fmt.Sprintf("%d", nonce))
			err = strg.Create(nil, short, string(body))
			if err == nil {
				break
			}
		}

		_, err = w.Write([]byte(short))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8000", http.DefaultServeMux); err != nil {
		panic(err)
	}
}
