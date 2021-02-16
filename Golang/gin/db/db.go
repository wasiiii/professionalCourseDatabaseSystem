package db

import (
	//mysql的包

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//DB 句柄
var DB *sqlx.DB

//CONN 事务句柄
var CONN *sql.Tx

//Start 事务开始
func Start() error {
	conn, err := DB.Begin()
	if err != nil {
		return err
	}
	CONN = conn
	return nil
}

//Rollback 事务回滚
func Rollback() error {
	return CONN.Rollback()
}

//Commit 事务结束
func Commit() error {
	return CONN.Commit()
}

//InitDB 初始化连接数据库
func InitDB() error {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/pcds")
	if err != nil {
		return err
	}
	err = database.Ping()
	if err != nil {
		return err
	}
	DB = database
	return nil
}

//Stu 学生结构体
type Stu struct {
	Stu_ID           string  `json:stu_id`
	Stu_name         string  `json:stu_name`
	Stu_sex          string  `json:stu_sex`
	Stu_class        string  `json:stu_class`
	Password         string  `json:password`
	Class_headmaster string  `json:class_headmaster`
	Gpa              float64 `json:gpa`
}

//SelectStuPassword 选择学生密码
func SelectStuPassword(stuID string) (string, error) {
	var stu Stu
	err := DB.Get(&stu, "select password from student where stu_id=?", stuID)
	if err != nil {
		return "", err
	}

	return stu.Password, nil
}

//Tea 教师结构体
type Tea struct {
	Tea_ID   string `json:tea_id`
	Tea_name string `json:tea_name`
	Tea_type int    `json:tea_type`
	Password string `json:password`
}

//SelectTeaPasswordAndType 选择密码和老师类型
func SelectTeaPasswordAndType(teaID string) (string, int, int, error) {
	var tea Tea
	err := DB.Get(&tea, "select tea_type, password from teacher where tea_id=?", teaID)
	if err != nil {
		return "", -1, -1, err
	}
	err = DB.Get(&tea, "select class_headmaster as tea_id from class where class_headmaster=?", teaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return tea.Password, tea.Tea_type, 0, nil
		}
		return "", -1, -1, err
	}
	return tea.Password, tea.Tea_type, 1, nil
}

type TestS struct {
	Course_ID  string `json:course_id`
	Tea_ID     string `json:tea_id`
	Class_ID   string `json:class_id`
	Begin_Time int    `json:begin_time`
	End_Time   int    `json:end_time`
	Week       int    `json:day`
	Section    int    `json:section`
}

type Course struct {
	Course_ID     string `json:course_id`
	Course_name   string `json:course_name`
	Course_type   int    `json:course_type`
	Course_credit int    `json:course_credit`
	Course_hour   int    `json:course_hour`
}

type Class struct {
	Class_id         string `json:class_id`
	Class_headmaster string `json:class_headmaster`
	Stu_num          int    `json:stu_num`
}

