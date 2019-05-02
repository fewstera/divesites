package web

import (
	"encoding/json"
	"net/http"

	"github.com/fewstera/divesites/pkg/commandhandler"
	"github.com/fewstera/divesites/pkg/site"
	"github.com/go-chi/chi"
)

func StartServer(addr string, ch *commandhandler.CommandHandler, siteRepo *site.Repository) error {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/sites", sitesRouter(ch, siteRepo))
	})

	return http.ListenAndServe(addr, r)
}

func errorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
