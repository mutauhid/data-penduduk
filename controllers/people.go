package controllers

import (
	"data-penduduk/models"
	"data-penduduk/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	sqlStatement := `SELECT * FROM people`

	var peoples []models.People

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var people models.People
		err := rows.Scan(&people.ID, &people.NIK, &people.Name, people.Gender, people.DOB, people.POB, &people.ProvinceID, people.RegencyID, people.DistrictID, &people.CreatedAt, &people.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		peoples = append(peoples, people)
	}
	c.JSON(http.StatusOK, peoples)
}

func CreatePeople(c *gin.Context) {
	sqlStatement := `INSERT INTO people (id,nik, name,gender, dob, pob, province_id, regency_id, district_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	var people struct {
		NIK        string `json:"nik"`
		Name       string `json:"name"`
		Gender     string `json:"gender"`
		DOB        string `json:"dob"` // String to handle custom format
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
	dob, err := time.Parse("02-01-2006", people.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use DD-MM-YYYY."})
		return
	}

	newPerson := models.People{
		Name:       people.Name,
		Gender:     people.Gender,
		DOB:        dob,
		POB:        people.POB,
		ProvinceID: people.ProvinceID,
		RegencyID:  people.RegencyID,
		DistrictID: people.DistrictID,
	}
	newPerson.NIK = utils.GenerateNIK(newPerson)
	err = db.QueryRow(sqlStatement, newPerson.ID, newPerson.NIK, newPerson.Name, newPerson.Gender, newPerson.DOB, newPerson.POB, newPerson.ProvinceID, newPerson.RegencyID, newPerson.DistrictID, newPerson.UpdatedAt).Scan(&newPerson.ID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newPerson.CreatedAt = time.Now()
	newPerson.UpdatedAt = time.Now()
	c.JSON(http.StatusOK, people)
}

func UpdatePeople(c *gin.Context) {
	sqlStatement := `UPDATE people SET name=$1, province_id=$2, updated_at=$3 WHERE id=$4`
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	var people models.People

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	people.UpdatedAt = time.Now()

	_, err := db.Exec(sqlStatement, people.Name, people.Gender, people.DOB, people.POB, people.ProvinceID, people.RegencyID, people.DistrictID, people.UpdatedAt, id)
	fmt.Println(db.Exec(sqlStatement, people.Name, people.ProvinceID, people.UpdatedAt, id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Regency Successfully",
	})

}

func DeletePeople(c *gin.Context) {
	sqlStatement := `DELETE FROM regency WHERE id=$1`
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Regency Successfully",
	})
}
