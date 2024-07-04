package controllers

import (
	"data-penduduk/models"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDistrict(c *gin.Context) {
	sqlStatement := `
	SELECT
		d.id, d.name, d.regency_id, r.id, r.name,r.created_at, r.updated_at,p.id, p.name, p.created_at, p.updated_at, d.created_at, d.updated_at
	FROM 
		district d
	JOIN
		regency r on r.id = d.regency_id
	JOIN 
		province p on p.id = r.province_id
	ORDER BY
		d.id ASC;
	`

	var districts []models.District

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var district models.District
		err := rows.Scan(&district.ID, &district.Name, &district.Regency.ID, &district.Regency.ID, &district.Regency.Name, &district.Regency.CreatedAt, &district.Regency.UpdatedAt, &district.Regency.Province.ID, &district.Regency.Province.Name, &district.Regency.Province.CreatedAt, &district.Regency.Province.UpdatedAt, &district.CreatedAt, &district.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		districts = append(districts, district)
	}
	c.JSON(http.StatusOK, districts)
}

func CreateDistrict(c *gin.Context) {
	sqlStatement := `
	INSERT INTO 
		district (id, name, regency_id, created_at, updated_at) 
	VALUES 
		($1, $2, $3, $4, $5) 
	RETURNING 
		id`

	sqlRow := `
	SELECT r.id, r.name, r.province_id, p.id, p.name, p.created_at, p.updated_at, r.created_at, r.updated_at 
	FROM regency r 
	JOIN province p on p.id = r.province_id
	WHERE r.id = $1
	`

	var input struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		RegencyID  string `json:"regency_id"`
		ProvinceID string `json:"province_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var regency models.Regency
	err := db.QueryRow(sqlRow, input.RegencyID).Scan(&regency.ID, &regency.Name, &regency.Province.ID, &regency.Province.ID, &regency.Province.Name, &regency.Province.CreatedAt, &regency.Province.UpdatedAt, &regency.CreatedAt, &regency.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Regency ID"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	var regencyID string
	err = db.QueryRow(sqlStatement, input.ID, input.Name, input.RegencyID, createdAt, updatedAt).Scan(&regencyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	district := models.District{
		ID:        regencyID,
		Name:      input.Name,
		Regency:   regency,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	c.JSON(http.StatusOK, district)
}

func UpdateDistrict(c *gin.Context) {
	sqlStatement := `
	UPDATE 
		district 
	SET 
		name=$1, regency_id=$2, updated_at=$3 
	WHERE 
		id=$4`
	var input struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		RegencyID string `json:"regency_id"`
	}
	id := c.Param("id")
	var regency models.Regency

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	regency.UpdatedAt = time.Now()
	fmt.Println(&regency)
	fmt.Println(id)

	result, err := db.Exec(sqlStatement, input.Name, input.RegencyID, regency.UpdatedAt, id)
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
		"message": "Update Regency Successfully",
	})

}

func DeleteDistrict(c *gin.Context) {
	sqlStatement := `DELETE FROM District WHERE id=$1`
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
			"error": "District with ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete District Successfully",
	})
}
