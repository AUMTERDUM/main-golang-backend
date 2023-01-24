package database

import (
	"golang-crud-rest-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB // เป็นตัวแปร Instance ที่เป็น pointer ของ gorm.DB
var err error // เป็นตัวแปร err ที่เป็น error

func Connect(connectionString string) { // ฟังก์ชัน Connect ที่รับ parameter connectionString และเป็น void
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{}) // แบบนี้คือการเรียกใช้ gorm เพื่อเปิดการเชื่อมต่อกับ MySQL โดยใช้ connectionString ที่ส่งเข้ามา
	if err != nil { // ถ้าเกิด error ให้ทำงานใน if
		log.Fatal(err) // แสดง error ที่เกิดขึ้น
		panic("Cannot connect to DB") // แสดง error ที่เกิดขึ้นเป็น panic
	}
	log.Println("Connected to Database...") // แสดงข้อความ Connected to Database...
}

func MigrateUSER() { // ฟังก์ชัน MigrateUSER ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.User{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.User โดยใช้ gorm ในการสร้าง
	log.Println("Database USER Completed...") // แสดงข้อความ Database USER Completed...
}

func MigrateSYSTEM() { // ฟังก์ชัน MigrateSYSTEM ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.System{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.System โดยใช้ gorm ในการสร้าง
	log.Println("Database SYSTEM Completed...") // แสดงข้อความ Database SYSTEM Completed...
}

func MigratePROBLEM() { // ฟังก์ชัน MigratePROBLEM ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.Problemtype{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.Problemtype โดยใช้ gorm ในการสร้าง
	log.Println("Database PROBLEM Completed...") // แสดงข้อความ Database PROBLEM Completed...
}

func MigrateLEVEL() { // ฟังก์ชัน MigrateLEVEL ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.Level{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.Level โดยใช้ gorm ในการสร้าง
	log.Println("Database LEVEL Completed...") // แสดงข้อความ Database LEVEL Completed...
}

func MigrateCONTACT() { // ฟังก์ชัน MigrateCONTACT ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.Contact{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.Contact โดยใช้ gorm ในการสร้าง
	log.Println("Database CONTACT Completed...") // แสดงข้อความ Database CONTACT Completed...
}

func MigrateANGENCY() { // ฟังก์ชัน MigrateANGENCY ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.Agency{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.Agency โดยใช้ gorm ในการสร้าง
	log.Println("Database ANGENCY Completed...") // แสดงข้อความ Database ANGENCY Completed...
}

func MigratePROBLEMRECORD() { // ฟังก์ชัน MigratePROBLEMRECORD ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.ProblemRecord{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.Problemrecord โดยใช้ gorm ในการสร้าง
	log.Println("Database PROBLEMRECORD Completed...") // แสดงข้อความ Database PROBLEMRECORD Completed...
}

func MigrateSTATUS() { // ฟังก์ชัน MigrateSTATUS ที่ไม่รับ parameter และเป็น void
	Instance.AutoMigrate(&entities.Status{}) // เป็นการสร้าง table ตาม struct ที่อยู่ใน entities.Status โดยใช้ gorm ในการสร้าง
	log.Println("Database STATUS Completed...") // แสดงข้อความ Database STATUS Completed...
}
