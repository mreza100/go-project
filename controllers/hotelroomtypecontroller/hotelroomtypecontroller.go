package hotelroomtypecontroller

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
	var data []models.HotelRoomType

	models.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Show(c *gin.Context) {
	var data models.HotelRoomType
	id := c.Param("id")

	if err := models.DB.First(&data, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, data)
}
func Create(c *gin.Context) {
	hotelid, _ := strconv.ParseInt(c.PostForm("hotel_id"), 10, 64)
	roomname := c.PostForm("room_name")
	title := c.PostForm("title")
	url := c.PostForm("url")
	sequence, _ := strconv.ParseInt(c.PostForm("sequence"), 10, 64)
	status, _ := strconv.ParseBool(c.PostForm("status"))
	file, _ := c.FormFile("image")
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	namafile := c.Request.Host + "/upload/" + newFileName

	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
		return
	}

	result := models.HotelRoomType{HotelId: hotelid, RoomName: roomname, Title: title, Url: url, Sequence: sequence, Image: namafile, Status: status}

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
	var data models.HotelRoomType
	id := c.Param("id")
	hotelid, _ := strconv.ParseInt(c.PostForm("hotel_id"), 10, 64)
	roomname := c.PostForm("room_name")
	title := c.PostForm("title")
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
		models.DB.Model(&data).Where("hotel_room_type_id = ?", id).Update("image", namafile)
		c.SaveUploadedFile(file, "upload/"+newFileName)
	}

	models.DB.Model(&data).Where("hotel_room_type_id = ?", id).Updates(models.HotelRoomType{HotelId: hotelid, RoomName: roomname, Title: title, Url: url, Sequence: sequence, Status: status})

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
}
func Delete(c *gin.Context) {
	var data models.HotelRoomType

	input := map[string]string{"hotel_room_type_id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["hotel_room_type_id"], 10, 64)
	if models.DB.Delete(&data, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}
