package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         int      `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"type:varchar(50)" json:"name"`
	Nickname   string   `gorm:"type:varchar(50)" json:"nickname"`
	Systems    string   `json:"systems"`
	ListSystem []System `json:"listsystems" gorm:"-"`
}

type MetaUser struct {
	Pageination Pageination `json:"pageination"`
	Users       []User      `json:"operator"`
}

type System struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Problemtype struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Level struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time int    `json:"time"`
}

type Contact struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Agency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// type ProblemRecord struct {
// 	ID     string `gorm:"primaryKey" json:"id"`
// 	Agency string `gorm:"type:varchar(50)" json:"agency"`
// 	//Agencys []Agency `gorm:"foreignKey:AgencyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL ;" json:"agency"`
// 	Contact         string `gorm:"type:varchar(50)" json:"contact"`
// 	Informer        string `gorm:"type:varchar(50)" json:"informer"`
// 	Informermessage string `gorm:"type:varchar(50)" json:"informermessage"`
// 	System          string `gorm:"type:varchar(50)" json:"system"`
// 	//System 		     *System  `gorm:"foreignKey:SystemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"system"`
// 	Problemtype    string `gorm:"type:varchar(50)" json:"problemttype"`
// 	Level          string `gorm:"type:varchar(50)" json:"level"`
// 	Problem        string `gorm:"type:varchar(50)" json:"problem"`
// 	File_name      string `json:"file_name"`
// 	Path_file      string `json:"path_file"`
// 	File_extension string `json:"file_extension"`
// 	File_size      int    `json:"file_size"`
// 	Status         int    `json:"status"`
// 	//Status
// 	Statusname    string `json:"statusname"`
// 	Casuseproblem string `gorm:"type:varchar(255)" json:"casuseproblem"`
// 	Solution      string `gorm:"type:varchar(255)" json:"solution"`
// 	Suggestion    string `gorm:"type:varchar(255)" json:"suggestion"`
// 	Operator      string `gorm:"type:varchar(50)" json:"operator"`

// 	CreatedAt time.Time `gorm:"<-:crate;type:timestamp" json:"created_at"`
// 	//SenderAt     	time.Time `gorm:"<-:update;type:timestamp" json:"sender_at"`
// 	SenderAt time.Time `gorm:"column:sender_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"sender_at"`
// 	// UpdatedAt time.Time `gorm:"<-:update;type:timestamp;column:sender_at" json:"updated_at"`
// 	CompletedAt time.Time `gorm:"column:completed_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"completed_at"`
// 	Time        int       `json:"timesla"`
// }

func (book *ProblemRecord) BeforeCreate(tx *gorm.DB) (err error) {
	uuidWithHyphen := uuid.New()
	//limit string length to 5 characters and remove hyphen (-) from uuid string using strings.ReplaceAll() function in Go language.
	//start useing text "TTRS" + uuid string to create unique id for each book. For example, TTRS-1a2b3c4d5e6f7g8h9i0j
	//heeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	uuid = uuid[:3]
	uuid = "TCC" + uuid
	book.ID = uuid
	return
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	ID     string `gorm:"primaryKey" json:"id"`
	Agency string `gorm:"type:varchar(50)" json:"agency"`
	//Agencys []Agency `gorm:"foreignKey:AgencyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL ;" json:"agency"`
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

	CreatedAt   time.Time `gorm:"<-:crate;type:timestamp" json:"created_at"`
	SenderAt    time.Time `gorm:"column:sender_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"sender_at"`
	CompletedAt time.Time `gorm:"column:completed_at;type:TIMESTAMP;DEFAULT:CURRENT_TIMESTAMP;not null;" json:"completed_at"`

	Time         int           `json:"timesla"`
	Status       int           `json:"-"`
	Statuse      Statuse       `json:"status" gorm:"foreignKey:Status;references:Id"`
	Systems      []System      `json:"systems" gorm:"-"`
	Contacts     []Contact     `json:"contacts" gorm:"-"`
	Agencies     []Agency      `json:"agencies" gorm:"-"`
	Problemtypes []Problemtype `json:"problemtype" gorm:"-"`
	Levels       []Level       `json:"levels" gorm:"-"`
	Users        []User        `json:"operators" gorm:"-"`
	//asdasdljsdl;kfj;ajriogjae[o'rjgadjfg;sldkfkjg]asdasdasdasdawdawdasdawdawdawdadawdawdasdawdadasdawdad
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