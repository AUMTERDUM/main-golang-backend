package problemrecord

import (
	//"encoding/json"
	"fmt"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	//sqldriver
	//"database/sql/driver"
	"math"

	//"golang.org/x/crypto/bcrypt"
	//"github.com/gorilla/mux"
	"github.com/gofiber/fiber/v2"
)

func CreateProblemRecord(c *fiber.Ctx) error {
	file, err := c.FormFile("problem_records")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	ext := filepath.Ext(file.Filename)
	//get file name
	fileName := fmt.Sprint(time.Now().Unix()) + ext
	//get file extension

	//get file size
	size := file.Size
	//get file path
	filePath := "/upload/" + fileName
	//log file data
	fmt.Println("File Info")
	fmt.Println("File Name:", fileName)
	fmt.Println("File Size:", size)
	fmt.Println("File Path:", filePath)
	fmt.Println("File Extension:", ext)

	//upload file in folder
	filename := path.Join("./uploads/", path.Base(fileName))
	//create file
	dest, err := os.Create(filename)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer dest.Close()
	// Copy data in file
	if err := c.SaveFile(file, dest.Name()); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	agency := c.FormValue("agency")
	contact := c.FormValue("contact")
	problem := c.FormValue("problem")
	level := c.FormValue("level")
	informer := c.FormValue("informer")
	informermessage := c.FormValue("informermessage")
	system := c.FormValue("system")
	problemtype := c.FormValue("problemtype")

	problemrecord := entities.ProblemRecord{
		File_name:       fileName,
		Path_file:       filePath,
		Agency:          agency,
		Contact:         contact,
		Problem:         problem,
		Level:           level,
		Informer:        informer,
		Informermessage: informermessage,
		System:          system,
		Problemtype:     problemtype,
		CreatedAt:       time.Now(),
		Status:          1,
		//Statusname       Statusname,

		File_extension: ext,
		File_size:      int(size),
	}
	entities.ReadJsonFormLocal()//read json file from local and save to database table
	database.Instance.Create(&problemrecord)
	return c.JSON(fiber.Map{"id": problemrecord.ID, "file_name": problemrecord.File_name, "path_file": problemrecord.Path_file, "agency": problemrecord.Agency, "contact": problemrecord.Contact, "problem": problemrecord.Problem, "level": problemrecord.Level, "informer": problemrecord.Informer, "informermessage": problemrecord.Informermessage, "system": problemrecord.System, "problemtype": problemrecord.Problemtype, "created_at": problemrecord.CreatedAt, "status": problemrecord.Status, "file_extension": problemrecord.File_extension, "file_size": problemrecord.File_size, "message": "Create Successfully"})
	
	

}

func GetProblemRecords(c *fiber.Ctx) error {
	var repo entities.Meta
	var systems []entities.System
	var contacts []entities.Contact
	var problemtype []entities.Problemtype
	var agencies []entities.Agency
	var levels []entities.Level
	var users []entities.User

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}
	offset := (page - 1) * limit
	//database.Instance.Preload("Statuse").Limit(limit).Offset(offset).Find(&repo.ProblemRecord)
	database.Instance.Order("created_at desc").Preload("Statuse").Limit(limit).Offset(offset).Find(&repo.ProblemRecord) //order by created_at desc
	database.Instance.Find(&systems)
	database.Instance.Find(&contacts)
	database.Instance.Find(&problemtype)
	database.Instance.Find(&agencies)
	database.Instance.Find(&levels)
	database.Instance.Find(&users)

	for index, data := range repo.ProblemRecord {
		repo.ProblemRecord[index].Systems = mapSystem(data.System, systems)
	}

	for index, data1 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Contacts = MapContact(data1.Contact, contacts)
	}

	for index, data2 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Problemtypes = MapProblemType(data2.Problemtype, problemtype)
	}

	for index, data3 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Agencies = MapAgnecy(data3.Agency, agencies)
	}

	for index, data4 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Levels = MapLevel(data4.Level, levels)
	}

	for index, data5 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Users = MapUser(data5.Operator, users)
		for index, data := range users {
			users[index].ListSystem = mapSystem(data.Systems, systems)
		}
	}

	repo.Pageination = Pagination(c)
	c.JSON(repo)
	c.Set("Content-Type", "application/json")
	return c.JSON(repo)
}

