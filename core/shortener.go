package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort/composer"
	"urlshort/storage"
)

type Shortener struct {
	storage  storage.URLStorage
	composer composer.Composer
}

func NewShortener(s storage.URLStorage, c composer.Composer) *Shortener {
	return &Shortener{
		storage:  s,
		composer: c,
	}
}

func (s *Shortener) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		template, err := ioutil.ReadFile("templates/create.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(template)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil || !r.Form.Has("long") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		long := r.Form.Get("long")
		short := ""

		for nonce := 0; ; nonce++ {
			short = s.composer.Compose(long, fmt.Sprintf("%d", nonce))
			if err := s.storage.Create(r.Context(), short, long); err == nil {
				break
			}
		}

		_, err = w.Write([]byte(short))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
