package authcontroller

import (
	"Waringin/models"
	"Waringin/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUser struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func Index(c *gin.Context) {
	var data []GetUser

	models.DB.Table("users").
		Select("id,email,name").
		Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"test": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password
	u.Name = input.Name

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func Update(c *gin.Context) {
	var user models.User
	userid, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	hotelid := c.PostForm("hotel_id")
	password := []byte(c.PostForm("password"))
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err := models.DB.Create(&models.HotelOwner{UserId: userid, HotelId: hotelid}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	if err := models.DB.Model(&user).Where("id = ?", userid).Updates(&models.User{Password: string(hashedPassword)}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "berhasil"})

}
