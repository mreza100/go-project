package articlenewscontroller

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
	var data []models.ArticleNews

	models.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Show(c *gin.Context) {
	var data models.ArticleNews
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
	articleid, _ := strconv.ParseInt(c.PostForm("article_categories_id"), 10, 64)
	title := c.PostForm("title")
	slug := c.PostForm("slug")
	metadescription := c.PostForm("meta_description")
	metakeyword := c.PostForm("meta_keyword")
	contentnews := c.PostForm("content_news")
	status, _ := strconv.ParseBool(c.PostForm("status"))
	file, _ := c.FormFile("image_thumbnail")
	icon, _ := c.FormFile("image_banner")
	extension := filepath.Ext(file.Filename)
	extensionicon := filepath.Ext(icon.Filename)
	newFileName := uuid.New().String() + extension
	newFileNameMobile := uuid.New().String() + extensionicon
	namafile := c.Request.Host + "/upload/" + newFileName
	namafilemobile := c.Request.Host + "/upload/" + newFileNameMobile

	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
		return
	}
	if extensionicon != ".jpg" && extensionicon != ".png" && extensionicon != ".jpeg" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
		return
	}

	result := models.ArticleNews{ArticleCategoryId: articleid, Title: title, Slug: slug, ImageBanner: namafile, ImageThumbnail: namafilemobile, MetaDescription: metadescription, MetaKeyword: metakeyword, ContentNews: contentnews, Status: status}

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
	var data models.ArticleNews
	id := c.Param("id")
	articleid, _ := strconv.ParseInt(c.PostForm("article_categories_id"), 10, 64)
	title := c.PostForm("title")
	slug := c.PostForm("slug")
	metadescription := c.PostForm("meta_description")
	metakeyword := c.PostForm("meta_keyword")
	contentnews := c.PostForm("content_news")
	status, _ := strconv.ParseBool(c.PostForm("status"))

	file, _ := c.FormFile("image_thumbnail")
	if file != nil {
		extension := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extension
		namafile := c.Request.Host + "/upload/" + newFileName
		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
			return
		}
		models.DB.Model(&data).Where("id = ?", id).Update("image_thumbnail", namafile)
		c.SaveUploadedFile(file, "upload/"+newFileName)
	}

	icon, _ := c.FormFile("image_banner")
	if icon != nil {
		extension := filepath.Ext(icon.Filename)
		newFileName := uuid.New().String() + extension
		namafile := c.Request.Host + "/upload/" + newFileName
		if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "file harus jpg, jpeg atau png"})
			return
		}
		models.DB.Model(&data).Where("id = ?", id).Update("image_banner", namafile)
		c.SaveUploadedFile(icon, "upload/"+newFileName)
	}

	models.DB.Model(&data).Where("id = ?", id).Updates(models.ArticleNews{ArticleCategoryId: articleid, Title: title, Slug: slug, MetaDescription: metadescription, MetaKeyword: metakeyword, ContentNews: contentnews, Status: status})

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di update"})
}
func Delete(c *gin.Context) {
	var data models.ArticleNews

	input := map[string]string{"id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&data, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}
