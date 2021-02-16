package controller

import (
	"fmt"
	"net/http"

	"../model"

	"github.com/gin-gonic/gin"
)

//Cors 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

//StuLogin 学生登录
func StuLogin(c *gin.Context) {
	var stu = model.Stu{}
	err := c.BindJSON(&stu)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		fmt.Println("err: ", err.Error())
		return
	}
	flag := model.CheckStu(stu)
	c.JSON(http.StatusOK, gin.H{
		"check": flag,
	})
}

//TeaLogin 教师登录
func TeaLogin(c *gin.Context) {
	var tea = model.Tea{}
	err := c.BindJSON(&tea)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}
	flag, teaType, isHeadmaster := model.CheckTea(tea)
	c.JSON(http.StatusOK, gin.H{
		"isHeadmaster": isHeadmaster,
		"type":         teaType,
		"check":        flag,
	})
}

//GetInformation 获取个人信息
func GetInformation(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	//先检验账号和role是否对应
	if !model.CheckAccountAndRole(data) {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	if data.Role == "0" {
		stu, rank := model.GetStuInfo(data.Account)
		c.JSON(http.StatusOK, gin.H{
			"check":           true,
			"stuID":           stu.StuID,
			"stuName":         stu.StuName,
			"stuSex":          stu.StuSex,
			"stuClass":        stu.StuClass,
			"classHeadmaster": stu.ClassHeadmaster,
			"gpa":             stu.Gpa,
			"rank":            rank,
		})

	} else {
		var tea = model.GetTeaInfo(data.Account)
		var class string
		for i, c := range tea.TeaClass {
			if i == 0 {
				class += c
			} else {
				class += "，"
				class += c
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"check":   true,
			"teaID":   tea.TeaID,
			"teaName": tea.TeaName,
			"class":   class,
		})
	}
}

func GetPersonalTable(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	if data.Role == "0" {
		table, err := model.GetStuTable(data.Account)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"check":   false,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check": true,
			"table": table,
		})
	} else {
		table, err := model.GetTeaTable(data.Account)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"check":   false,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check": true,
			"table": table,
		})
	}
}

func GetClassTable(c *gin.Context) {
	var data = model.AccountAndRoleWhenChangeClassTable{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	if data.Role == "0" {
		table, err := model.GetClassTableFromStu(data.Account)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"check":   false,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check": true,
			"table": table,
		})
	} else {
		if data.Class == "" {
			table, class, err := model.GetClassTableFromTea(data.Account)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"check":   false,
					"message": err,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"check": true,
				"table": table,
				"class": class,
			})
		} else {
			table, err := model.GetClassTableFromTeaChange(data.Class)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, gin.H{
					"check":   false,
					"message": err,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"check": true,
				"table": table,
			})
		}
	}
}

//CourseSelectInit 选课页面初始化
func CourseSelectInit(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	//先检验账号和role是否对应
	if !model.CheckAccountAndRole(data) {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": "账号角色不匹配",
		})
		return
	}

	flag, err := model.IsCourseSelectTime()
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}
	if flag {
		table, err := model.GetCourseWaitSelectTable(data.Account)
		if err != nil {
			fmt.Println("err: ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"check":   false,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check":  true,
			"isTime": true,
			"table":  table,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"check":  true,
			"isTime": false,
		})
	}
}

//CourseSelect 选课退课
func CourseSelect(c *gin.Context) {
	var data = model.CourseSelectIn{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	flag, err := model.IsCourseSelectTime()
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	if flag {
		success, err := model.CourseSelect(data.Account, data.CourseID, data.IsChosed)
		if err != nil {
			fmt.Println("err: ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"check":   false,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check":   true,
			"isTime":  true,
			"success": success,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"check":  true,
			"isTime": false,
		})
	}
}

func GetAllStu(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	//先检验账号和role是否对应
	if !model.CheckAccountAndRole(data) {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	allStu, class, tea, teaIDs, err := model.GetAllStu()
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"check":  true,
		"allStu": allStu,
		"class":  class,
		"tea":    tea,
		"teaIDs": teaIDs,
	})
}

func GetAllTea(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	//先检验账号和role是否对应
	if !model.CheckAccountAndRole(data) {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	allTea, class, err := model.GetAllTeaS()
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"check":  true,
		"allTea": allTea,
		"class":  class,
	})
}

func ChangeClassHeadmaster(c *gin.Context) {
	var data = model.ChangeCHS{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	flag, err := model.ChangeCH(data.ClassID, data.TeaID)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   flag,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"check": true,
	})
}

func GetCourseInfo(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	//先检验账号和role是否对应
	if !model.CheckAccountAndRole(data) {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	allCourse, teaMap, err := model.GetAllCourseInfo()
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"check":     true,
		"allCourse": allCourse,
		"teaMap":    teaMap,
	})

}

func ChangeGroup(c *gin.Context) {
	var data = model.ChangeGroup{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	if data.Acction == 1 {
		flag, err := model.DelGroup(data.CourseID, data.TeaID)
		if err != nil {
			fmt.Println("err: ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"check":   flag,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check": true,
		})
	} else {
		flag, err := model.AddGroup(data.CourseID, data.TeaID)
		if err != nil {
			fmt.Println("err: ", err.Error())
			c.JSON(http.StatusOK, gin.H{
				"check":   flag,
				"message": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"check": true,
		})
	}
}

func GradesRecordInit(c *gin.Context) {
	var data = model.AccountAndRole{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	//先检验账号和role是否对应
	if !model.CheckAccountAndRole(data) {
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	allCourse, courseMap, err := model.GetAllCourseMark(data.Account)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"check":     true,
		"allCourse": allCourse,
		"courseMap": courseMap,
	})
}

func CommitGrades(c *gin.Context) {
	var data = model.UpdateGradesStruct{}
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	flag, err := model.CommitGradesUpdate(data.AllCourse)
	if err != nil {
		fmt.Println("err: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"check":   false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"check": flag,
	})

}
