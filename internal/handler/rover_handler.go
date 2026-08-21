package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GerardoHuHer/go_crud/internal/service"
	"github.com/gorilla/mux"
)

type RoverHandler struct {
	svc *service.RoverService
}

func NewRoverHandler(svc *service.RoverService) *RoverHandler {
	return &RoverHandler{svc: svc}
}

type CreateRoverRequest struct {
	Pos_x int `json:"pos_x"`
	Pos_y int `json:"pos_y"`
}

func (h *RoverHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRoverRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	rover, err := h.svc.CreateRover(r.Context(), req.Pos_x, req.Pos_y)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rover)
}

func (h *RoverHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	rover, err := h.svc.GetRoverById(r.Context(), uint(id))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(rover)

}

func (h *RoverHandler) UpdateByIdHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateRoverRequest
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	rover, err := h.svc.UpdateRoverById(r.Context(), uint(id), req.Pos_x, req.Pos_y)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(rover)

}

func (h *RoverHandler) DeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.svc.DeleteRoverById(r.Context(), uint(id))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"msg": fmt.Sprintf("%d se ha eliminado con éxito", id),
	}
	json.NewEncoder(w).Encode(response)

}

func (h *RoverHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	vehiculos, err := h.svc.GetAll(r.Context())
	if err != nil {
		http.Error(w, "There was an error", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(vehiculos)
}
