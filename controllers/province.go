package controllers

import (
	"data-penduduk/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProvince(c *gin.Context) {
	sqlStatement := `SELECT * FROM province`

	var provinces []models.Province

	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var province models.Province
		err := rows.Scan(&province.ID, &province.Name, &province.CreatedAt, &province.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		provinces = append(provinces, province)
	}
	c.JSON(http.StatusOK, provinces)
}

func CreateProvince(c *gin.Context) {
	sqlStatement := `INSERT INTO province (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var province models.Province
	if err := c.ShouldBindJSON(&province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	province.CreatedAt = time.Now()
	province.UpdatedAt = time.Now()

	err := db.QueryRow(sqlStatement, province.ID, province.Name, province.CreatedAt, province.UpdatedAt).Scan(&province.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, province)
}

func UpdateProvince(c *gin.Context) {
	sqlStatement := `UPDATE province SET name=$1, updated_at=$2 WHERE id=$3`
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	var province models.Province

	if err := c.ShouldBindJSON(&province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	province.UpdatedAt = time.Now()

	_, err := db.Exec(sqlStatement, province.Name, province.UpdatedAt, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Province Successfully",
	})

}

func DeleteProvince(c *gin.Context) {
	sqlStatement := `DELETE FROM province WHERE id=$1`
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Province Successfully",
	})
}