//just for test
func GetProblemRecordByTime(c *fiber.Ctx) error {
	var repo entities.Meta
	var systems []entities.System
	var contacts []entities.Contact
	var problemtype []entities.Problemtype
	var agencies []entities.Agency
	var levels []entities.Level
	var users []entities.User

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}
	offset := (page - 1) * limit
	database.Instance.Order("created_at").Preload("Statuse").Limit(limit).Offset(offset).Find(&repo.ProblemRecord) //.Where("created_at BETWEEN ? AND ?", "2021-01-01", "2021-01-31")
	//database.Instance.Preload("Statuse").Limit(limit).Offset(offset).Find(&repo.ProblemRecord)
	database.Instance.Find(&systems)
	database.Instance.Find(&contacts)
	database.Instance.Find(&problemtype)
	database.Instance.Find(&agencies)
	database.Instance.Find(&levels)
	database.Instance.Find(&users)


	for index, data := range repo.ProblemRecord {
		repo.ProblemRecord[index].Systems = mapSystem(data.System, systems)
	}

	for index, data1 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Contacts = MapContact(data1.Contact, contacts)
	}

	for index, data2 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Problemtypes = MapProblemType(data2.Problemtype, problemtype)
	}

	for index, data3 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Agencies = MapAgnecy(data3.Agency, agencies)
	}

	for index, data4 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Levels = MapLevel(data4.Level, levels)
	}

	for index, data5 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Users = MapUser(data5.Operator, users)
		for index, data := range users {
			users[index].ListSystem = mapSystem(data.Systems, systems)
		}
	}

	repo.Pageination = Pagination(c) // ส่งค่า pageination ไปให้ Pagination ทำงาน แล้ว return ค่ากลับมา แล้วเก็บไว้ใน repo.Pageination แล้วส่งค่าไปให้ client ด้วย
	c.JSON(repo)
	c.Set("Content-Type", "application/json")
	return c.JSON(repo)
}


func GetProblemById(c *fiber.Ctx) error {
	id := (c.Params("id"))
	fmt.Printf("id: %s ", id)
	var repo entities.Meta
	var systems []entities.System
	var contacts []entities.Contact
	var problemtype []entities.Problemtype
	var agencies []entities.Agency
	var levels []entities.Level
	var users []entities.User
	database.Instance.Preload("Statuse").First(&repo.ProblemRecord, fmt.Sprintf(`id = '%s'`, id)) // ค้นหาข้อมูลจาก id ที่ส่งมา แล้วเก็บไว้ใน repo.ProblemRecord
	//database.Instance.Preload(&systems).First(&repo.ProblemRecord)
	database.Instance.Find(&systems)
	database.Instance.Find(&contacts)
	database.Instance.Find(&problemtype)
	database.Instance.Find(&agencies)
	database.Instance.Find(&levels)
	database.Instance.Find(&users)

	for index, data := range repo.ProblemRecord {
		repo.ProblemRecord[index].Systems = mapSystem(data.System, systems)
	}

	for index, data1 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Contacts = MapContact(data1.Contact, contacts)
	}

	for index, data2 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Problemtypes = MapProblemType(data2.Problemtype, problemtype)
	}

	for index, data3 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Agencies = MapAgnecy(data3.Agency, agencies)
	}

	for index, data4 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Levels = MapLevel(data4.Level, levels)
	}

	for index, data5 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Users = MapUser(data5.Operator, users) // ส่งค่า user ไปให้ MapUser ทำงาน แล้ว return ค่ากลับมา แล้วเก็บไว้ใน repo.ProblemRecord[index].Users
		for index, data := range users {
			users[index].ListSystem = mapSystem(data.Systems, systems) // ส่งค่า system ไปให้ mapSystem ทำงาน แล้ว return ค่ากลับมา แล้วเก็บไว้ใน users[index].ListSystem
		} // แล้วเก็บไว้ใน repo.ProblemRecord[index].Users
	}

	
	c.JSON(repo)
	c.Set("Content-Type", "application/json")
	return c.JSON(repo)
}

func mapSystem(listStr string, systems []entities.System) []entities.System {
	// Split the comma-separated list of system IDs into a slice of strings
list := strings.Split(listStr, ",")
	var data []entities.System
	// Loop over the list of strings and compare each ID to the system ID
	for _, v := range list {
		for _, s := range systems {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				// If there is a match, append the system data to the data slice
				data = append(data, s)
			}
		}
	}
	return data
}

func MapContact(listStr string, contacts []entities.Contact) []entities.Contact {
	list := strings.Split(listStr, ",")
	var data1 []entities.Contact
	for _, v := range list {
		for _, s := range contacts {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				data1 = append(data1, s)
			}
		}
	}
	return data1
}

