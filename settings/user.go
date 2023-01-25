package settings

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"math"

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
		return c.JSON("Operator Not Found! ไม่พบผู้รับผิดชอบ")
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

func GetUserss(c *fiber.Ctx) error {
	var products []entities.User
	var systems []entities.System
	var total_row int64

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "100"))
	if err != nil {
		limit = 10
	}

	offset := (page - 1) * limit
	//database.Instance.Select("status").Where("status = ?","Inactive").Find(&products)
	database.Instance.Order("id desc").Preload("users").Limit(limit).Offset(offset).Find(&products)
	database.Instance.Model(&entities.User{}).Count(&total_row)
	database.Instance.Find(&products)
	database.Instance.Find(&systems)
	for index, data := range products {
		products[index].ListSystem = mapSystem(data.Systems, systems)
	}

	products = products[offset : offset+limit]
	
	

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
		return c.JSON("Operator Not Found! ไม่พบผู้รับผิดชอบ")
	}
	var product entities.User
	database.Instance.First(&product, productId)
	c.BodyParser(&product)
	database.Instance.Save(&product)
	return c.JSON(product)
}

func DeleteUser(c *fiber.Ctx) error {
	productId := c.Params("id")
	if checkIfSystemExists(productId) == false {
		return c.JSON("Operator Not Found! ไม่พบระบบ")
	}
	database.Instance.Delete(&entities.User{}, productId)
	return c.JSON("Operator Deleted! ลบผู้รับผิดชอบเรียบร้อยแล้ว")
}

func checkIfUserExists(id string) bool {
	var product entities.User
	database.Instance.First(&product, id)
	if product.ID == 0 {
		return false
	}
	return true
}

//DeleteUser But not Delete in Database  just for test

func DeleteUserButNotDelete(c *fiber.Ctx) error {
	productId := c.Params("id")
	database.Instance.Delete(&entities.User{}, productId)
	return c.JSON("Operator Deleted! ลบผู้รับผิดชอบเรียบร้อยแล้ว")
}

func Pagination(total_row int64, page int, limit int) entities.Pageination {

	return entities.Pageination{
		Page:     page,
		Limit:    limit,
		Pages:    int(math.Ceil(float64(total_row) / float64(limit))),
		TotalRow: total_row,
	}
}
