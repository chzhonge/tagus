package api

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"tagus/model"
	"tagus/repository"
)

type SignOnReq struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignOn(c *gin.Context) {
	var r SignOnReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	userRepo := &repository.UserRepository{Repository: repository.Repository{DB: model.DBConn}}

	_, err := userRepo.Find([]string{"id"}, model.User{UserName: r.User})

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "the user already exists"})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	_, err = userRepo.Create(r.User, string(hashPassword), "")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "ok"})
}

func SignIn(c *gin.Context) {
	var r SignOnReq
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	var u model.User

	err := model.DBConn.Select("id", "password").Where(&model.User{UserName: r.User}).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "user or password wrong1"})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(r.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "user or password wrong2"})
		return
	}

	h := md5.New()
	h.Write([]byte(strings.ToLower(u.Password)))

	c.JSON(http.StatusOK, gin.H{"error": false, "data": gin.H{"token": hex.EncodeToString(h.Sum(nil))}})
}
