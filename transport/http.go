package transport

import (
	"fmt"
	"net/http"
	"urlshort/core/composer"
	"urlshort/storage"
	"urlshort/templates/renderer"
)

type Shortener struct {
	storage  storage.Storage
	composer composer.Composer
}

func NewHTTPShortener(s storage.Storage, c composer.Composer) *Shortener {
	return &Shortener{
		storage:  s,
		composer: c,
	}
}

func (s *Shortener) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := renderer.RenderCreate(w, renderer.CreateParams{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
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

		params := renderer.CreateSuccessParams{ResultURL: "http://localhost:8000/" + short}
		err = renderer.RenderCreateSuccess(w, params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Shortener) Open(w http.ResponseWriter, r *http.Request) {
	long, err := s.storage.GetLong(r.Context(), r.RequestURI[1:])
	if err == storage.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, long, http.StatusPermanentRedirect)
	return
}