func InsertTest(Datass [5][5][]TestS) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for _, d := range Datass[i][j] {
				for ii := 0; ii < len(d.Class_ID); ii += 9 {
					for jj := 0; jj < len(d.Tea_ID); jj += 6 {
						_, err := DB.Exec("insert into time_table(course_id, tea_id, class_id, begin_time, end_time, section, day)values(?, ?, ?, ?, ?, ?, ?)", d.Course_ID, d.Tea_ID[jj:jj+6], d.Class_ID[ii:ii+9], d.Begin_Time, d.End_Time, d.Section, d.Week)
						if err != nil {
							fmt.Println("exec failed, ", err.Error())
							return
						}
					}
				}
			}
		}
	}

	//选必修课
	var cs []Course
	err := DB.Select(&cs, "select course_id from course where course_type=?", 0)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	for _, c := range cs {
		fmt.Println(c.Course_ID)
	}
	//选全部班级
	var cls []Class
	err = DB.Select(&cls, "select class_id from class")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	for _, c := range cls {
		fmt.Println(c.Class_id)
	}
	//查找班级对应学生列表，map
	var classMapStu map[string][]Stu
	classMapStu = make(map[string][]Stu)
	for _, cl := range cls {
		var empty []Stu
		err = DB.Select(&empty, "select stu_id from student where stu_class = ?", cl.Class_id)
		if err != nil {
			fmt.Println("err: ", err.Error())
			return
		}
		classMapStu[cl.Class_id] = empty
	}
	for class := range classMapStu {
		for _, stu := range classMapStu[class] {
			fmt.Println(class, ": ", stu.Stu_ID)
		}
	}
	//插入选课表
	for _, c := range cs {
		for class := range classMapStu {
			for _, stu := range classMapStu[class] {
				_, err := DB.Exec("insert into course_select(course_id, stu_id)values(?, ?)", c.Course_ID, stu.Stu_ID)
				if err != nil {
					fmt.Println("exec failed, ", err.Error())
					return
				}
			}
		}
	}

	//插入待选课表
	//选选修课
	cs = cs[0:0]
	err = DB.Select(&cs, "select course_id from course where course_type=?", 1)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	for _, c := range cs {
		fmt.Println(c.Course_ID)
	}

	//从排课表中选出课程和对应班级
	var csAndcls []Table
	err = DB.Select(&csAndcls, "select distinct course_id, class_id from time_table")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}

	var csMapcls map[string][]string
	csMapcls = make(map[string][]string)
	for _, cac := range csAndcls {
		csMapcls[cac.Course_ID] = append(csMapcls[cac.Course_ID], cac.Class_ID)
	}

	//插入待选课表
	for _, c := range cs {
		for _, ci := range csMapcls[c.Course_ID] {
			_, err := DB.Exec("insert into course_wait_for_select(course_id, class_id)values(?,?)", c.Course_ID, ci)
			if err != nil {
				fmt.Println("exec failed, ", err.Error())
				return
			}
		}
	}
}

func SelectStuInfo(stuID string) (Stu, error) {
	var stu Stu
	err := DB.Get(&stu, "select stu_id, stu_name, stu_sex, stu_class, gpa from student where stu_id=?", stuID)
	if err != nil {
		return Stu{}, err
	}

	var cls Class
	err = DB.Get(&cls, "select class_headmaster from class where class_id=?", stu.Stu_class)
	if err != nil {
		if err == sql.ErrNoRows {
			stu.Class_headmaster = "待定"
			return stu, nil
		}
		return Stu{}, err
	}

	var tea Tea
	err = DB.Get(&tea, "select tea_name from teacher where tea_id=?", cls.Class_headmaster)
	if err != nil {
		return Stu{}, err
	}
	stu.Class_headmaster = tea.Tea_name

	return stu, nil
}

func SelectAllStuIDAndGpa() ([]Stu, error) {
	var stus []Stu
	err := DB.Select(&stus, "select stu_id, gpa from student")
	return stus, err
}

func SelectTeaInfo(teaID string) (Tea, []string, error) {
	var tea Tea
	var class []string
	err := DB.Get(&tea, "select tea_id, tea_name from teacher where tea_id=?", teaID)
	if err != nil {
		return Tea{}, class, err
	}

	err = DB.Select(&class, "select class_id from class where class_headmaster=?", teaID)
	if err != nil && err != sql.ErrNoRows {
		return Tea{}, class, err
	}

	return tea, class, nil
}

func SelectStuClass(stuID string) (string, error) {
	var stu Stu
	err := DB.Get(&stu, "select stu_class from student where stu_id=?", stuID)
	if err != nil {
		return "", err
	}
	return stu.Stu_class, nil
}

func SelectStuCourse(stuID string) ([]string, error) {
	var courseS []Course
	var courses []string
	err := DB.Select(&courseS, "select course_id from course_select where stu_id=?", stuID)
	if err != nil {
		return courses, err
	}
	for _, c := range courseS {
		courses = append(courses, c.Course_ID)
	}
	return courses, nil
}

