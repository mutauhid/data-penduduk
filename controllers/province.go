package controllers

import (
	"data-penduduk/models"
	"data-penduduk/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProvince(c *gin.Context) {
	sqlStatement := `SELECT * FROM province`

	var provinces []models.Province

	rows, err := db.Query(sqlStatement)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var province models.Province
		err := rows.Scan(&province.ID, &province.Name, &province.CreatedAt, &province.UpdatedAt)
		if err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		provinces = append(provinces, province)
	}
	utils.JSONResponse(c, http.StatusOK, "Success", provinces)
}

func CreateProvince(c *gin.Context) {
	sqlStatement := `INSERT INTO province (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var province models.Province
	if err := c.ShouldBindJSON(&province); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	province.CreatedAt = time.Now()
	province.UpdatedAt = time.Now()

	err := db.QueryRow(sqlStatement, province.ID, province.Name, province.CreatedAt, province.UpdatedAt).Scan(&province.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Success", province)
}

func UpdateProvince(c *gin.Context) {
	sqlStatement := `UPDATE province SET name=$1, updated_at=$2 WHERE id=$3`
	id := c.Param("id")
	fmt.Println(id)
	var province models.Province

	if err := c.ShouldBindJSON(&province); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	province.UpdatedAt = time.Now()

	result, err := db.Exec(sqlStatement, province.Name, province.UpdatedAt, id)

	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if rowsAffected == 0 {
		utils.JSONResponse(c, http.StatusNotFound, "Province with ID not found", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "Update Province Successfully", nil)

}

func DeleteProvince(c *gin.Context) {
	sqlStatement := `DELETE FROM province WHERE id=$1`
	id := c.Param("id")

	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if rowsAffected == 0 {
		utils.JSONResponse(c, http.StatusNotFound, "Province with ID not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Delete Province Successfully", nil)
}
