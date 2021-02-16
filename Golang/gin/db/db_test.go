package db

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	stu, err := SelectStuInfo("1806100202")
	fmt.Println(111)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	fmt.Println(stu.Stu_ID)
	fmt.Println(stu.Stu_name)
	//fmt.Println(stu.Gpa)
	//fmt.Println(stu.Class_headmaster)
}
