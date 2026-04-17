package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ikennarichard/contrib-tracker/internal/domain"
	"github.com/ikennarichard/contrib-tracker/internal/service"
)

type ContributionHandler struct {
    service *service.ContributionService
}

func NewContributionHandler(svc *service.ContributionService) *ContributionHandler {
    return &ContributionHandler{service: svc}
}

func (h *ContributionHandler) RegisterRoutes(r chi.Router) {
    r.Route("/api/contributions", func(r chi.Router) {
        r.Post("/", h.Create)
        r.Get("/", h.List)
        r.Get("/{repo}", h.ListByRepo)
    })
}

type createRequest struct {
    Title string `json:"title"`
    Repo  string `json:"repository"`
    Date  string `json:"date"`
    URL   string `json:"url"`
}

func (h *ContributionHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req createRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Basic validation
    if req.Title == "" || req.Repo == "" {
        http.Error(w, "title and repository are required", http.StatusBadRequest)
        return
    }

    err := h.service.Add(req.Title, req.Repo, req.Date, req.URL)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Contribution added successfully",
    })
}

func (h *ContributionHandler) List(w http.ResponseWriter, r *http.Request) {
    repoFilter := r.URL.Query().Get("repository")

    var contribs []*domain.Contribution
    var err error

    if repoFilter != "" {
        contribs, err = h.service.FindByRepo(repoFilter)
    } else {
        contribs, err = h.service.List()
    }

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if contribs == nil {
        contribs = []*domain.Contribution{}
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(contribs); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

func (h *ContributionHandler) ListByRepo(w http.ResponseWriter, r *http.Request) {
    repo := chi.URLParam(r, "repo")
    if repo == "" {
        http.Error(w, "repo is required", http.StatusBadRequest)
        return
    }

    contribs, err := h.service.FindByRepo(repo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if contribs == nil {
        contribs = []*domain.Contribution{}
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(contribs); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}