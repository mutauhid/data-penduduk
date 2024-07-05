package controllers

import (
	"data-penduduk/models"
	"data-penduduk/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte("secret_key")

func Register(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	user.Password = string(hashedPassword)
	user.ID = utils.GenerateID()

	sqlStatement := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, user.ID, user.Username, user.Password)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Registration Success", nil)

}

func Login(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var storedUser models.Users
	sqlStatement := `SELECT id, username, password FROM users WHERE username = $1`
	err := db.QueryRow(sqlStatement, user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.JSONResponse(c, http.StatusUnauthorized, "Invalid username or password", nil)
		} else {
			utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": storedUser.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
