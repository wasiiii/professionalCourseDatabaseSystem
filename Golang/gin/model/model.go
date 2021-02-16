package model

import (
	"database/sql"
	"fmt"
	"sort"

	"../db"
)

//Stu 学生结构体
type Stu struct {
	StuID           string  `json:stu_id`
	StuName         string  `json:stu_name`
	StuSex          string  `json:stu_sex`
	StuClass        string  `json:stu_class`
	Password        string  `json:password`
	ClassHeadmaster string  `json:class_headmaster`
	Gpa             float64 `json:gpa`
}

//CheckStu 学生登录检查密码
func CheckStu(stu Stu) bool {
	password, err := db.SelectStuPassword(stu.StuID)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return false
	}
	return password == stu.Password
}

//Tea 教师结构体
type Tea struct {
	TeaID    string   `json:tea_id`
	TeaName  string   `json:tea_name`
	TeaType  int      `json:tea_type`
	TeaClass []string `json:tea_class`
	Password string   `json:password`
}

//CheckTea 教师登录检查密码，获取教师类型，是否为班主任
func CheckTea(tea Tea) (bool, int, int) {
	password, TeaType, isHeadmaster, err := db.SelectTeaPasswordAndType(tea.TeaID)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return false, -1, -1
	}
	return password == tea.Password, TeaType, isHeadmaster
}

type AccountAndRole struct {
	Account string `json:account`
	Role    string `json:role`
}
type AccountAndRoleWhenChangeClassTable struct {
	Account string `json:account`
	Role    string `json:role`
	Class   string `json:class`
}

func CheckAccountAndRole(aar AccountAndRole) bool {
	return true
}

//GetStuInfo 获取学生个人信息
func GetStuInfo(stuID string) (Stu, int) {
	stuDB, err := db.SelectStuInfo(stuID)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return Stu{}, -1
	}
	stu := Stu{
		StuID:           stuDB.Stu_ID,
		StuName:         stuDB.Stu_name,
		StuSex:          stuDB.Stu_sex,
		StuClass:        stuDB.Stu_class,
		ClassHeadmaster: stuDB.Class_headmaster,
		Gpa:             stuDB.Gpa,
	}

	//求所有学生的ID和绩点
	allStu, err := db.SelectAllStuIDAndGpa()
	if err != nil {
		fmt.Print("err: ", err)
		return Stu{}, -1
	}
	//两个map
	//这个map，stuID映射Gpa
	var stuIDMapGpa map[string]float64
	stuIDMapGpa = make(map[string]float64)
	for _, aS := range allStu {
		stuIDMapGpa[aS.Stu_ID] = aS.Gpa
	}
	//这个map是为了清除相同的gpa
	var gpaMapRank map[float64]int
	//切片为了排序key
	var gpaS []float64
	gpaMapRank = make(map[float64]int)
	for _, aS := range allStu {
		gpaMapRank[aS.Gpa] = 1
	}
	for k := range gpaMapRank {
		gpaS = append(gpaS, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(gpaS)))

	var rank int
	for i, r := range gpaS {
		if r == stuIDMapGpa[stuID] {
			rank = i + 1
			break
		}
	}

	return stu, rank
}

//GetTeaInfo 获取教师个人信息
func GetTeaInfo(teaID string) Tea {
	teaDB, class, err := db.SelectTeaInfo(teaID)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return Tea{}
	}
	tea := Tea{
		TeaID:    teaDB.Tea_ID,
		TeaName:  teaDB.Tea_name,
		TeaClass: class,
	}

	return tea
}

type Table struct {
	Course_ID   string `json:course_id`
	Course_name string `json:course_name`
	Tea_name    string `json:tea_name`
	Begin_Time  int    `json:begin_time`
	End_Time    int    `json:end_time`
	Day         int    `json:day`
	Section     int    `json:section`
	Class_id    string `json:class_id`
}

//根据教师（班主任）id获取班级id，因为一个老师可以任多个班主任，所以用[]string类型
func getTeaClass(teaID string) ([]string, error) {
	class, err := db.SelectTeaClass(teaID)
	if err != nil {
		return class, err
	}
	return class, nil
}

//根据学生id获取所选课程id，即从选课表中获得
func getStuCourse(stuID string) ([]string, error) {
	var course []string
	course, err := db.SelectStuCourse(stuID)
	if err != nil {
		return course, err
	}
	return course, nil
}

