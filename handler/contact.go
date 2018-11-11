package handler

import (
	"net/http"

	"../models"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAllContacts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	contacts := []models.Contact{}
	db.Debug().Find(&contacts)
	var count int
	db.Model(&models.Contact{}).Count(&count)
	// db.Find(&contacts)
	spew.Dump(count)
	respondJSON(w, http.StatusOK, contacts)
}

// func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	employee := models.Contact{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&employee); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, employee)
// }

// func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&employee); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	if err := db.Delete(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }

// func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	employee.Disable()
// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	name := vars["name"]
// 	employee := getEmployeeOr404(db, name, w, r)
// 	if employee == nil {
// 		return
// 	}
// 	employee.Enable()
// 	if err := db.Save(&employee).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, employee)
// }

// // getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
// func getEmployeeOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *models.Contact {
// 	employee := models.Contact{}
// 	if err := db.First(&employee, models.Contact{Name: name}).Error; err != nil {
// 		respondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return &employee
// }
