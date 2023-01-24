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
func CreateStatus(c *fiber.Ctx) error {
	var product entities.Status
	c.BodyParser(&product)
	database.Instance.Create(&product)
	return c.JSON(product)
}

func GetStatusById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfStatusExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Status
	database.Instance.First(&product, productId)
	return c.JSON(product)
}

func GetStatuss(c *fiber.Ctx) error {
	var products []entities.Status
	database.Instance.Find(&products)
	return c.JSON(products)
}

func UpdateStatus(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfStatusExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Status
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteStatus(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfStatusExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.Status
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("Product Deleted!")
}

func checkIfStatusExists(id string) bool {
	var product entities.Status
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}

