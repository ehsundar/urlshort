package renderer

import (
	"html/template"
	"io"
	"urlshort/storage"
)

var (
	create        *template.Template
	createSuccess *template.Template
	list          *template.Template
)

func init() {
	create = template.Must(template.ParseFiles("templates/create.html"))
	createSuccess = template.Must(template.ParseFiles("templates/create_success.html"))
	list = template.Must(template.ParseFiles("templates/list.html"))
}

type CreateParams struct {
}

type CreateSuccessParams struct {
	ResultURL string
}

type ListParams struct {
	Items []storage.Item
}

func RenderCreate(w io.Writer, params CreateParams) error {
	return create.Execute(w, params)
}

func RenderCreateSuccess(w io.Writer, params CreateSuccessParams) error {
	return createSuccess.Execute(w, params)
}

func RenderList(w io.Writer, params ListParams) error {
	return list.Execute(w, params)
}
