package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Handler struct {
	Pattern string
	Page    string
	Data    interface{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join("templates", h.Page)
	templates := template.Must(template.ParseFiles(p))

	if err := templates.ExecuteTemplate(w, h.Page, h.Data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
