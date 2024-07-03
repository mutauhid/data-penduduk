package controllers

import (
	"data-penduduk/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRegency(c *gin.Context) {
	sqlStatement := `SELECT * FROM regency`

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
		err := rows.Scan(&regency.ID, &regency.Name, &regency.ProvinceID, &regency.CreatedAt, &regency.UpdatedAt)
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
	sqlStatement := `INSERT INTO regency (id, name, province_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var regency models.Regency
	if err := c.ShouldBindJSON(&regency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	regency.CreatedAt = time.Now()
	regency.UpdatedAt = time.Now()

	err := db.QueryRow(sqlStatement, regency.ID, regency.Name, regency.ProvinceID, regency.CreatedAt, regency.UpdatedAt).Scan(&regency.ID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, regency)
}

func UpdateRegency(c *gin.Context) {
	sqlStatement := `UPDATE regency SET name=$1, province_id=$2, updated_at=$3 WHERE id=$4`
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	var regency models.Regency

	if err := c.ShouldBindJSON(&regency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	regency.UpdatedAt = time.Now()

	_, err := db.Exec(sqlStatement, regency.Name, regency.ProvinceID, regency.UpdatedAt, id)
	fmt.Println(db.Exec(sqlStatement, regency.Name, regency.ProvinceID, regency.UpdatedAt, id))
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

func DeleteRegency(c *gin.Context) {
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
