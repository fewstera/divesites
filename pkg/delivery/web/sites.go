package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fewstera/divesites/pkg/commandhandler"
	"github.com/fewstera/divesites/pkg/site"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type sitesHandler struct {
	ch       *commandhandler.CommandHandler
	siteRepo *site.Repository
}

type CreateSiteDTO struct {
	Name     string  "json:name"
	Location string  "json:location"
	Depth    float32 "json:depth"
}

func sitesRouter(ch *commandhandler.CommandHandler, siteRepo *site.Repository) chi.Router {
	h := &sitesHandler{
		ch:       ch,
		siteRepo: siteRepo,
	}

	r := chi.NewRouter()
	r.Post("/", h.create)
	r.Get("/{id}", h.getByID)

	return r
}

func (h *sitesHandler) getByID(w http.ResponseWriter, r *http.Request) {
	siteID := chi.URLParam(r, "id")
	_, err := h.siteRepo.Get(siteID)
	if err != nil {
		errorResponse(w, r, fmt.Errorf("retrieving site: %s", err))
		return
	}
}

func (h *sitesHandler) create(w http.ResponseWriter, r *http.Request) {
	var dto CreateSiteDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		errorResponse(w, r, fmt.Errorf("parsing request body: %s", err))
		return
	}

	siteID := uuid.New().String()
	err = h.ch.Handle(&site.CreateCommand{
		ID:       siteID,
		Name:     dto.Name,
		Location: dto.Location,
		Depth:    dto.Depth,
	})
	if err != nil {
		errorResponse(w, r, fmt.Errorf("creating site: %s", err))
		return
	}

	site, err := h.siteRepo.Get(siteID)
	if err != nil {
		errorResponse(w, r, fmt.Errorf("retrieving site after creation: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(site)
}
