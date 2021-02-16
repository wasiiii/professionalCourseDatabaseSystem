package courseschedalgorithm

import (
	"fmt"
	"strconv"

	"../gin/db"
	"../testdata"
)

//CSAData 排课算法使用的数据结构体
type CSAData struct {
	CourseID  int
	TeaID     []int
	ClassID   []string
	BeginTime int
	EndTime   int
	IsEvening int
}

//getClass 将class拆分成切片，一个班级9个字节，例如"计科184"
func getClass(classID string) []string {
	var retClass []string
	for i := 0; i < len(classID); i += 9 {
		retClass = append(retClass, classID[i:i+9])
	}
	return retClass
}

//DataPreprocessing 将数据处理成排课算法能用的格式，传进来的数据确保正确能用
func dataPreprocessing(vs []testdata.Variable) []CSAData {
	var datas []CSAData
	for _, v := range vs {
		var data CSAData
		data.CourseID = v.CourseID
		data.TeaID = v.TeaID
		data.ClassID = getClass(v.ClassID[0])
		//这里的时间范围当作是最优，之后如果不是记得调整
		data.BeginTime = v.BeginTime
		data.EndTime = v.EndTime
		//1节/每周
		if v.ClassNumPerWeek == 1 {
			var weekNum int = v.EndTime - v.BeginTime + 1
			if weekNum*2 >= v.CourseTime {
				data.IsEvening = 0
			} else {
				data.IsEvening = 1
			}
			datas = append(datas, data)
		} else {
			data.IsEvening = 0
			datas = append(datas, data)
			datas = append(datas, data)
		}
	}

	return datas
}

//isConflictSection 判断当天当节课教师、班级、时间是否冲突
func isConflictSection(existingCourse CSAData, course CSAData) bool {
	//判断结课时间
	if course.BeginTime > existingCourse.EndTime {
		return false
	}
	//判断教师
	for _, existingTea := range existingCourse.TeaID {
		for _, tea := range course.TeaID {
			if tea == existingTea {
				return true
			}
		}
	}
	//判断班级
	for _, existingClass := range existingCourse.ClassID {
		for _, class := range course.ClassID {
			if class == existingClass {
				return true
			}
		}
	}

	return false
}

//isExistThisSection 判断当天当节课和要插入的课是否冲突，调用isConflictSection
func isExistThisSection(data []CSAData, course CSAData) bool {
	for _, existingCourse := range data {
		if isConflictSection(existingCourse, course) {
			return true
		}
	}

	return false
}

//isConflictToday 当天同班不能上同一门课，即一个班一天不能上相同的课
func isConflictToday(existingCourse CSAData, course CSAData) bool {
	if existingCourse.CourseID == course.CourseID {
		for _, existingClass := range existingCourse.ClassID {
			for _, class := range course.ClassID {
				if class == existingClass {
					return true
				}
			}
		}
	}

	return false
}

//isExistSameClassAndCourseToday 判断当天是否出现冲突的课，调用isConflictToday
func isExistSameClassAndCourseToday(datas [5][]CSAData, course CSAData) bool {
	for _, data := range datas {
		for _, existingCourse := range data {
			if isConflictToday(existingCourse, course) {
				return true
			}
		}
	}

	return false
}

//canSet 判断课表当天是否能插入待排的课
func canSet(datas [5][]CSAData, course CSAData, section int) bool {
	if isExistThisSection(datas[section], course) {
		return false
	}
	if isExistSameClassAndCourseToday(datas, course) {
		return false
	}
	return true
}

//getOriginalTimeTable 获取初始课表
func getOriginalTimeTable(datas []CSAData) [5][5][]CSAData {
	var originalTimeTable [5][5][]CSAData
	for _, data := range datas {
		if data.IsEvening == 1 {
			//var isSet bool = false
			for i := 0; i < 5; i++ {
				if canSet(originalTimeTable[i], data, 4) {
					originalTimeTable[i][4] = append(originalTimeTable[i][4], data)
					//isSet = true
					break
				}
			}
		} else {
			for i := 0; i < 5; i++ {
				var isSet bool = false
				for j := 0; j < 4; j++ {
					if canSet(originalTimeTable[i], data, j) {
						originalTimeTable[i][j] = append(originalTimeTable[i][j], data)
						isSet = true
						break
					}
				}
				if isSet {
					break
				}
			}
		}
	}

	return originalTimeTable
}

func printCSAData(datas []CSAData) {
	for _, d := range datas {
		fmt.Println(d.CourseID)
		for _, t := range d.TeaID {
			fmt.Println(t)
		}
		for _, t := range d.ClassID {
			fmt.Println(t)
		}
		fmt.Println(d.BeginTime)
		fmt.Println(d.EndTime)
		fmt.Println(d.IsEvening)
		fmt.Println()
	}
}

//showTable 打印课表
func showTable(originalTimeTable [5][5][]CSAData) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("星期", i+1, "第", j+1, "大节")
			printCSAData(originalTimeTable[i][j])
		}
	}
}

//CSA 排课算法
func CSA() {
	var datas []CSAData = dataPreprocessing(testdata.Vs)
	//贪心获取一个初始课表
	var originalTimeTable [5][5][]CSAData = getOriginalTimeTable(datas)
	//showTable(originalTimeTable)
	var Datass [5][5][]db.TestS = makeTest(originalTimeTable)
	db.InitDB()
	db.InsertTest(Datass)
	//遗传算法优化
}

func makeTest(originalTimeTable [5][5][]CSAData) [5][5][]db.TestS {
	var datass [5][5][]db.TestS
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("星期", i+1, "第", j+1, "大节")
			printCSAData(originalTimeTable[i][j])
			for _, d := range originalTimeTable[i][j] {
				var ty db.TestS
				ty.Course_ID = strconv.Itoa(d.CourseID)
				ty.Week = i + 1
				ty.Section = j + 1
				for _, t := range d.TeaID {
					ty.Tea_ID += strconv.Itoa(t)
				}
				for _, t := range d.ClassID {
					ty.Class_ID += t
				}
				ty.Begin_Time = d.BeginTime
				ty.End_Time = d.EndTime
				datass[i][j] = append(datass[i][j], ty)
			}
		}

	}
	return datass
}
