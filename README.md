# Implementing CRUD in GoLang REST API with Mux & GORM

![CRUD in GoLang REST API with Fiber & GORM]
CRUD, หรือ Create, Read, Update, Delete เป็นการเขียนแอปพลิเคชันในแบบ REST API โดยใช้ GoLang เป็นเทคโนโลยีหลัก ร่วมกับ Fiber และ GORM เป็นไลบรารี่สำหรับการจัดการฐานข้อมูล
Fiber เป็น web framework สำหรับ GoLang ที่มีความเร็วและปลอดภัย ซึ่งจะใช้ในการเขียน routing และ middleware ของแอปพลิเคชัน
GORM เป็น ORM สำหรับ GoLang ที่ช่วยในการจัดการฐานข้อมูลแบบมาตรฐาน โดยสามารถเชื่อมต่อกับฐานข้อมูลต่างๆ เช่น MySQL, PostgreSQL, SQLite ได้
ในการเขียน CRUD ด้วย GoLang กับ Fiber & GORM จะเริ่มจากการติดตั้งไลบรารี่ และกำหนดค่าการเชื่อมต่อฐานข้อมูล จากนั้นจะต้องเขียนโค้ดเพื่อสร้าง endpoint สำหรับ CRUD ต่างๆ โดยการใช้ Fiber สำหรับ routing และ GORM สำหรับการจัดการฐานข้อมูล

## Topics Covered
- Setting up the Golang Project
- Loading Configurations using Viper
- Defining the Product Entity
- Connecting to the database
- Routing
- Implementing CRUD in Golang Rest API
	 - Create
	 - Get By ID
	 - Get All
	 - Update
	 - Delete
- Testing CRUD Operations
# TESTSS
