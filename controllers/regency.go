package controllers

import (
	"data-penduduk/models"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRegency(c *gin.Context) {
	sqlStatement := `
	SELECT
		r.id, r.name,r.province_id, p.id, p.name,p.created_at, p.updated_at, r.created_at, r.updated_at
	FROM 
		regency r
	JOIN 
		province p on p.id = r.province_id
	ORDER BY
		r.id ASC
	`

	var regencys []models.Regency

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var regency models.Regency
		err := rows.Scan(&regency.ID, &regency.Name, &regency.Province.ID, &regency.Province.ID, &regency.Province.Name, &regency.Province.CreatedAt, &regency.Province.UpdatedAt, &regency.CreatedAt, &regency.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		regencys = append(regencys, regency)
	}
	c.JSON(http.StatusOK, regencys)
}

func CreateRegency(c *gin.Context) {
	sqlStatement := `
	INSERT INTO 
		regency (id, name, province_id, created_at, updated_at) 
	VALUES 
		($1, $2, $3, $4, $5) 
	RETURNING 
		id`

	var input struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		ProvinceID string `json:"province_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Request Body:", c.Request.Body)

	var province models.Province
	err := db.QueryRow("SELECT id, name, created_at, updated_at FROM province WHERE id = $1", input.ProvinceID).Scan(&province.ID, &province.Name, &province.CreatedAt, &province.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province ID"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	var regencyID string
	err = db.QueryRow(sqlStatement, input.ID, input.Name, input.ProvinceID, createdAt, updatedAt).Scan(&regencyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	regency := models.Regency{
		ID:        regencyID,
		Name:      input.Name,
		Province:  province,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	c.JSON(http.StatusOK, regency)
}

func UpdateRegency(c *gin.Context) {

	sqlStatement := `
	UPDATE 
		regency 
	SET 
		name=$1, province_id=$2, updated_at=$3 
	WHERE 
		id=$4`
	var input struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		ProvinceID string `json:"province_id"`
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

	result, err := db.Exec(sqlStatement, input.Name, input.ProvinceID, regency.UpdatedAt, id)
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
			"error": "Regency with ID not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Regency Successfully",
	})

}

func DeleteRegency(c *gin.Context) {
	sqlStatement := `DELETE FROM regency WHERE id=$1`
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
			"error": "Regency with ID not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Regency Successfully",
	})
}
