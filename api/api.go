package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(SetApplicationJson)

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post(
		"/api/shorten",
		handlePost(db),
	)

	r.Get(
		"/{code}",
		handleGet(db),
	)
	return r

}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		r.Body = http.MaxBytesReader(w, r.Body, 1000)
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJson(
				w,
				Response{
					Error: "invalid-body",
				},
				http.StatusUnprocessableEntity,
			)
			return
		}
		if _, err := url.Parse(body.URL); err != nil {
			// verificando se a url é válida
			SendJson(
				w,
				Response{
					Error: "invalid-url",
				},
				http.StatusBadRequest,
			)
			return
		}
		code := GenCode()
		db[code] = body.URL
		SendJson(
			w,
			Response{
				Data: code,
			},
			http.StatusCreated,
		)
	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		data, ok := db[code]
		if !ok {
			SendJson(
				w,
				Response{
					Error: "invalid-url",
				},
				http.StatusNotFound,
			)
			return
		}

		http.Redirect(
			w,
			r,
			data,
			http.StatusPermanentRedirect,
		) // enviando usuário para o URL

		SendJson(
			w,
			Response{
				Data: data,
			},
			http.StatusOK,
		)
	}
}
