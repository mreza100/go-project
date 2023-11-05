package models

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	DB.AutoMigrate(&Brands{}, &User{}, &About{}, &HotelLocation{}, &SpecialOffer{}, &ManagementInvestor{}, &HotelOwner{},
		&Hotel{}, &HotelFacilities{}, &HotelRoomType{}, &HotelGallery{}, &HotelInformation{}, &ArticleNews{}, &ArticleCategories{})
}

func Seeder() {
	var user User

	if DB.Model(&user).Where("id = ?", 1).Find(&user).RowsAffected == 0 {
		user := &User{Name: "admin", Password: "MicroAd!234%", Email: "devmob@microad.co.id"}
		DB.Create(&user)
	}
}

func ConnectCDN() {
	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String("https://sgp1.cdn.digitaloceanspaces.com"),
		Region:           aws.String("sgp1"),
		S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	input := &s3.ListObjectsInput{
		Bucket: aws.String("maimedia"),
	}

	objects, err := s3Client.ListObjects(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, obj := range objects.Contents {
		fmt.Println(aws.StringValue(obj.Key))
	}
}
