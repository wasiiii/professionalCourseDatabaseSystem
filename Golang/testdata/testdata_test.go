package testdata

import (
	"fmt"
	"testing"
)

//打印Variable
func PrintVariable(v Variable) {
	fmt.Println("courseID: ", v.CourseID)
	fmt.Println("teaID: ", v.TeaID)
	fmt.Println("classID: ", v.ClassID)
	fmt.Println("courseTime: ", v.CourseTime)
	fmt.Println("beginTime: ", v.BeginTime)
	fmt.Println("endTime: ", v.EndTime)
	fmt.Println("classNumPerWeek: ", v.ClassNumPerWeek)
	fmt.Println()
}

func Test(t *testing.T) {
	for _, v := range Vs {
		PrintVariable(v)
	}
}
