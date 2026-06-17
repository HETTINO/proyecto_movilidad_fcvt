package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage"

	"github.com/go-chi/chi/v5"
)

type EspacioHandler struct {
	repo storage.EspacioRepository
}

func NewEspacioHandler(repo storage.EspacioRepository) *EspacioHandler {
	return &EspacioHandler{
		repo: repo,
	}
}

// GET /api/espacios
func (h *EspacioHandler) Listar(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.repo.ListarEspacios())
}

// GET /api/espacios/{id}
func (h *EspacioHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	espacio, ok := h.repo.BuscarEspacioPorID(id)
	if !ok {
		http.Error(w, "Espacio no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(espacio)
}

// POST /api/espacios
func (h *EspacioHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var espacio modelos.Espacio

	if err := json.NewDecoder(r.Body).Decode(&espacio); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	nuevo := h.repo.CrearEspacio(espacio)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

// PUT /api/espacios/{id}
func (h *EspacioHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Espacio

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizado, ok := h.repo.ActualizarEspacio(id, datos)
	if !ok {
		http.Error(w, "Espacio no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actualizado)
}

// DELETE /api/espacios/{id}
func (h *EspacioHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.repo.BorrarEspacio(id) {
		http.Error(w, "Espacio no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
