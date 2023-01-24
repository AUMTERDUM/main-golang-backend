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

func CreateSystem(c *fiber.Ctx) error {
	var product entities.System
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetSystemById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("System Not Found! ไม่พบระบบ")
	}
	var product entities.System
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetSystems(c *fiber.Ctx) error {
	var products []entities.System
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateSystem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("System Not Found! ไม่พบระบบ")
	}
	var product entities.System
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteSystem(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("System Not Found! ไม่พบระบบ")
	}
	var product entities.System
	database.Instance.Delete(&product, productId)
	return c.JSON("System Deleted! ลบระบบแล้ว")
}

func checkIfSystemExists(id string) bool {
	var product entities.System
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}