func MapProblemType(listStr string, problemtype []entities.Problemtype) []entities.Problemtype {
	list := strings.Split(listStr, ",")
	var data2 []entities.Problemtype
	for _, v := range list {
		for _, s := range problemtype {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				data2 = append(data2, s)
			}
		}
	}
	return data2
}

func MapAgnecy(listStr string, agnecy []entities.Agency) []entities.Agency {
	list := strings.Split(listStr, ",")
	var data3 []entities.Agency
	for _, v := range list {
		for _, s := range agnecy {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				data3 = append(data3, s)
			}
		}
	}
	return data3
}

func MapLevel(listStr string, level []entities.Level) []entities.Level {
	list := strings.Split(listStr, ",")
	var data4 []entities.Level
	for _, v := range list {
		for _, s := range level {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				data4 = append(data4, s)
			}
		}
	}
	return data4
}

func MapUser(listStr string, operator []entities.User) []entities.User {
	list := strings.Split(listStr, ",")
	var data5 []entities.User
	for _, v := range list {
		for _, s := range operator {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				data5 = append(data5, s)
			}
		}
	}
	return data5
}

func GetProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("id = ?", id).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)

	return c.JSON(problemrecord)
}

func UpdateProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	operator := c.FormValue("operator")
	var data_problem entities.ProblemRecord
	database.Instance.Where("id = ?", id).Find(&data_problem)
	fmt.Println(data_problem.ID)
	if id != data_problem.ID {

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	problemrecord := entities.ProblemRecord{
		Operator: operator,
		SenderAt: time.Now(),
		Status:   2,
	}

	if err := c.BodyParser(&problemrecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if database.Instance.Where("id = ?", id).Updates(&problemrecord).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Update File",
		})
	}

	return c.JSON(fiber.Map{"operator": problemrecord.Operator, "status": problemrecord.Status, "sender_at": problemrecord.SenderAt, "message": "Update Successfully"})

}

func CompletedProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	Casuseproblem := c.FormValue("casuseproblem")
	Solution := c.FormValue("solution")
	Suggestion := c.FormValue("suggestion")
	

	var data_problem entities.ProblemRecord
	database.Instance.Where("id = ?", id).Find(&data_problem)
	fmt.Println(data_problem.ID)
	if id != data_problem.ID {

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	problemrecord := entities.ProblemRecord{
		Casuseproblem: Casuseproblem,
		Solution:      Solution,
		Suggestion:    Suggestion,
		CompletedAt:   time.Now(),
		Status:        3,
		Time: 		   (data_problem.CompletedAt).Sub(data_problem.CreatedAt).String(), //คำนวนเวลา วันที่เส็จ-วันที่สร้าง แล้วเก็บไว้ในตัวแปร Time ในตาราง ProblemRecord ในฐานข้อมูล
	}



	if err := c.BodyParser(&problemrecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if database.Instance.Where("id = ?", id).Updates(&problemrecord).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Update File",
		})
	}

	//problemrecord.Time = time.Now().Sub(problemrecord.SenderAt).String()
	CalculateTime(c) 

	

	return c.JSON(fiber.Map{"casuseproblem": problemrecord.Casuseproblem, "solution": problemrecord.Solution, "suggestion": problemrecord.Suggestion, "status": problemrecord.Status, "completed_at": problemrecord.CompletedAt, "message": "Update Successfully", "time": problemrecord.Time})
	//database.Instance.Where("id = ?",id).Save(&problemrecord)
	//return c.JSON(problemrecord)
}

func CancalProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	Casuseproblem := c.FormValue("casuseproblem")
	Solution := c.FormValue("solution")
	Suggestion := c.FormValue("suggestion")


	var data_problem entities.ProblemRecord
	database.Instance.Where("id = ?", id).Find(&data_problem)
	fmt.Println(data_problem.ID)
	if id != data_problem.ID {

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	problemrecord := entities.ProblemRecord{
		Casuseproblem: Casuseproblem,
		Solution:      Solution,
		Suggestion:    Suggestion,
		CompletedAt:   time.Now(),
		Status:        4,
		Time :         (data_problem.CompletedAt).Sub(data_problem.CreatedAt).String(),
	}

	if err := c.BodyParser(&problemrecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if database.Instance.Where("id = ?", id).Updates(&problemrecord).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Update File",
		})
	}

	return c.JSON(fiber.Map{"casuseproblem": problemrecord.Casuseproblem, "solution": problemrecord.Solution, "suggestion": problemrecord.Suggestion, "status": problemrecord.Status, "completed_at": problemrecord.CompletedAt, "message": "Update Successfully"})
	//database.Instance.Where("id = ?",id).Save(&problemrecord)
	//return c.JSON(problemrecord)
}


