package handler

import (
	"computerextra/elaerning-go/db"
	"computerextra/elaerning-go/internal/component"
	"computerextra/elaerning-go/internal/service/realip"
	"computerextra/elaerning-go/internal/util/flash"

	"fmt"
	"log/slog"
	"net/http"

	"github.com/lucsky/cuid"
)

type Handler struct {
	logger     *slog.Logger
	db         *db.PrismaClient
	ipResolver *realip.Service
}

func New(logger *slog.Logger, db *db.PrismaClient, ipService *realip.Service) *Handler {
	return &Handler{
		logger:     logger,
		db:         db,
		ipResolver: ipService,
	}
}

// 1MB Maxbodysize
const maxBodySize = 1 << 20

func (h *Handler) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the request body
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	content := r.FormValue("content")
	if content == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ctx := r.Context()
	// ip := h.ipResolver.RealIPForRequest(r)
	id := cuid.New()

	// TODO: Write to DB!
	_, err := h.db.Snippet.CreateOne(
		db.Snippet.Snippet.Set(content),
		db.Snippet.ID.Set(id),
	).Exec(ctx)
	if err != nil {
		h.logger.Error("failed to write snippet", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	uri := fmt.Sprintf("%s://%s/%s", scheme, host, id)

	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) GetSnippet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	// TODO: Get Snippet from DB
	content, err := h.db.Snippet.FindUnique(db.Snippet.ID.Equals(id)).Exec(ctx)
	if err != nil {
		h.logger.Error("failed to get snippet", slog.Any("error", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	uri := fmt.Sprintf("%s://%s/%s", scheme, host, id)

	component.SnippetPage(content.Snippet, uri).Render(ctx, w)
}
