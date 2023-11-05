package hotellocationcontroller

import (
	"Waringin/models"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var hottellocation []models.HotelLocation

	models.DB.Find(&hottellocation)
	c.JSON(http.StatusOK, gin.H{"data": hottellocation})
}

func Show(c *gin.Context) {
	var hottellocation models.HotelLocation
	id := c.Param("id")

	if err := models.DB.First(&hottellocation, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, hottellocation)
}
func Create(c *gin.Context) {
	name := c.PostForm("name_location")
	sequence, _ := strconv.ParseInt(c.PostForm("sequence"), 10, 64)
	status, _ := strconv.ParseBool(c.PostForm("status"))
	file, _ := c.FormFile("image_desktop")
	file_mobile, _ := c.FormFile("image_mobile")
	extension := filepath.Ext(file.Filename)
	extension_mobile := filepath.Ext(file_mobile.Filename)
	newFileName := uuid.New().String() + extension
	newFileNameMobile := uuid.New().String() + extension_mobile
	namafile := c.Request.Host + "/upload/" + newFileName
	namafilemobile := c.Request.Host + "/upload/" + newFileNameMobile

	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
		return
	}

	if extension_mobile != ".jpg" && extension_mobile != ".png" && extension_mobile != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
		return
	}

	result := models.HotelLocation{ImageDesktop: namafile, ImageMobile: namafilemobile, NameLocation: name, Sequence: sequence, Status: status}

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
	var hotellocation models.HotelLocation
	id := c.Param("id")
	name := c.PostForm("name_location")
	sequence, _ := strconv.ParseInt(c.PostForm("sequence"), 10, 64)
	status, _ := strconv.ParseBool(c.PostForm("status"))

	file, _ := c.FormFile("image_desktop")
	if file != nil {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		namafile := c.Request.Host + "/upload/" + newFileName
		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
			return
		}
		models.DB.Model(&hotellocation).Where("hotel_location_id = ?", id).Update("image_desktop", namafile)
		c.SaveUploadedFile(file, "upload/"+newFileName)
	}

	mobile, _ := c.FormFile("image_mobile")
	if mobile != nil {
		extension := filepath.Ext(mobile.Filename)
		newFileName := uuid.New().String() + extension
		namafile := c.Request.Host + "/upload/" + newFileName
		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
			return
		}
		models.DB.Model(&hotellocation).Where("hotel_location_id = ?", id).Update("image_mobile", namafile)
		c.SaveUploadedFile(mobile, "upload/"+newFileName)
	}

	models.DB.Model(&hotellocation).Where("hotel_location_id = ?", id).Updates(models.HotelLocation{NameLocation: name, Sequence: sequence, Status: status})

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
}
func Delete(c *gin.Context) {
	var hotellocation models.HotelLocation

	input := map[string]string{"hotel_location_id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["hotel_location_id"], 10, 64)
	if models.DB.Delete(&hotellocation, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}

type result struct {
	NameLocation string `json:"location_name"`
	TitleHotel   string `json:"hotel_title"`
	ImageDesktop string `json:"hotel_location_image_desktop"`
	ImageMobile  string `json:"hotel_location_image_mobile"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func Result(c *gin.Context) {
	var result []result
	id := c.Param("id")

	models.DB.Table("hotel_locations").
		Joins("left join hotels on hotels.hotel_location_id = hotel_locations.hotel_location_id").
		Select("hotel_locations.name_location, hotels.title_hotel, hotels.image_desktop, hotels.image_mobile").
		Where("hotel_locations.hotel_location_id = ?", id).
		Find(&result)

	c.JSON(http.StatusOK, gin.H{"data": &result})
}
