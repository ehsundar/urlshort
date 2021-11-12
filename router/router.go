package router

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
)

type registration struct {
	reg *regexp.Regexp
	hnd http.HandlerFunc
}

type Router struct {
	r []registration
}

func (r *Router) Register(pattern string, handler http.HandlerFunc) {
	reg := regexp.MustCompile(pattern)
	r.r = append(r.r, registration{reg: reg, hnd: handler})
}

func (r *Router) Handle(w http.ResponseWriter, req *http.Request) {
	reqWithCtx := req

	for _, reg := range r.r {
		if reg.reg.MatchString(req.RequestURI) {
			if reg.reg.NumSubexp() > 0 {
				ctx := req.Context()
				submatch := reg.reg.FindStringSubmatch(req.RequestURI)[1:]
				for i, name := range reg.reg.SubexpNames()[1:] {
					ctx = context.WithValue(ctx, name, submatch[i])
				}
				reqWithCtx = req.WithContext(ctx)
			}

			reg.hnd(w, reqWithCtx)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("no route matched"))
	if err != nil {
		fmt.Println(err)
	}
}
