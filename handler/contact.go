package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAllContacts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	contacts := []models.Contact{}
	if err := db.Find(&contacts).Error; err != nil {
		log.Fatal(err)
	}

	respondJSON(w, http.StatusOK, contacts)
}

func CreateContact(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	contact := models.Contact{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&contact); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&contact).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, contact)
}

func GetContact(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	contact := getContactOr404(db, id, w, r)
	if contact == nil {
		return
	}
	respondJSON(w, http.StatusOK, contact)
}

func UpdateContact(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	contact := getContactOr404(db, id, w, r)
	if contact == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&contact); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&contact).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, contact)
}

func DeleteContact(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	contact := getContactOr404(db, id, w, r)
	if contact == nil {
		return
	}
	if err := db.Delete(&contact).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getContactOr404 gets a contact instance if exists, or respond the 404 error otherwise
func getContactOr404(db *gorm.DB, id int64, w http.ResponseWriter, r *http.Request) *models.Contact {
	contact := models.Contact{}
	if err := db.First(&contact, models.Contact{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &contact
}