func SelectCourseName(courses []string) (map[string]string, error) {
	var courseIDMapName map[string]string
	courseIDMapName = make(map[string]string)
	for _, c := range courses {
		var course Course
		err := DB.Get(&course, "select course_name from course where course_id=?", c)
		if err != nil {
			return courseIDMapName, err
		}
		courseIDMapName[c] = course.Course_name
	}
	return courseIDMapName, nil
}

func SelectTeaName(teaIDs []string) (map[string]string, error) {
	var teaIDMapName map[string]string
	teaIDMapName = make(map[string]string)
	for _, t := range teaIDs {
		var tea Tea
		err := DB.Get(&tea, "select tea_name from teacher where tea_id=?", t)
		if err != nil {
			return teaIDMapName, err
		}
		teaIDMapName[t] = tea.Tea_name
	}
	return teaIDMapName, nil
}

type Table struct {
	Course_ID  string `json:course_id`
	Tea_ID     string `json:tea_id`
	Begin_time int    `json:begin_time`
	End_time   int    `json:end_time`
	Day        int    `json:day`
	Section    int    `json:section`
	Class_ID   string `json:class_id`
}

func SelectStuTable(class string, courses []string) ([]Table, error) {
	var stuTable []Table
	for _, c := range courses {
		var stuTableP []Table
		err := DB.Select(&stuTableP, "select course_id, group_concat(tea_id) as tea_id, begin_time, end_time, day, section from time_table where class_id=? and course_id=? group by course_id, begin_time, end_time, day, section", class, c)
		if err != nil {
			return stuTable, err
		}
		stuTable = append(stuTable, stuTableP...)
	}
	return stuTable, nil
}

func SelectTeaCourse(teaID string) ([]string, error) {
	var courseS []Table
	var courses []string
	err := DB.Select(&courseS, "select distinct course_id from time_table where tea_id=?", teaID)
	if err != nil {
		return courses, err
	}
	for _, c := range courseS {
		courses = append(courses, c.Course_ID)
	}
	return courses, nil
}

func SelectTeaAllTable(teaID string) ([]Table, error) {
	var teaAllTable []Table
	err := DB.Select(&teaAllTable, "select course_id, group_concat(class_id) as class_id, begin_time, end_time, day, section from time_table where tea_id=? group by course_id, begin_time, end_time, day, section", teaID)
	if err != nil {
		return teaAllTable, err
	}
	return teaAllTable, nil
}

type Temporary struct {
	Course_id string `json:course_id`
	Class_id  string `json:class_id`
}

func SelectCourseAndClassForTea() ([]Temporary, error) {
	var ret []Temporary
	err := DB.Select(&ret, "select distinct course_id, stu_class as class_id from course_select join student on course_select.stu_id = student.stu_id")
	return ret, err
}

//SelectTeaClass 根据教师（班主任)ID选出班级ID
func SelectTeaClass(teaID string) ([]string, error) {
	var ret []string
	var classS []Class
	err := DB.Select(&ret, "select class_id from class where class_headmaster=?", teaID)
	if err != nil {
		return ret, err
	}
	for _, c := range classS {
		ret = append(ret, c.Class_id)
	}

	return ret, nil
}

//SelectClassCourse 根据班级ID从排课表中选出对应的所有课程id
func SelectClassCourse(classID string) ([]string, error) {
	var ret []string
	var courses []Course
	err := DB.Select(&courses, "select distinct course_id from time_table where class_id=?", classID)
	if err != nil {
		return ret, err
	}
	for _, c := range courses {
		ret = append(ret, c.Course_ID)
	}

	return ret, nil
}

//BeginOrEnd 用于SelectBeginOrEnd方法的结构体
type Begin struct {
	Begin int `json:begin`
}

