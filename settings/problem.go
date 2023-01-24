package settings

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	//"net/http"

	//"github.com/gorilla/mux"
	"github.com/gofiber/fiber/v2"
)

//convert to fiber

func CreateProblem(c *fiber.Ctx) error {
	var product entities.Problemtype
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetProblemById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfProblemExists(productId) == false {
		return c.JSON("ProbelmType Not Found! ไม่พบประเภทของปัญหา")
	}
	var product entities.Problemtype
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetProblems(c *fiber.Ctx) error {
	var products []entities.Problemtype
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateProblem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfProblemExists(productId) == false {
		return c.JSON("ProblemType Not Found! ไม่พบประเภทของปัญหา")
	}
	var product entities.Problemtype 
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteProblem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfProblemExists(productId) == false {
		return c.JSON("ProblemType Not Found! ไม่พบประเภทของปัญหา")
	}
	var product entities.Problemtype
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("ProbelmType Deleted! ลบประเภทของปัญหาสำเร็จ")
}

func checkIfProblemExists(id string) bool {
	var product entities.Problemtype
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}
