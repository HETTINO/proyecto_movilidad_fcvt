package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage"

	"github.com/go-chi/chi/v5"
)

type ParqueaderoHandler struct {
	repo storage.ParqueaderoRepository
}

func NewParqueaderoHandler(repo storage.ParqueaderoRepository) *ParqueaderoHandler {
	return &ParqueaderoHandler{
		repo: repo,
	}
}

// GET /api/parqueaderos
func (h *ParqueaderoHandler) Listar(w http.ResponseWriter, r *http.Request) {
	parqueaderos := h.repo.ListarParqueaderos()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parqueaderos)
}

// GET /api/parqueaderos/{id}
func (h *ParqueaderoHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	parqueadero, ok := h.repo.BuscarParqueaderoPorID(id)
	if !ok {
		http.Error(w, "Parqueadero no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parqueadero)
}

// POST /api/parqueaderos
func (h *ParqueaderoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var parqueadero modelos.Parqueadero

	if err := json.NewDecoder(r.Body).Decode(&parqueadero); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	nuevo := h.repo.CrearParqueadero(parqueadero)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(nuevo)
}

// PUT /api/parqueaderos/{id}
func (h *ParqueaderoHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Parqueadero

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok := h.repo.ActualizarParqueadero(id, datos)
	if !ok {
		http.Error(w, "Parqueadero no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actualizado)
}

// DELETE /api/parqueaderos/{id}
func (h *ParqueaderoHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.repo.BorrarParqueadero(id) {
		http.Error(w, "Parqueadero no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
