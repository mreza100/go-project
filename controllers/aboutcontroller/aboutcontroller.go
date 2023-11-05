package aboutcontroller

import (
	"Waringin/models"
	"Waringin/utils/token"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var abouts []models.About

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

	if u.Name == "admin" {

	}

	models.DB.Find(&abouts)
	c.JSON(http.StatusOK, gin.H{"data": abouts})
}

func Show(c *gin.Context) {
	var abouts models.About
	id := c.Param("id")

	if err := models.DB.First(&abouts, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, abouts)
}
func Create(c *gin.Context) {
	// var brands models.Brands
	file, _ := c.FormFile("image")
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	namafile := c.Request.Host + "/upload/" + newFileName
	title := c.PostForm("title")
	body := c.PostForm("body_content")

	result := models.About{Title: title, Image: namafile, BodyContent: body}

	if err := c.SaveUploadedFile(file, "upload/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	models.DB.Create(&result)
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di tambah"})
}

func Update(c *gin.Context) {
	var abouts models.About
	id := c.Param("id")
	title := c.PostForm("title")
	body := c.PostForm("body_content")

	file, _ := c.FormFile("image")
	if file != nil {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		namafile := c.Request.Host + "/upload/" + newFileName
		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
			return
		}
		models.DB.Model(&abouts).Where("id = ?", id).Update("image", namafile)
		c.SaveUploadedFile(file, "upload/"+newFileName)
	}

	models.DB.Model(&abouts).Where("id = ?", id).Updates(models.About{Title: title, BodyContent: body})

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
}
func Delete(c *gin.Context) {
	var abouts models.About

	input := map[string]string{"id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&abouts, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}

func Result(c *gin.Context) {
	var data []models.About

	models.DB.Find(&data).Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