//GetStuTable 获取学生课表
func GetStuTable(stuID string) ([]Table, error) {
	var table []Table
	//获取所属班级
	class, err := db.SelectStuClass(stuID)
	if err != nil {
		return table, err
	}

	//获取所选课程
	courses, err := getStuCourse(stuID)
	if err != nil {
		return table, err
	}

	//用map，课程id映射课程名字
	courseIDMapName, err := db.SelectCourseName(courses)
	if err != nil {
		return table, err
	}

	//从排课表中取得学生课表
	tableDB, err := db.SelectStuTable(class, courses)
	if err != nil {
		return table, err
	}

	//优化教师显示，将多个教师的字符串拆分成[]string
	var teaID []string
	for _, t := range tableDB {
		for jj := 0; jj < len(t.Tea_ID); jj += 7 {
			teaID = append(teaID, t.Tea_ID[jj:jj+6])
		}
	}

	//用map，教师id映射教师名字
	teaIDMapName, err := db.SelectTeaName(teaID)

	//构造可以返回前端的形式
	for _, t := range tableDB {
		var nameStr string
		for jj := 0; jj < len(t.Tea_ID); jj += 7 {
			if jj == 0 {
				nameStr += teaIDMapName[t.Tea_ID[jj:jj+6]]
			} else {
				nameStr += "," + teaIDMapName[t.Tea_ID[jj:jj+6]]
			}
		}

		tableP := Table{
			Course_ID:   t.Course_ID,
			Course_name: courseIDMapName[t.Course_ID],
			Tea_name:    nameStr,
			Begin_Time:  t.Begin_time,
			End_Time:    t.End_time,
			Day:         t.Day,
			Section:     t.Section,
		}
		table = append(table, tableP)
	}

	return table, nil
}

type temporary struct {
	Course_ID string
	Class_ID  string
}

//GetTeaTable 获取教师课表
func GetTeaTable(teaID string) ([]Table, error) {
	var table []Table

	//获取排课表中该教师的课
	courses, err := db.SelectTeaCourse(teaID)
	if err != nil {
		return table, err
	}
	//获取map，课程id映射课程名字
	courseIDMapName, err := db.SelectCourseName(courses)
	if err != nil {
		return table, err
	}

	//获取排课表中该教师的课的所有信息
	tableDB, err := db.SelectTeaAllTable(teaID)
	if err != nil {
		return table, err
	}

	//构造成一个切片
	for _, t := range tableDB {
		tableP := Table{
			Course_ID:   t.Course_ID,
			Course_name: courseIDMapName[t.Course_ID],
			Begin_Time:  t.Begin_time,
			End_Time:    t.End_time,
			Day:         t.Day,
			Section:     t.Section,
			Class_id:    t.Class_ID,
		}
		table = append(table, tableP)
	}

	//获取map，课程id+班级id对应一个课在排课表中的所有信息
	var ccMapTableDB map[temporary][]Table
	ccMapTableDB = make(map[temporary][]Table)
	for _, t := range table {
		var tp temporary
		tp.Course_ID = t.Course_ID
		for ii := 0; ii < len(t.Class_id); ii += 10 {
			tp.Class_ID = t.Class_id[ii : ii+9]
			ccMapTableDB[tp] = append(ccMapTableDB[tp], t)
		}
	}

	//清空切片
	table = table[0:0]

	//从选课表中获取所有被选的课
	courseChosed, err := db.SelectCourseAndClassForTea()
	if err != nil {
		return table, err
	}
	//构造成map的key
	var key []temporary
	for _, t := range courseChosed {
		kk := temporary{
			Course_ID: t.Course_id,
			Class_ID:  t.Class_id,
		}
		key = append(key, kk)
	}

	//用key在map中找到对应的value，构造返回前端的形式
	var tableDBMapInsert map[Table]int
	tableDBMapInsert = make(map[Table]int)
	for _, k := range key {
		t, ok1 := ccMapTableDB[k]
		for _, tt := range t {
			_, ok2 := tableDBMapInsert[tt]
			if ok1 && !ok2 {
				tableDBMapInsert[tt] = 1
				table = append(table, tt)
			}
		}
	}

	return table, nil
}

//GetClassTableFromStu 学生获取班级课表
func GetClassTableFromStu(stuID string) ([]Table, error) {
	var table []Table
	class, err := db.SelectStuClass(stuID)
	if err != nil {
		return table, err
	}

	return getTableFromClass(class)
}

