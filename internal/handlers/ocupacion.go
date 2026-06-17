package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto_movilidad_fcvt/internal/modelos"
	"proyecto_movilidad_fcvt/internal/storage"

	"github.com/go-chi/chi/v5"
)

type OcupacionHandler struct {
	repo storage.OcupacionesRepository
}

func NewOcupacionHandler(repo storage.OcupacionesRepository) *OcupacionHandler {
	return &OcupacionHandler{
		repo: repo,
	}
}

// GET /api/ocupaciones
func (h *OcupacionHandler) Listar(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.repo.ListarOcupaciones())
}

// GET /api/ocupaciones/{id}
func (h *OcupacionHandler) Obtener(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	ocupacion, ok := h.repo.BuscarOcupacionPorID(id)
	if !ok {
		http.Error(w, "Ocupación no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ocupacion)
}

// POST /api/ocupaciones
func (h *OcupacionHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var ocupacion modelos.Ocupacion

	if err := json.NewDecoder(r.Body).Decode(&ocupacion); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	nueva := h.repo.CrearOcupacion(ocupacion)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nueva)
}

// PUT /api/ocupaciones/{id}
func (h *OcupacionHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var datos modelos.Ocupacion

	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	actualizada, ok := h.repo.ActualizarOcupacion(id, datos)
	if !ok {
		http.Error(w, "Ocupación no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(actualizada)
}

// DELETE /api/ocupaciones/{id}
func (h *OcupacionHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if !h.repo.BorrarOcupacion(id) {
		http.Error(w, "Ocupación no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PATCH /api/ocupaciones/{id}/liberar
func (h *OcupacionHandler) Liberar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	ocupacion, ok := h.repo.LiberarOcupacion(id)
	if !ok {
		http.Error(w, "Ocupación no encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ocupacion)
}
