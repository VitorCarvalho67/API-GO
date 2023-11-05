package main

import (
	"APIgo/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func FazerAlgo() {
	fmt.Println("Começando a fazer algo...")
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	fmt.Println("Terminando de fazer algo...")
}

func GetDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/Goapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := GetDb()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Create(&user)
		c.JSON(200, user)
	})

	// mostrar todos os usuarios
	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		db.Find(&users)
		c.JSON(200, users)
	})

	// mostrar um usuario
	r.GET("/users/:id", func(c *gin.Context) {
		var user models.User
		db.First(&user, c.Param("id"))
		c.JSON(200, user)
	})

	// login
	r.POST("/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		var userDB models.User
		db.Where("email = ?", user.Email).First(&userDB)
		if userDB.Password != user.Password {
			c.JSON(401, gin.H{
				"error": "Senha incorreta",
			})
			return
		}
		c.JSON(200, userDB)
	})

	r.Run()

	// listen and serve on 0.0.0.0:8080

	// c := make(chan int)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		FazerAlgo()
	// 		c <- 1
	// 	}()
	// }
	// c <- 1

	// db.Create(&models.User{
	// 	Name:     "Teste",
	// 	Email:    "vitor@gmail.com",
	// 	Password: "123456",
	// })

}