//SelectBeginOrEnd 从begin_or_end表中选出两个字段，用于判断是否为选课阶段
func SelectBeginOrEnd() (int, error) {
	var b Begin
	err := DB.Get(&b, "select begin from begin")
	return b.Begin, err
}

//CourseWaitSelect 用于SelectCourseFromClassOnWaitSelect方法的结构体
type CourseWaitSelect struct {
	Course_ID string `json:course_id`
	Class_ID  string `json:class_id`
	Stu_num   int    `json:stu_num`
}

//SelectCourseFromClassOnWaitSelect 从待选课表中选出班级对应的课程ID
func SelectCourseFromClassOnWaitSelect(classID string) ([]CourseWaitSelect, error) {
	var ret []CourseWaitSelect
	err := DB.Select(&ret, "select course_id, stu_num, class_id from course_wait_for_select where class_id=?", classID)
	return ret, err
}

//SelectCourseFromTimetable 从排课表中选出班级对应的课程信息
func SelectCourseFromTimetable(classID string, courseID string) ([]Table, error) {
	var ret []Table
	err := DB.Select(&ret, "select course_id, group_concat(tea_id) as tea_id , group_concat(section) as section, group_concat(day) as day, begin_time, end_time from time_table where class_id=? and course_id=? group by course_id, begin_time, end_time", classID, courseID)
	return ret, err
}

//SelectCourseFromTimetableDiv 从排课表中选出班级对应的课程信息，信息不合并
func SelectCourseFromTimetableDiv(classID string, courseID string) ([]Table, error) {
	var ret []Table
	err := DB.Select(&ret, "select course_id, tea_id, section, day, begin_time, end_time from time_table where class_id=? and course_id=?", classID, courseID)
	return ret, err
}

//SelectClassFromTimetable 从排课表中选出班级（选课时使用）
func SelectClassFromTimetable(table Table) ([]string, error) {
	var ret []string
	var cs []Class
	err := DB.Select(&cs, "select class_id from time_table where course_id=? and tea_id=? and begin_time=? and end_time=? and section=? and day=?", table.Course_ID, table.Tea_ID, table.Begin_time, table.End_time, table.Section, table.Day)
	if err != nil {
		return ret, err
	}
	for _, c := range cs {
		ret = append(ret, c.Class_id)
	}
	return ret, nil
}

func InsertCourseSelect(stuID string, courseID string) error {
	_, err := CONN.Exec("insert into course_select(course_id, stu_id) values (?,?)", courseID, stuID)
	return err
}

func UpdateCourseWaitSelectAdd(classID string, courseID string) error {
	_, err := CONN.Exec("update course_wait_for_select set stu_num = stu_num+1 where class_id=? and course_id=?", classID, courseID)
	return err
}

func DeleteCourseSelect(stuID string, courseID string) error {
	_, err := CONN.Exec("delete from course_select where course_id=? and stu_id=?", courseID, stuID)
	return err
}

func UpdateCourseWaitSelectSub(classID string, courseID string) error {
	_, err := CONN.Exec("update course_wait_for_select set stu_num = stu_num-1 where class_id=? and course_id=?", classID, courseID)
	return err
}

func SelectClass() ([]Class, error) {
	var ret []Class
	err := DB.Select(&ret, "select class_id, class_headmaster from class")
	return ret, err
}

func SelectAllStu() ([]Stu, error) {
	var ret []Stu
	err := DB.Select(&ret, "select stu_id, stu_class, stu_name, gpa, stu_sex from student")
	return ret, err
}

func SelectAllTeaID() ([]string, error) {
	var teaS []Tea
	var ret []string
	err := DB.Select(&teaS, "select tea_id from teacher")
	if err != nil {
		return ret, err
	}
	for _, t := range teaS {
		ret = append(ret, t.Tea_ID)
	}
	return ret, nil
}

