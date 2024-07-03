package controllers

import (
	"data-penduduk/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDistrict(c *gin.Context) {
	sqlStatement := `SELECT * FROM district`

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
		err := rows.Scan(&district.ID, &district.Name, &district.RegencyID, &district.CreatedAt, &district.UpdatedAt)
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
	sqlStatement := `INSERT INTO district(id, name, regency_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var district models.District
	if err := c.ShouldBindJSON(&district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	district.CreatedAt = time.Now()
	district.UpdatedAt = time.Now()

	err := db.QueryRow(sqlStatement, district.ID, district.Name, district.RegencyID, district.CreatedAt, district.UpdatedAt).Scan(&district.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, district)
}

func UpdateDistrict(c *gin.Context) {
	sqlStatement := `UPDATE district SET name=$1,regency_id=$2, updated_at=$3 WHERE id=$4`
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	var district models.District

	if err := c.ShouldBindJSON(&district); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	district.UpdatedAt = time.Now()

	_, err := db.Exec(sqlStatement, district.Name, district.RegencyID, district.UpdatedAt, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update district Successfully",
	})

}

func DeleteDistrict(c *gin.Context) {
	sqlStatement := `DELETE FROM district WHERE id=$1`
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete district Successfully",
	})
}