//GetClassTableFromTea 教师（班主任）获取班级课表，初始化
func GetClassTableFromTea(teaID string) ([]Table, []string, error) {
	var ret []Table
	class, err := getTeaClass(teaID)
	if err != nil {
		return ret, class, err
	}

	ret, err = getTableFromClass(class[0])
	if err != nil {
		return ret, class, err
	}

	return ret, class, nil
}

//getTableFromClass 根据班级ID获取课表
func getTableFromClass(classID string) ([]Table, error) {
	var table []Table

	//从排课表中获取班级ID对应课程
	courses, err := db.SelectClassCourse(classID)
	if err != nil {
		return table, err
	}
	//用map，课程id映射课程名字
	courseIDMapName, err := db.SelectCourseName(courses)
	if err != nil {
		return table, err
	}

	//从排课表中取得学生课表
	tableDB, err := db.SelectStuTable(classID, courses)
	if err != nil {
		return table, err
	}

	//优化教师显示，将多个教师的字符串拆分成[]string
	var teaID []string
	for _, t := range tableDB {
		for jj := 0; jj < len(t.Tea_ID); jj += 7 {
			teaID = append(teaID, t.Tea_ID[jj:jj+6])
		}
	}

	//用map，教师id映射教师名字
	teaIDMapName, err := db.SelectTeaName(teaID)

	//构造成一个切片
	for _, t := range tableDB {
		var nameStr string
		for jj := 0; jj < len(t.Tea_ID); jj += 7 {
			if jj == 0 {
				nameStr += teaIDMapName[t.Tea_ID[jj:jj+6]]
			} else {
				nameStr += "," + teaIDMapName[t.Tea_ID[jj:jj+6]]
			}
		}

		tableP := Table{
			Class_id:    classID,
			Course_ID:   t.Course_ID,
			Course_name: courseIDMapName[t.Course_ID],
			Tea_name:    nameStr,
			Begin_Time:  t.Begin_time,
			End_Time:    t.End_time,
			Day:         t.Day,
			Section:     t.Section,
		}
		table = append(table, tableP)
	}

	//获取map，课程id+班级id对应一个课在排课表中的所有信息
	var ccMapTableDB map[temporary][]Table
	ccMapTableDB = make(map[temporary][]Table)
	for _, t := range table {
		var tp temporary
		tp.Course_ID = t.Course_ID
		for ii := 0; ii < len(t.Class_id); ii += 10 {
			tp.Class_ID = t.Class_id[ii : ii+9]
			ccMapTableDB[tp] = append(ccMapTableDB[tp], t)
		}
	}

	//清空切片
	table = table[0:0]

	//从选课表中获取所有被选的课
	courseChosed, err := db.SelectCourseAndClassForTea()
	if err != nil {
		return table, err
	}
	//构造成map的key
	var key []temporary
	for _, t := range courseChosed {
		kk := temporary{
			Course_ID: t.Course_id,
			Class_ID:  t.Class_id,
		}
		key = append(key, kk)
	}

	//用key在map中找到对应的value，构造返回前端的形式
	var tableDBMapInsert map[Table]int
	tableDBMapInsert = make(map[Table]int)
	for _, k := range key {
		t, ok1 := ccMapTableDB[k]
		for _, tt := range t {
			_, ok2 := tableDBMapInsert[tt]
			if ok1 && !ok2 {
				tableDBMapInsert[tt] = 1
				table = append(table, tt)
			}
		}
	}

	return table, nil
}

//GetClassTableFromTeaChange 教师（班主任）获取特定班级课表
func GetClassTableFromTeaChange(classID string) ([]Table, error) {
	return getTableFromClass(classID)
}

//IsCourseSelectTime 判断是否为选课阶段
func IsCourseSelectTime() (bool, error) {
	begin, err := db.SelectBeginOrEnd()
	if err != nil {
		return false, err
	}
	if begin == 1 {
		return true, nil
	}
	return false, nil
}

//CourseWaitSelect 待选课程信息结构体
type CourseWaitSelect struct {
	CourseID   string
	CourseName string
	BeginTime  int
	EndTime    int
	Day        int
	Section    int
	TeaName    string
	StuNum     int
	IsChosed   bool
}

