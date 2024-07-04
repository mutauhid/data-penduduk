package controllers

import (
	"data-penduduk/models"
	"data-penduduk/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	sqlStatement := `
	SELECT 
		p.id, p.nik, p.name, p.gender, p.dob, p.pob,  p.created_at, p.updated_at,
		prov.id, prov.name, prov.created_at, prov.updated_at,
		reg.id, reg.name, reg.created_at, reg.updated_at,
		dist.id, dist.name, dist.created_at, dist.updated_at,
		p.province_id, p.regency_id, p.district_id
	FROM 
		people p
	JOIN 
		province prov ON p.province_id = prov.id
	JOIN 
		regency reg ON p.regency_id = reg.id
	JOIN 
		district dist ON p.district_id = dist.id
	ORDER BY
		p.id ASC
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var peoples []models.People
	for rows.Next() {
		var people models.People
		err := rows.Scan(
			&people.ID, &people.NIK, &people.Name, &people.Gender, &people.DOB, &people.POB, &people.CreatedAt, &people.UpdatedAt,
			&people.ProvinceID, &people.Province.Name, &people.Province.CreatedAt, &people.Province.UpdatedAt,
			&people.Regency.ID, &people.Regency.Name, &people.Regency.CreatedAt, &people.Regency.UpdatedAt,
			&people.District.ID, &people.District.Name, &people.District.CreatedAt, &people.District.UpdatedAt, &people.ProvinceID, &people.RegencyID, &people.DistrictID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		peoples = append(peoples, people)
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "SUCCESS", "data": peoples})
}

func GetPeopleByNIK(c *gin.Context) {
	id := c.Param("nik")
	sqlStatement := `
	SELECT 
		p.id, p.nik, p.name, p.gender, p.dob, p.pob, p.created_at, p.updated_at,
		prov.id, prov.name, prov.created_at, prov.updated_at,
		reg.id, reg.name, reg.created_at, reg.updated_at,
		dist.id, dist.name, dist.created_at, dist.updated_at
	FROM 
		people p
	JOIN 
		province prov ON p.province_id = prov.id
	JOIN 
		regency reg ON p.regency_id = reg.id
	JOIN 
		district dist ON p.district_id = dist.id
	WHERE 
		p.nik = $1
	`

	var people models.People
	err := db.QueryRow(sqlStatement, id).Scan(
		&people.ID, &people.NIK, &people.Name, &people.Gender, &people.DOB, &people.POB, &people.CreatedAt, &people.UpdatedAt,
		&people.Province.ID, &people.Province.Name, &people.Province.CreatedAt, &people.Province.UpdatedAt,
		&people.Regency.ID, &people.Regency.Name, &people.Regency.CreatedAt, &people.Regency.UpdatedAt,
		&people.District.ID, &people.District.Name, &people.District.CreatedAt, &people.District.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "People not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "SUCCESS", "data": people})
}

func CreatePeople(c *gin.Context) {
	sqlStatement := `INSERT INTO people (id, nik, name, gender, dob, pob, province_id, regency_id, district_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	var people struct {
		ID         string `json:"id"`
		NIK        string `json:"nik"`
		Name       string `json:"name"`
		Gender     string `json:"gender"`
		DOB        string `json:"dob"`
		POB        string `json:"pob"`
		ProvinceID string `json:"province_id"`
		RegencyID  string `json:"regency_id"`
		DistrictID string `json:"district_id"`
	}

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dob, err := time.Parse("2006-Jan-02", people.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	newPerson := models.People{
		ID:         utils.GenerateID(),
		NIK:        people.NIK,
		Name:       people.Name,
		Gender:     people.Gender,
		DOB:        dob,
		POB:        people.POB,
		ProvinceID: people.ProvinceID,
		RegencyID:  people.RegencyID,
		DistrictID: people.DistrictID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	newPerson.NIK = utils.GenerateNIK(newPerson)
	err = db.QueryRow(sqlStatement, newPerson.ID, newPerson.NIK, newPerson.Name, newPerson.Gender, newPerson.DOB, newPerson.POB, newPerson.ProvinceID, newPerson.RegencyID, newPerson.DistrictID, newPerson.CreatedAt, newPerson.UpdatedAt).Scan(&newPerson.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newPerson.CreatedAt = time.Now()
	newPerson.UpdatedAt = time.Now()
	c.JSON(http.StatusOK, newPerson)
}

func UpdatePeople(c *gin.Context) {
	sqlStatement := `UPDATE people SET name=$1, province_id=$2, regency_id=$3, district_id=$4, updated_at=$5 WHERE id=$6`
	id := c.Param("id")
	var people models.People

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	people.UpdatedAt = time.Now()

	result, err := db.Exec(sqlStatement, people.Name, people.ProvinceID, people.RegencyID, people.DistrictID, people.UpdatedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "People with ID not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update People Successfully",
	})

}

func DeletePeople(c *gin.Context) {
	sqlStatement := `DELETE FROM people WHERE id=$1`
	id := c.Param("id")

	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "People with ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete People Successfully",
	})
}
