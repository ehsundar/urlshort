package renderer

import (
	"html/template"
	"io"
)

var (
	create        *template.Template
	createSuccess *template.Template
)

func init() {
	create = template.Must(template.ParseFiles("templates/create.html"))
	createSuccess = template.Must(template.ParseFiles("templates/create_success.html"))
}

type CreateParams struct {
}

type CreateSuccessParams struct {
	ResultURL string
}

func RenderCreate(w io.Writer, params CreateParams) error {
	return create.Execute(w, params)
}

func RenderCreateSuccess(w io.Writer, params CreateSuccessParams) error {
	return createSuccess.Execute(w, params)
}
