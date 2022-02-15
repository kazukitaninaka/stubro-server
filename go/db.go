package openapi

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TypeMentor struct {
	*gorm.Model
	MentorId uint `json:"mentorId"`
	TypeId   uint `json:"typeId"`
}

type TermMentor struct {
	*gorm.Model
	MentorId uint `json:"mentorId"`
	TermId   uint `json:"termId"`
}

var Db *gorm.DB

func InitDb() {
	var err error
	dsn := "kazukitaninaka:db-password@/test_db?charset=utf8&parseTime=True"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	} else {
		fmt.Println("Sucessfully connected to database")
	}
	// setup logger
	Db.Logger.LogMode(logger.Info)

	// db, err := Db.DB()
	// defer db.Close()

	// Auto Migration
	// Db.AutoMigrate(&Consultation{})

	// err = Db.SetupJoinTable(&Mentor{}, "Types", &TypeMentor{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// seed(Db)
}

func seed(d *gorm.DB) {
	// terms := []Term{
	// 	{Name: "〜1週間"},
	// 	{Name: "1週間〜1ヶ月"},
	// 	{Name: "1ヶ月〜3ヶ月"},
	// 	{Name: "3ヶ月〜半年"},
	// 	{Name: "半年〜1年"},
	// 	{Name: "2年"},
	// 	{Name: "3年"},
	// 	{Name: "4年"},
	// }

	// for _, t := range terms {
	// 	d.Create(&t)
	// }
	// mentors := []Mentor{
	// 	{
	// 		Username:    "Kazuki",
	// 		Email:       "kazuki@gmail.com",
	// 		Image:       "https://placehold.jp/150x150.png",
	// 		ClickCount:  0,
	// 		Destination: "アメリカ",
	// 		EduOrg:      "Reading High School",
	// 		Term:        Term{ID: 5, Name: "半年〜1年"},
	// 		Types: []Type{
	// 			{ID: 7, Name: "大学留学"},
	// 			{ID: 10, Name: "私費留学"},
	// 			{ID: 9, Name: "寮"},
	// 		},
	// 		Price: 100,
	// 	},
	// 	{
	// 		Username:    "Eren",
	// 		Email:       "eren@gmail.com",
	// 		Image:       "https://placehold.jp/150x150.png",
	// 		ClickCount:  0,
	// 		Destination: "イギリス",
	// 		EduOrg:      "University of Oxford",
	// 		Term:        Term{ID: 8, Name: "4年"},
	// 		Types: []Type{
	// 			{ID: 7, Name: "大学留学"},
	// 			{ID: 10, Name: "私費留学"},
	// 			{ID: 9, Name: "寮"},
	// 		},
	// 		Price: 3000,
	// 	},
	// 	{
	// 		Username:    "Armin",
	// 		Email:       "armin@gmail.com",
	// 		Image:       "https://placehold.jp/150x150.png",
	// 		ClickCount:  0,
	// 		Destination: "デンマーク",
	// 		EduOrg:      "University of Copenhagen",
	// 		Term:        Term{ID: 5, Name: "半年〜1年"},
	// 		Types: []Type{
	// 			{ID: 7, Name: "大学留学"},
	// 			{ID: 6, Name: "交換留学"},
	// 		},
	// 		Price: 2500,
	// 	},
	// 	{
	// 		Username:    "Mikasa",
	// 		Email:       "mikasa@gmail.com",
	// 		Image:       "https://placehold.jp/150x150.png",
	// 		ClickCount:  0,
	// 		Destination: "アメリカ",
	// 		EduOrg:      "University of California, Los Angeles",
	// 		Term:        Term{ID: 5, Name: "半年〜1年"},
	// 		Types: []Type{
	// 			{ID: 7, Name: "大学留学"},
	// 			{ID: 6, Name: "交換留学"},
	// 			{ID: 9, Name: "寮"},
	// 			{ID: 8, Name: "奨学金利用"},
	// 		},
	// 		Price: 2000,
	// 	},
	// }

	// for _, m := range mentors {
	// 	d.Create(&m)
	// }

}
