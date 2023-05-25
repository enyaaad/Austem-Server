package handlers

import (
	"AustemServer/models"
	"AustemServer/posgtresql/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func Autharization(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Json"})
		return
	}

	var user models.User
	if err := db.DB.Where("name = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid data"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(loginData.Password), bcrypt.DefaultCost)

	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid data"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Token{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	key, exist := os.LookupEnv("key")
	if !exist {
		panic("no key for jwt")
	}

	//correct json
	//{
	//	"username": "asd",
	//	"password": "123"
	//}

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "token": tokenString})
}