func UpdateClassHeadmaster(classID string, teaID string) error {
	_, err := DB.Exec("update class set class_headmaster=? where class_id=?", teaID, classID)
	return err
}

func SelectAllTeaCourse(teaID string) ([]string, error) {
	var courseS []Table
	var courses []string
	err := DB.Select(&courseS, "select course_id from time_table where tea_id=?", teaID)
	if err != nil {
		return courses, err
	}
	for _, c := range courseS {
		courses = append(courses, c.Course_ID)
	}
	return courses, nil
}

func SelectAllTea() ([]Tea, error) {
	var teas []Tea
	err := DB.Select(&teas, "select tea_id, tea_name from teacher")
	return teas, err
}

func SelectCourseFromTeaIDOnCG(teaID string) ([]string, error) {
	var ret []string
	var courses []Course
	err := DB.Select(&courses, "select course_id from course_group where tea_id=?", teaID)
	if err != nil {
		return ret, err
	}
	for _, c := range courses {
		ret = append(ret, c.Course_ID)
	}
	return ret, nil
}

func SelectAllCourses() ([]Course, error) {
	var courses []Course
	err := DB.Select(&courses, "select course_id, course_name from course")
	return courses, err
}

func SelectTeaFromCourseIDOnGroup(courseID string) ([]string, error) {
	var ret []string
	var teaS []Tea
	err := DB.Select(&teaS, "select tea_id from course_group where course_id=?", courseID)
	if err != nil {
		return ret, err
	}
	for _, t := range teaS {
		ret = append(ret, t.Tea_ID)
	}
	return ret, nil
}

func DeleteGroup(courseID string, teaID string) error {
	_, err := DB.Exec("Delete from course_group where tea_id=? and course_id=?", teaID, courseID)
	return err
}
func AddGroup(courseID string, teaID string) error {
	_, err := DB.Exec("insert into course_group(course_id, tea_id)values(?,?)", courseID, teaID)
	return err
}

type CC struct {
	Course_ID string `json:course_id`
	Class_ID  string `json:class_id`
}

func SelectCCOnTimetable(teaID string) ([]CC, error) {
	var ret []CC
	err := DB.Select(&ret, "select course_id, class_id from time_table where tea_id=?", teaID)
	return ret, err
}

type CCStu struct {
	Course_ID string        `json:course_id`
	Stu_ID    string        `json:stu_id`
	Mark      sql.NullInt32 `json:mark`
}

func SelectAllStuGrades(courseIDAndClsID CC) ([]CCStu, error) {
	var ret []CCStu
	err := DB.Select(&ret, "select student.stu_id, course_id, mark from course_select join student on student.stu_id = course_select.stu_id where course_id=? and stu_class=?", courseIDAndClsID.Course_ID, courseIDAndClsID.Class_ID)
	return ret, err
}

func SelectAllStuMap() (map[string]string, error) {
	var StuIDMapName map[string]string
	StuIDMapName = make(map[string]string)
	var stu []Stu
	err := DB.Select(&stu, "select stu_id, stu_name from student")
	if err != nil {
		return StuIDMapName, err
	}
	for _, s := range stu {
		StuIDMapName[s.Stu_ID] = s.Stu_name
	}
	return StuIDMapName, nil
}

func UpdateGrades(courseID string, stuID string, grade int) error {
	_, err := DB.Exec("update course_select set mark = ? where course_id=? and stu_id=?", grade, courseID, stuID)
	return err
}

type GG struct {
	Course_credit int           `json:course_credit`
	Mark          sql.NullInt32 `json:mark`
}

func SelectGG(stuID string) ([]GG, error) {
	var ret []GG
	err := DB.Select(&ret, "select course_credit, mark from course_select join course on course_select.course_id = course.course_id where stu_id=?", stuID)
	return ret, err
}

func UpdateGpa(gpa float64, stuID string) error {
	_, err := DB.Exec("update student set gpa =? where stu_id=?", gpa, stuID)
	return err
}
