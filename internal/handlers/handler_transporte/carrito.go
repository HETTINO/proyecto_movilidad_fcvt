package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	modelos "proyecto_movilidad_fcvt/internal/modelos"
)

// ListarCarritos atiende GET /api/v1/carritos
func (s *Server) ListarCarritos(w http.ResponseWriter, _ *http.Request) {
	carritos := s.Carrito.Listar()
	responderJSON(w, http.StatusOK, carritos)
}

// ObtenerCarrito atiende GET /api/v1/carritos/{id}
func (s *Server) ObtenerCarrito(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	carrito, encontrado := s.Carrito.Obtener(id)
	if !encontrado {
		responderError(w, http.StatusNotFound, "carrito no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, carrito)
}

// CrearCarrito atiende POST /api/v1/carritos
func (s *Server) CrearCarrito(w http.ResponseWriter, r *http.Request) {
	var nuevo modelos.Carrito

	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	creado, err := s.Carrito.Crear(nuevo)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, creado)
}

// ActualizarCarrito atiende PUT /api/v1/carritos/{id}
func (s *Server) ActualizarCarrito(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	var datos modelos.Carrito
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	actualizado, encontrado, err := s.Carrito.Actualizar(id, datos)
	if err != nil {
		responderError(w, statusDeError(err), err.Error())
		return
	}
	if !encontrado {
		responderError(w, http.StatusNotFound, "carrito no encontrado")
		return
	}

	responderJSON(w, http.StatusOK, actualizado)
}

// BorrarCarrito atiende DELETE /api/v1/carritos/{id}
func (s *Server) BorrarCarrito(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responderError(w, http.StatusBadRequest, "id debe ser un número entero")
		return
	}

	if err := s.Carrito.Borrar(id); err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
