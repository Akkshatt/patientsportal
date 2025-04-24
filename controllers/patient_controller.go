package controllers

import (
	"net/http"
	"patientsportal/models"


	"github.com/gin-gonic/gin"
)

func CreatePatient(c *gin.Context) {
	var p models.Patient
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := models.DB.Exec("INSERT INTO patients (name, age, gender, address, notes) VALUES ($1, $2, $3, $4, $5)",
		p.Name, p.Age, p.Gender, p.Address, p.Notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Patient created"})
}

func GetPatients(c *gin.Context) {
	rows, err := models.DB.Query("SELECT id, name, age, gender, address, notes FROM patients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var p models.Patient
		rows.Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Address, &p.Notes)
		patients = append(patients, p)
	}
	c.JSON(http.StatusOK, patients)
}

func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var p models.Patient
	err := models.DB.QueryRow("SELECT id, name, age, gender, address, notes FROM patients WHERE id=$1", id).
		Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Address, &p.Notes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var p models.Patient
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := models.DB.Exec("UPDATE patients SET name=$1, age=$2, gender=$3, address=$4, notes=$5 WHERE id=$6",
		p.Name, p.Age, p.Gender, p.Address, p.Notes, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient updated"})
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	_, err := models.DB.Exec("DELETE FROM patients WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
