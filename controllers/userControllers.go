package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Itsu8/Auth/initializers"
	"github.com/Itsu8/Auth/modules"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	//Structure with user registration input
	var user modules.User
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read user data"})
		return
	}
	//Check if username taken
	if err := initializers.DB.Where("username = ?", user.GetUsername()).First(&user).Error; err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Username already taken"})
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := modules.User{
		ID:         uuid.New(),
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Deleted_at: time.Time{},
		Password:   string(hashPass),
		Username:   user.Username,
		Age:        user.Age,
		Bio:        user.Bio,
	}

	//creating user on database from recieved
	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered!"})
}

func LoginUser(c *gin.Context) {
	//user login input
	var userInput modules.User
	if c.Bind(&userInput) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read user data"})
		return
	}

	//Look up requested user
	var user modules.User
	if initializers.DB.First(&user, "username = ?", userInput.GetUsername()).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid username or password"})
		return
	}

	//Compare password input and actual user
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)) != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid username or password"})
		return
	}

	//Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//Sign the token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*25*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Username:": user.GetUsername(),
		"Age:":      user.GetUserAge(),
		"Bio:":      user.GetUserBio(),
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
