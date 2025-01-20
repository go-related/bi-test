package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		h.registerUserRoutes(r)
		h.registerGroupRoutes(r)
	})

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("ent demo application"))
	})
}

func (h *Handler) registerUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", h.GetUserById)
			r.Route("/friends", func(r chi.Router) {
				r.Get("/", h.GetUserFriends)
				r.Put("/", h.UpdateUserFriends)
			})
		})

		r.Get("/", h.ListUsers)
		r.Post("/", h.CreateUser)
		r.Put("/{userID}", h.UpdateUser)
		r.Delete("/{userID}", h.DeleteUser)

	})

}

func (h *Handler) registerGroupRoutes(r chi.Router) {
	r.Route("/groups", func(r chi.Router) {
		//r.Get("/{groupID}", h.Update)
		//r.Get("/", h.List)
		//r.Post("/", h.Create)
		//r.Put("/{groupID}", h.Update)
		//r.Delete("/{groupID}", h.Delete)
	})
}
