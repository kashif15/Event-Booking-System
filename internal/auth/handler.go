package auth

import (
	"event-booking-api/internal/user"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)


type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}


type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}


func RegisterHandler(c *gin.Context) {
	var req registerRequest

	err := c.ShouldBindBodyWithJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request data",
		})
		return
	}

	user, err := Register(req.Name, req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := GenerateToken(user.ID, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})

}

func LoginHandler(c *gin.Context) {
	var req loginRequest

	err := c.ShouldBindBodyWithJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request data",
		})
		return
	}

	user, err := Login(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := GenerateToken(user.ID, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	refereshToken := GenerateRefreshToken()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	err = SaveRefreshToken(user.ID, refereshToken, expiresAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	
	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"access_token": accessToken,
		"refresh_token": refereshToken,
	})


}

func RefreshHandler(c *gin.Context) {
	var req refreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	rt, err := GetRefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := user.GetByID(rt.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	newAccessToken, _ := GenerateToken(user.ID, user.Role)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}

func LogoutHandler(c *gin.Context) {
	var req refreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := DeleteRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})

}