//calculate time ไม่น่าจะได้ใช้
func CalculateTime(c *fiber.Ctx) error {

	
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("id = ?", id).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	fmt.Println(problemrecord.CreatedAt)	
	fmt.Println(problemrecord.CompletedAt) 
	//fmt.Println(problemrecord.CompletedAt.Sub(problemrecord.CreatedAt).Hours())   // 1.5 hours difference between the two times in hours (1.5)
	//fmt.Sprintf("2023-01-23 10:01:10").Sub("2023-01-23 10:53:46").Hours()
	//fmt.Println(problemrecord.CreatedAt.Sub(problemrecord.CompletedAt).Minutes()) // 90 minutes difference between the two times in minutes (90)
	//fmt.Println(problemrecord.CreatedAt.Sub(problemrecord.CompletedAt).Seconds()) // 5400 seconds difference between the two times in seconds (5400)

	//convert time to string
	fmt.Println(problemrecord.CompletedAt.Sub(problemrecord.CreatedAt).String()) // 1h30m0s difference between the two times in string (1h30m0s)
	//convert time to int
	//fmt.Println(int(problemrecord.CreatedAt.Sub(problemrecord.CompletedAt).Hours())) // 1 hours difference between the two times in int (1)
	//fmt.Println(int(problemrecord.CreatedAt.Sub(problemrecord.CompletedAt).Minutes())) // 90 minutes difference between the two times in int (90)
	//fmt.Println(int(problemrecord.CreatedAt.Sub(problemrecord.CompletedAt).Seconds())) // 5400 seconds difference between the two times in int (5400)
	


	return c.JSON(fiber.Map{"time": problemrecord.CreatedAt.Sub(problemrecord.CompletedAt).String()})
}

func DeleteProblemRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var problemrecord entities.ProblemRecord
	database.Instance.First(&problemrecord, id)
	if problemrecord.ID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	database.Instance.Delete(&problemrecord)
	return c.JSON(fiber.Map{
		"message": "Record deleted",
	})
}

//ค้นหาข้อมูล แบบเลือกเฉพาะ 1 อย่าง แล้วค้นหาแล้วแสดงผลทั้งหมดที่เจอออกมาแบบเป็น array ที่เป็น json 
func GetProblemRecordByAgency(c *fiber.Ctx) error {
	agency := c.Params("agency")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("agency = ?", agency).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByContact(c *fiber.Ctx) error {
	contact := c.Params("contact")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("contact = ?", contact).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblem(c *fiber.Ctx) error {
	problem := c.Params("problem")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problem = ?", problem).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByLevel(c *fiber.Ctx) error {
	level := c.Params("level")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("level = ?", level).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByInformer(c *fiber.Ctx) error {
	informer := c.Params("informer")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("informer = ?", informer).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByInformermessage(c *fiber.Ctx) error {
	informermessage := c.Params("informermessage")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("informermessage = ?", informermessage).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordBySystem(c *fiber.Ctx) error {
	system := c.Params("system")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("system = ?", system).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemtype(c *fiber.Ctx) error {
	problemtype := c.Params("problemtype")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemtype = ?", problemtype).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemstatus(c *fiber.Ctx) error {
	problemstatus := c.Params("problemstatus")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemstatus = ?", problemstatus).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemtime(c *fiber.Ctx) error {
	problemtime := c.Params("problemtime")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemtime = ?", problemtime).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

func GetProblemRecordByProblemdescription(c *fiber.Ctx) error {
	problemdescription := c.Params("problemdescription")
	var problemrecord entities.ProblemRecord
	database.Instance.Where("problemdescription = ?", problemdescription).Find(&problemrecord)
	c.Set("Content-Type", "application/json")
	c.JSON(problemrecord)
	return nil
}

//pagination การจัดหน้า
func Pagination(c *fiber.Ctx) entities.Pageination {
	var problemrecord entities.ProblemRecord
	var total_row int64
	var page, limit int
	var err error
	page, err = strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		limit = 10
	}
	offset := (page - 1) * limit
	database.Instance.Model(&problemrecord).Count(&total_row)
	database.Instance.Limit(limit).Offset(offset).Find(&problemrecord)

	return entities.Pageination{
		Page:     page,
		Limit:    limit,
		Pages:    int(math.Ceil(float64(total_row) / float64(limit))),
		TotalRow: total_row,
	}
}