//GetCourseWaitSelectTable 获取待选课程信息
func GetCourseWaitSelectTable(stuID string) ([]CourseWaitSelect, error) {
	var ret []CourseWaitSelect

	//学生ID对应班级ID
	classID, err := db.SelectStuClass(stuID)
	if err != nil {
		return ret, err
	}

	chosedCourseID, err := db.SelectStuCourse(stuID)
	var courseIDMapChosed map[string]bool
	courseIDMapChosed = make(map[string]bool)
	for _, c := range chosedCourseID {
		courseIDMapChosed[c] = true
	}

	//班级ID在待选课程表中的课程信息
	courseWs, err := db.SelectCourseFromClassOnWaitSelect(classID)
	if err != nil {
		return ret, err
	}

	//根据课程信息选出对应的课程ID切片
	var courseIDs []string
	for _, c := range courseWs {
		courseIDs = append(courseIDs, c.Course_ID)
	}

	//构造map，课程ID映射选课人数
	var courseIDMapNum map[string]int
	courseIDMapNum = make(map[string]int)
	for _, c := range courseWs {
		courseIDMapNum[c.Course_ID] = c.Stu_num
	}

	//用map，课程ID映射课程名字
	courseIDMapName, err := db.SelectCourseName(courseIDs)
	if err != nil {
		return ret, err
	}

	var courses []db.Table
	for _, c := range courseIDs {
		coursesP, err := db.SelectCourseFromTimetable(classID, c)
		if err != nil {
			return ret, err
		}
		courses = append(courses, coursesP...)
	}

	//优化教师显示，将多个教师的字符串拆分成[]string
	var teaID []string
	for _, t := range courses {
		for jj := 0; jj < len(t.Tea_ID); jj += 7 {
			teaID = append(teaID, t.Tea_ID[jj:jj+6])
		}
	}

	//用map，教师id映射教师名字
	teaIDMapName, err := db.SelectTeaName(teaID)

	for _, c := range courses {
		var nameStr string
		for jj := 0; jj < len(c.Tea_ID); jj += 7 {
			if jj == 0 {
				nameStr += teaIDMapName[c.Tea_ID[jj:jj+6]]
			} else {
				nameStr += "," + teaIDMapName[c.Tea_ID[jj:jj+6]]
			}
		}
		retP := CourseWaitSelect{
			CourseID:   c.Course_ID,
			TeaName:    nameStr,
			CourseName: courseIDMapName[c.Course_ID],
			BeginTime:  c.Begin_time,
			EndTime:    c.End_time,
			Day:        c.Day,
			Section:    c.Section,
			StuNum:     courseIDMapNum[c.Course_ID],
			IsChosed:   courseIDMapChosed[c.Course_ID],
		}
		ret = append(ret, retP)
	}

	return ret, nil
}

//CourseSelectIn 接收选课的前端信息结构体
type CourseSelectIn struct {
	Account  string `json:account`
	Role     string `json:role`
	CourseID string `json:courseID`
	IsChosed bool   `json:isChosed`
}

//CourseSelect 选课退课
func CourseSelect(stuID string, courseID string, isChosed bool) (bool, error) {
	var flag bool
	var err error
	if isChosed {
		flag, err = drop(stuID, courseID)
	} else {
		flag, err = choose(stuID, courseID)
	}
	return flag, err
}

//选课
func choose(stuID string, courseID string) (bool, error) {
	//选择学生班级
	classID, err := db.SelectStuClass(stuID)
	if err != nil {
		return false, err
	}
	//从排课表中选出这门课和学生班级对应的所有信息
	courseTable, err := db.SelectCourseFromTimetableDiv(classID, courseID)
	if err != nil {
		return false, err
	}
	//从排课表中选出这门课对应的所有班级
	classIDs, err := db.SelectClassFromTimetable(courseTable[0])
	if err != nil {
		return false, err
	}
	//开始事务
	err = db.Start()
	if err != nil {
		return false, err
	}
	//插入选课表
	err = db.InsertCourseSelect(stuID, courseID)
	if err != nil {
		db.Rollback()
		return false, err
	}
	//更新待选课表
	for _, c := range classIDs {
		err = db.UpdateCourseWaitSelectAdd(c, courseID)
		if err != nil {
			db.Rollback()
			return false, err
		}
	}
	//提交事务
	err = db.Commit()
	if err != nil {
		db.Rollback()
		return false, err
	}

	return true, nil
}

