package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	NAME string    `json:"name"`
}

type Server struct {
	*mux.Router
	shoppingItems []Item
}

func NewServer() *Server {
	server := &Server{
		Router:        mux.newRouter(),
		shoppingItems: []Items{},
	}
	server.routes()
	return server
}

func (s *Server) routes() {
	s.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("GET")
	s.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("POST")
	s.HandleFunc("/shopping-items/{id}", s.listShoppingItems()).Methods("DELETE")
}

func (s *Server) createShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.Decoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.ID = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingItems {
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}
