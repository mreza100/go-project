package brandcontroller

import (
	"Waringin/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var brands []models.Brands

	models.DB.Find(&brands)

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func Show(c *gin.Context) {
	var brands models.Brands
	id := c.Param("id")

	if err := models.DB.First(&brands, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, brands)
}
func Create(c *gin.Context) {
	// var brands models.Brands
	file, _ := c.FormFile("image")
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	title := c.PostForm("brand_title")
	description := c.PostForm("description")
	url := c.PostForm("url")
	sequence, _ := strconv.ParseInt(c.PostForm("sequence"), 10, 64)
	status, _ := strconv.ParseBool(c.PostForm("status"))
	namafile := c.Request.Host + "/upload/" + newFileName

	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
		return
	}

	result := models.Brands{TitleBrand: title, Description: description, Image: namafile, Url: url, Sequence: sequence, Status: status}

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
	var brands models.Brands
	id := c.Param("id")
	title := c.PostForm("brand_title")
	description := c.PostForm("description")
	url := c.PostForm("url")
	sequence, _ := strconv.ParseInt(c.PostForm("sequence"), 10, 64)
	status, _ := strconv.ParseBool(c.PostForm("status"))

	file, _ := c.FormFile("image")
	if file != nil {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		namafile := c.Request.Host + "/upload/" + newFileName
		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
			return
		}
		models.DB.Model(&brands).Where("id = ?", id).Update("image", namafile)
		c.SaveUploadedFile(file, "upload/"+newFileName)
	}

	models.DB.Model(&brands).Where("id = ?", id).Updates(models.Brands{TitleBrand: title, Description: description, Url: url, Sequence: sequence, Status: status})

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
}
func Delete(c *gin.Context) {
	var brands models.Brands

	input := map[string]string{"id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&brands, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}

func Result(c *gin.Context) {
	var brands []models.Brands

	models.DB.Where("status = ?", false).Find(&brands)

	c.JSON(http.StatusOK, gin.H{"data": brands})
}