//退课
func drop(stuID string, courseID string) (bool, error) {
	//选择学生班级
	classID, err := db.SelectStuClass(stuID)
	if err != nil {
		return false, err
	}
	//从排课表中选出这门课和学生班级对应的所有信息
	courseTable, err := db.SelectCourseFromTimetableDiv(classID, courseID)
	if err != nil {
		return false, err
	}
	//从排课表中选出这门课对应的所有班级
	classIDs, err := db.SelectClassFromTimetable(courseTable[0])
	if err != nil {
		return false, err
	}
	//开始事务
	err = db.Start()
	if err != nil {
		return false, err
	}
	//从选课表中删除
	err = db.DeleteCourseSelect(stuID, courseID)
	if err != nil {
		db.Rollback()
		return false, err
	}
	//更新待选课表
	for _, c := range classIDs {
		err = db.UpdateCourseWaitSelectSub(c, courseID)
		if err != nil {
			db.Rollback()
			return false, err
		}
	}
	//提交事务
	err = db.Commit()
	if err != nil {
		db.Rollback()
		return false, err
	}

	return true, nil
}

type AllStu struct {
	StuID    string  `json:stu_id`
	StuName  string  `json:stu_name`
	StuSex   string  `json:stu_sex`
	Gpa      float64 `json:gpa`
	StuClass string  `json:stu_class`
}

type Class struct {
	ClassID         string `json:class_id`
	ClassHeadmaster string `json:class_headmaster`
}

func GetAllStu() ([]AllStu, map[string]string, map[string]string, []string, error) {
	var ret []AllStu
	var ret2 map[string]string
	var ret3 map[string]string
	var ret4 []string
	classInfo, err := db.SelectClass()
	if err != nil {
		return ret, ret2, ret3, ret4, err
	}
	ret2 = make(map[string]string)
	for _, cI := range classInfo {
		ret2[cI.Class_id] = cI.Class_headmaster
	}

	stuS, err := db.SelectAllStu()
	if err != nil {
		return ret, ret2, ret3, ret4, err
	}
	for _, sS := range stuS {
		aS := AllStu{
			StuID:    sS.Stu_ID,
			StuName:  sS.Stu_name,
			StuSex:   sS.Stu_sex,
			Gpa:      sS.Gpa,
			StuClass: sS.Stu_class,
		}
		ret = append(ret, aS)
	}

	ret4, err = db.SelectAllTeaID()
	if err != nil {
		return ret, ret2, ret3, ret4, err
	}

	ret3, err = db.SelectTeaName(ret4)
	if err != nil {
		return ret, ret2, ret3, ret4, err
	}

	return ret, ret2, ret3, ret4, nil
}

type AllTea struct {
	TeaID   string   `json:tea_id`
	TeaName string   `json:tea_name`
	Course  []string `json:course`
}

func GetAllTeaS() ([]AllTea, map[string]string, error) {
	var ret []AllTea
	var ret2 map[string]string
	ret2 = make(map[string]string)

	tea, err := db.SelectAllTea()
	if err != nil {
		return ret, ret2, err
	}

	var teaIDMapStruct map[string]AllTea
	teaIDMapStruct = make(map[string]AllTea)
	for _, t := range tea {
		tt := AllTea{
			TeaID:   t.Tea_ID,
			TeaName: t.Tea_name,
		}
		teaIDMapStruct[tt.TeaID] = tt
	}

	var courses []string
	for i, t := range teaIDMapStruct {
		c, err := db.SelectCourseFromTeaIDOnCG(i)
		if err != nil {
			return ret, ret2, err
		}
		t.Course = c
		courses = append(courses, c...)
		ret = append(ret, t)
	}

	ret2, err = db.SelectCourseName(courses)
	if err != nil {
		return ret, ret2, err
	}

	return ret, ret2, nil
}

type ChangeCHS struct {
	Account string `json:account`
	Role    string `json:role`
	ClassID string `json:classID`
	TeaID   string `json:teaID`
}

type ChangeGroup struct {
	Account  string `json:account`
	Role     string `json:role`
	TeaID    string `json:teaID`
	CourseID string `json:courseID`
	Acction  int    `json:acction`
}

