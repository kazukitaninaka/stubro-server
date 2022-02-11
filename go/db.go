package openapi

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Type struct {
	*gorm.Model
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"unique"`
}

type TypeMentor struct {
	*gorm.Model
	MentorId uint `json:"mentorId"`
	TypeId   uint `json:"typeId"`
}

var Db *gorm.DB

func InitDb() {
	dsn := "kazukitaninaka:db-password@/test_db?charset=utf8&parseTime=True"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	} else {
		fmt.Println("Sucessdully connected to database")
	}
	// setup logger
	Db.Logger.LogMode(logger.Info)

	db, err := Db.DB()
	defer db.Close()

	// Auto Migration
	// Db.AutoMigrate(&Mentor{}, &Type{})

	// err = Db.SetupJoinTable(&Mentor{}, "Types", &TypeMentor{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// seed(Db)
}

func seed(d *gorm.DB) {
	mentors := []Mentor{
		// {
		// 	Username:    "Kazuki",
		// 	Email:       "kazuki@gmail.com",
		// 	Image:       "https://placehold.jp/150x150.png",
		// 	ClickCount:  0,
		// 	Destination: "アメリカ",
		// 	EduOrg:      "Reading High School",
		// 	Term:        "半年〜1年",
		// 	Type: []Type{
		// 		{ID: 7, Name: "大学留学"},
		// 		{ID: 10, Name: "私費留学"},
		// 		{ID: 9, Name: "寮"},
		// 	},
		// 	Price: 100,
		// },
		{
			Username:    "Eren",
			Email:       "eren@gmail.com",
			Image:       "https://placehold.jp/150x150.png",
			ClickCount:  0,
			Destination: "イギリス",
			EduOrg:      "University of Oxford",
			Term:        "4年",
			Type: []Type{
				{ID: 7, Name: "大学留学"},
				{ID: 10, Name: "私費留学"},
				{ID: 9, Name: "寮"},
			},
			Price: 3000,
		},
		{
			Username:    "Armin",
			Email:       "armin@gmail.com",
			Image:       "https://placehold.jp/150x150.png",
			ClickCount:  0,
			Destination: "デンマーク",
			EduOrg:      "University of Copenhagen",
			Term:        "半年〜1年",
			Type: []Type{
				{ID: 7, Name: "大学留学"},
				{ID: 6, Name: "交換留学"},
			},
			Price: 2500,
		},
		{
			Username:    "Mikasa",
			Email:       "mikasa@gmail.com",
			Image:       "https://placehold.jp/150x150.png",
			ClickCount:  0,
			Destination: "アメリカ",
			EduOrg:      "University of California, Los Angeles",
			Term:        "半年〜1年",
			Type: []Type{
				{ID: 7, Name: "大学留学"},
				{ID: 6, Name: "交換留学"},
				{ID: 9, Name: "寮"},
				{ID: 8, Name: "奨学金利用"},
			},
			Price: 2000,
		},
	}

	for _, m := range mentors {
		d.Create(&m)
	}

}
