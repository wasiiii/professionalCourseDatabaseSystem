package main

import (
	"fmt"

	"./gin/controller"
	"./gin/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}

	router := gin.Default()
	router.Use(controller.Cors())
	router.POST("/stuLogin", controller.StuLogin)
	router.POST("/teaLogin", controller.TeaLogin)
	router.POST("/getInformation", controller.GetInformation)
	router.POST("/getPersonalTable", controller.GetPersonalTable)
	router.POST("/getClassTable", controller.GetClassTable)
	router.POST("/courseSelectInit", controller.CourseSelectInit)
	router.POST("/courseSelect", controller.CourseSelect)
	router.POST("/getAllStu", controller.GetAllStu)
	router.POST("/getAllTea", controller.GetAllTea)
	router.POST("/changeClassHeadmaster", controller.ChangeClassHeadmaster)
	router.POST("/getCourseInfo", controller.GetCourseInfo)
	router.POST("/changeGroup", controller.ChangeGroup)
	router.POST("/gradesRecordInit", controller.GradesRecordInit)
	router.POST("/commitGrades", controller.CommitGrades)

	router.Run(":8080")
}
