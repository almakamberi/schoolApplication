package handlers

import (
	"encoding/json"
	"net/http"
	"schoolApplication/models"
	"schoolApplication/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type ClassHandler struct {
	Repo repository.ClassRepository
}

func NewClassHandler(repo repository.ClassRepository) *ClassHandler {
	return &ClassHandler{Repo: repo}
}

func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&class); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := h.Repo.Create(&class); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, class)
}

func (h *ClassHandler) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.Repo.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, classes)
}

func (h *ClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid class ID")
		return
	}

	var class models.Class
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&class); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	class.ID = id

	if err := h.Repo.Update(&class); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, class)
}

func (h *ClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid class ID")
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
