//not use


package problemrecord

import (
	//"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	//"net/http"
	"github.com/gofiber/fiber/v2"

)

func PublicLinks(c *fiber.Ctx) error {
	id := c.Params("id")
	var repo entities.Meta
	var systems []entities.System
	var contacts []entities.Contact
	var problemtype []entities.Problemtype
	var agencies []entities.Agency
	var levels []entities.Level
	var users []entities.User
	
	database.Instance.Preload("Statuse").First(&repo.ProblemRecord, id)
	database.Instance.Find(&systems)
	database.Instance.Find(&contacts)
	database.Instance.Find(&problemtype)
	database.Instance.Find(&agencies)
	database.Instance.Find(&levels)
	database.Instance.Find(&users)
	
	// for index, data := range systems {
	// 	fmt.Println(data.Name, index)
	// }

	for index, data := range repo.ProblemRecord {
		repo.ProblemRecord[index].Systems = mapSystem(data.System, systems)
	}

	for index, data1 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Contacts = MapContact(data1.Contact,contacts)
	}

	for index, data2 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Problemtypes = MapProblemType(data2.Problemtype,problemtype)
	}

	for index, data3 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Agencies = MapAgnecy(data3.Agency,agencies)
	}

	for index, data4 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Levels = MapLevel(data4.Level,levels)
	}

	for index, data5 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Users = MapUser(data5.Operator,users)
		for index, data := range users {
			users[index].ListSystem = mapSystem(data.Systems, systems)
		}
	}


	c.JSON(repo)
	c.Set("Content-Type", "application/json")
	return c.JSON(repo)
}

func PublicLink(c *fiber.Ctx) error {
	id := c.Params("id")
	var repo entities.Meta
	var systems []entities.System
	var contacts []entities.Contact
	var problemtype []entities.Problemtype
	var agencies []entities.Agency
	var levels []entities.Level
	var users []entities.User

	database.Instance.Preload("Statuse").First(&repo.ProblemRecord, id)
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
		repo.ProblemRecord[index].Contacts = MapContact(data1.Contact,contacts)
	}

	for index, data2 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Problemtypes = MapProblemType(data2.Problemtype,problemtype)
	}

	for index, data3 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Agencies = MapAgnecy(data3.Agency,agencies)
	}

	for index, data4 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Levels = MapLevel(data4.Level,levels)
	}

	for index, data5 := range repo.ProblemRecord {
		repo.ProblemRecord[index].Users = MapUser(data5.Operator,users)
		for index, data := range users {
			users[index].ListSystem = mapSystem(data.Systems, systems)
		}
	}
	c.JSON(repo)
	return c.JSON(repo)
}
