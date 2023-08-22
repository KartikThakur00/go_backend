package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type View3 struct {
	Name  string `json:"name"`
	Count float64   `json:"count"`
}

type View4 struct {
	Manager string `json:"manager"`
	Leavecount   int    `json:"leavecount"`
}

type View6 struct {
	TeamName      string `json:"team"`
	LeaveType string `json:"leaveType"`
	Count     float64    `json:"count"`
}

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=api port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func main() {
	router := gin.Default()

	Connect()

	router.GET("/view3", func(ctx *gin.Context) {
		view := []View3{}
		DB.Raw("SELECT * from view3").Scan(&view)
		ctx.JSON(200, &view)
	})

	router.GET("/view4", func(ctx *gin.Context) {
		view := []View4{}
		DB.Raw("SELECT * from view4").Scan(&view)
		ctx.JSON(200, &view)
	})

	router.GET("/view6", func(ctx *gin.Context) {
		view := []View6{}
		DB.Raw("SELECT * from view6").Scan(&view)
		ctx.JSON(200, &view)
	})

	router.Run(":8080")
}
