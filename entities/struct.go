package entities

import (
	"encoding/json"
	"io"
	"strings"
	"time"

	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         int      `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"type:varchar(50)" json:"name"`
	Nickname   string   `gorm:"type:varchar(50)" json:"nickname"`
	Systems    string   `json:"systems"`
	ListSystem []System `json:"listsystems" gorm:"-"`
	//Status	 string   `gorm:"type:varchar(50)" json:"status"`
	Deleted gorm.DeletedAt
}

type MetaUser struct {
	Pageination Pageination `json:"pageination"`
	Users       []User      `json:"operator"`
}

type System struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Deleted gorm.DeletedAt
}

type Problemtype struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Deleted gorm.DeletedAt
}

type Level struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Time    int    `json:"time"`
	Deleted gorm.DeletedAt
}

type Contact struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Deleted gorm.DeletedAt
}

type Agency struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Deleted gorm.DeletedAt
}

func (book *ProblemRecord) BeforeCreate(tx *gorm.DB) (err error) {
	//create uuid
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	//read setting.json
	config := ReadJsonFormLocal()
	//set id
	uuid = config.Code + fmt.Sprintf("%04d", config.CurrentId)
	//update current id
	config.CurrentId = config.CurrentId + 1
	//write to file
	jsonFile, err := os.OpenFile("./entities/setting.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := json.Marshal(config)
	fmt.Println(string(byteValue))
	jsonFile.Write(byteValue)
	//set id
	book.ID = uuid
	return
}

func ReadJsonFormLocal() Config {
	// Read json file
	jsonFile, err := os.Open("./entities/setting.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	fmt.Println(string(byteValue))
	var config Config

	json.Unmarshal(byteValue, &config)

	fmt.Printf("config: %#v", config)

	return config

}

type Config struct {
	Code      string `json:"code"`
	CurrentId int    `json:"currentId"`
}

type Status struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Deleted gorm.DeletedAt
}

type Pageination struct {
	Page     int   `json:"page"`
	Limit    int   `json:"limit"`
	Pages    int   `json:"pages"`
	TotalRow int64 `json:"total_row"`
}

type Meta struct {
	Pageination   Pageination     `json:"pageination"`
	ProblemRecord []ProblemRecord `json:"problemrecord"`
}

type ProblemRecord struct {
	ID              string `gorm:"primaryKey" json:"id"`
	Agency          string `gorm:"type:varchar(50)" json:"agency"`
	Contact         string `gorm:"type:varchar(50)" json:"contact"`
	Informer        string `gorm:"type:varchar(50)" json:"informer"`
	Informermessage string `gorm:"type:varchar(50)" json:"informermessage"`
	System          string `gorm:"type:varchar(50)" json:"system"`
	Problemtype     string `gorm:"type:varchar(50)" json:"problemttype"`
	Level           string `gorm:"type:varchar(50)" json:"level"`
	Problem         string `gorm:"type:varchar(50)" json:"problem"`
	File_name       string `json:"file_name"`
	Path_file       string `json:"path_file"`
	File_extension  string `json:"file_extension"`
	File_size       int    `json:"file_size"`

	//Status
	Casuseproblem string `gorm:"type:varchar(255)" json:"casuseproblem"`
	Solution      string `gorm:"type:varchar(255)" json:"solution"`
	Suggestion    string `gorm:"type:varchar(255)" json:"suggestion"`
	Operator      string `gorm:"type:varchar(50)" json:"operator"`

	//CreatedAt   time.Time `gorm:"<-:crate;type:timestamp" json:"created_at"`
	//CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"created_at"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	SenderAt    time.Time `gorm:"column:sender_at;type:TIMESTAMP NULL;DEFAULT:NULL;" json:"-"`
	CompletedAt time.Time `gorm:"column:completed_at;type:TIMESTAMP NULL;DEFAULT:NULL;" json:"-"`

	Time           string        `json:"time"`
	Status         int           `json:"-"`
	Statuse        Statuse       `json:"status" gorm:"foreignKey:Status;references:Id"`
	Systems        []System      `json:"systems" gorm:"-"`
	Contacts       []Contact     `json:"contacts" gorm:"-"`
	Agencies       []Agency      `json:"agencies" gorm:"-"`
	Problemtypes   []Problemtype `json:"problemtype" gorm:"-"`
	Levels         []Level       `json:"levels" gorm:"-"`
	Users          []User        `json:"operators" gorm:"-"`
	Deleted        gorm.DeletedAt
	TmpCreateAt    int64 `json:"created_at" gorm:"-"`
	TmpSenderAt    int64 `json:"sender_at" gorm:"-"`
	TmpCompletedAt int64 `json:"completed_at" gorm:"-"`
}

type Statuse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MapSystem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MapContact struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MapAgnecy struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MapProblemType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MapLevel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Time int    `json:"time"`
}

type MapUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MapOperator struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
