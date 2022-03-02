package seed

import (
	"log"

	"github.com/aniruddha2000/blogger/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username: "test1",
		Email:    "test@test1.com",
		Password: "1234",
	},
	models.User{
		Username: "test2",
		Email:    "test@test2.com",
		Password: "1234",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "title 1",
		Content: "Content 1",
	},
	models.Post{
		Title:   "title2",
		Content: "Content 2",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.AutoMigrate(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attatching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed post table: %v", err)
		}
	}
}