func ChangeCH(classID string, teaID string) (bool, error) {
	err := db.UpdateClassHeadmaster(classID, teaID)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

type AllCourse struct {
	Course_ID   string   `json:course_id`
	Course_name string   `json:course_name`
	Group       []string `json:group`
}

func GetAllCourseInfo() ([]AllCourse, map[string]string, error) {
	var ret []AllCourse
	var ret2 map[string]string
	ret2 = make(map[string]string)

	courses, err := db.SelectAllCourses()
	if err != nil {
		return ret, ret2, err
	}

	for _, c := range courses {
		aC := AllCourse{
			Course_ID:   c.Course_ID,
			Course_name: c.Course_name,
		}
		aC.Group, err = db.SelectTeaFromCourseIDOnGroup(aC.Course_ID)
		if err != nil {
			return ret, ret2, err
		}
		ret = append(ret, aC)
	}

	teaIDs, err := db.SelectAllTeaID()
	if err != nil {
		return ret, ret2, err
	}
	ret2, err = db.SelectTeaName(teaIDs)
	if err != nil {
		return ret, ret2, err
	}

	return ret, ret2, nil
}

func DelGroup(courseID string, teaID string) (bool, error) {
	err := db.DeleteGroup(courseID, teaID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddGroup(courseID string, teaID string) (bool, error) {
	err := db.AddGroup(courseID, teaID)
	if err != nil {
		return false, err
	}
	return true, nil
}

type GradesRecordStruct struct {
	CourseID   string        `json:course_id`
	CourseName string        `json:course_name`
	StuID      string        `json:stu_id`
	StuName    string        `json:stu_name`
	Grade      sql.NullInt32 `json:grade`
}

func GetAllCourseMark(teaID string) ([]GradesRecordStruct, map[string]string, error) {
	var ret []GradesRecordStruct
	var ret2 map[string]string
	ret2 = make(map[string]string)

	courseIDAndClsID, err := db.SelectCCOnTimetable(teaID)
	if err != nil {
		return ret, ret2, err
	}

	var courses []string
	var class []string
	for _, cac := range courseIDAndClsID {
		courses = append(courses, cac.Course_ID)
		class = append(class, cac.Class_ID)
	}

	//用map，课程id映射课程名字
	ret2, err = db.SelectCourseName(courses)
	if err != nil {
		return ret, ret2, err
	}

	stuIDMapName, err := db.SelectAllStuMap()
	if err != nil {
		return ret, ret2, err
	}
	var grs []db.CCStu
	for _, cac := range courseIDAndClsID {
		save, err := db.SelectAllStuGrades(cac)
		if err != nil {
			return ret, ret2, err
		}
		grs = append(grs, save...)
	}

	for _, g := range grs {
		gg := GradesRecordStruct{
			StuID:      g.Stu_ID,
			StuName:    stuIDMapName[g.Stu_ID],
			CourseID:   g.Course_ID,
			CourseName: ret2[g.Course_ID],
			Grade:      g.Mark,
		}
		ret = append(ret, gg)
	}
	return ret, ret2, nil
}

type UpdateGradesStruct struct {
	Account   string                `json:account`
	Role      string                `json:role`
	AllCourse []UpdateGradesStructP `json:allCourse`
}

type UpdateGradesStructP struct {
	CourseID   string        `json:course_id`
	CourseName string        `json:course_name`
	StuID      string        `json:stu_id`
	StuName    string        `json:stu_name`
	Grade      sql.NullInt32 `json:grade`
	NewGrade   int           `json:newGrade`
}

func CommitGradesUpdate(allCourse []UpdateGradesStructP) (bool, error) {
	db.Start()

	//更新选课表中成绩
	for _, aC := range allCourse {
		err := db.UpdateGrades(aC.CourseID, aC.StuID, aC.NewGrade)
		if err != nil {
			db.Rollback()
			return false, err
		}
	}

	//更新绩点
	var stuIDs []string
	for _, aC := range allCourse {
		stuIDs = append(stuIDs, aC.StuID)
	}
	for _, stuID := range stuIDs {
		var gpa float64
		up := 0
		down := 0
		gg, err := db.SelectGG(stuID)
		if err != nil {
			db.Rollback()
			return false, err
		}
		for _, g := range gg {
			if !g.Mark.Valid {
				continue
			}
			newG := int(g.Mark.Int32)
			up += (newG * g.Course_credit)
			down += g.Course_credit
		}
		gpa = float64(up) / float64(down)
		gpa = (gpa - 60) / 10
		if gpa < 0 {
			gpa = 0
		}
		err = db.UpdateGpa(gpa, stuID)
		if err != nil {
			db.Rollback()
			return false, err
		}
	}
	db.Commit()

	return true, nil
}
