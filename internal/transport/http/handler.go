package http

import (
	"fmt"
	"net/http"

	"github.com/ecam900/go-rest/internal/comment"
	"github.com/gorilla/mux"
)

// Hanlder - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all the routes for our application.
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "I am alive")
	})
}
