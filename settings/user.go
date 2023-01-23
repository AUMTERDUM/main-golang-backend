package settings

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"

	//"net/http"

	//"github.com/gorilla/mux"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//convert to fiber

func CreateUser(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var product entities.User
	c.BodyParser(&product)
	database.Instance.Create(&product)
	c.JSON(product)
	return c.JSON(product)
}

func GetUserById(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfUserExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product []entities.User
	var systems []entities.System
	database.Instance.First(&product, productId)
	database.Instance.Find(&systems)
	for index, data := range product {
		product[index].ListSystem = mapSystem(data.Systems, systems)
	}
	
	return c.JSON(product)
}


func GetUsers(c *fiber.Ctx) error {
	var products []entities.User
	var systems []entities.System
	database.Instance.Find(&products)
	database.Instance.Find(&systems)
	for index, data := range products {
		products[index].ListSystem = mapSystem(data.Systems, systems)
	}

	c.JSON(products)
	return c.JSON(products)

}

func mapSystem(listStr string, systems []entities.System) []entities.System {
	list := strings.Split(listStr, ",")
	var data []entities.System
	for _, v := range list {
		for _, s := range systems {
			id, _ := strconv.Atoi(v)
			if id == s.ID {
				data = append(data, s)
			}
		}
	}
	return data
}

func UpdateUser(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfUserExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.User
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteUser(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfUserExists(productId) == false {
		return c.JSON("Product Not Found!")
	}
	var product entities.User
	database.Instance.First(&product, productId)
	database.Instance.Delete(&product)
	return c.JSON("Product Deleted!")
}

func checkIfUserExists(id string) bool {
	var product entities.User
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}


