package hotelcontroller

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
	var hotel []models.Hotel

	models.DB.Find(&hotel)
	c.JSON(http.StatusOK, gin.H{"data": hotel})
}

func Show(c *gin.Context) {
	var hotel models.Hotel
	id := c.Param("id")

	if err := models.DB.First(&hotel, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, hotel)
}
func Create(c *gin.Context) {
	hotellocation, _ := strconv.ParseInt(c.PostForm("hotel_location_id"), 10, 64)
	brandid, _ := strconv.ParseInt(c.PostForm("brand_id"), 10, 64)
	title := c.PostForm("title_hotel")
	slug := c.PostForm("slug")
	description := c.PostForm("description")
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

	result := models.Hotel{ImageDesktop: namafile, ImageMobile: namafilemobile, HotelLocationId: hotellocation, BrandId: brandid, TitleHotel: title, Slug: slug, Description: description, Status: status}

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
	var hotel models.Hotel
	id := c.Param("id")
	hotellocation, _ := strconv.ParseInt(c.PostForm("hotel_location_id"), 10, 64)
	brandid, _ := strconv.ParseInt(c.PostForm("brand_id"), 10, 64)
	title := c.PostForm("title_hotel")
	slug := c.PostForm("slug")
	description := c.PostForm("description")
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
		models.DB.Model(&hotel).Where("hotel_id = ?", id).Update("image_desktop", namafile)
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
		models.DB.Model(&hotel).Where("hotel_id = ?", id).Update("image_mobile", namafile)
		c.SaveUploadedFile(mobile, "upload/"+newFileName)
	}

	models.DB.Model(&hotel).Where("hotel_id = ?", id).Updates(models.Hotel{HotelLocationId: hotellocation, BrandId: brandid, TitleHotel: title, Slug: slug, Description: description, Status: status})

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
}
func Delete(c *gin.Context) {
	var hotel models.Hotel

	input := map[string]string{"hotel_id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["hotel_id"], 10, 64)
	if models.DB.Delete(&hotel, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}

type data struct {
	TitleHotel   string `json:"hotel_title"`
	ImageDesktop string `json:"hotel_location_image_desktop"`
	ImageMobile  string `json:"hotel_location_image_mobile"`
	Description  string `json:"description"`
	RoomName     string `json:"room_name"`
	Title        string `json:"title"`
	Image        string `json:"hotel_room_type_image"`
	Url          string `json:"url"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func Result(c *gin.Context) {
	var data []data
	id := c.Param("id")

	models.DB.Table("hotels as h").
		Joins("left join hotel_room_types as t on t.hotel_id = h.hotel_id").
		Joins("left join hotel_facilities as f on f.hotel_id = h.hotel_id").
		Joins("left join hotel_galleries as g on g.hotel_id = h.hotel_id").
		Joins("left join hotel_informations as i on i.hotel_id = h.hotel_id").
		Select("h.title_hotel, h.image_desktop, h.image_mobile, h.description, t.room_name, t.title, t.image, t.url").
		Where("h.hotel_id = ?", id).
		Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
