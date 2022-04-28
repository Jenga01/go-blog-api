package Seeder

import (
	"first/Config"
	models2 "first/Model"
	"github.com/jinzhu/gorm"
	"log"
)

var users = []models2.User{
	models2.User{
		Name:     "Vardenis",
		Email:    "test@gmail.com",
		Password: "password",
	},
	models2.User{
		Name:     "Vardenis2",
		Email:    "test2@gmail.com",
		Password: "password",
	},
	models2.User{
		Name:     "Vardenis3",
		Email:    "test3@gmail.com",
		Password: "password",
	},
	models2.User{
		Name:     "Vardenis4",
		Email:    "test4@gmail.com",
		Password: "password",
	},
}

var articles = []models2.Article{
	models2.Article{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models2.Article{
		Title:   "Title 2",
		Content: "Hello world 2",
	},

	models2.Article{
		Title:   "Title 3",
		Content: "Hello world 3",
	},
	models2.Article{
		Title:   "Title 4",
		Content: "Hello world 4",
	},
}

var comments = []models2.Comment{
	models2.Comment{
		Title:   "Comment 1",
		Content: "Hello comment 1",
	},
	models2.Comment{
		Title:   "Comment 2",
		Content: "Hello comment 2",
	},

	models2.Comment{
		Title:   "Comment 3",
		Content: "Hello comment 3",
	},
	models2.Comment{
		Title:   "Comment 4",
		Content: "Hello comment 4",
	},
	models2.Comment{
		Title:   "Comment 5",
		Content: "Hello comment 5",
	},
	models2.Comment{
		Title:   "Comment 6",
		Content: "Hello comment 6",
	},

	models2.Comment{
		Title:   "Comment 7",
		Content: "Hello comment 7",
	},
	models2.Comment{
		Title:   "Comment 8",
		Content: "Hello comment 8",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models2.Comment{}, &models2.Article{}, &models2.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = Config.DB.Debug().AutoMigrate(&models2.User{}, &models2.Article{}, &models2.Comment{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models2.Article{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models2.Comment{}).AddForeignKey("article_id", "articles(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models2.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		articles[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models2.Article{}).Create(&articles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed articles table: %v", err)
		}

		comments[i].ArticleID = articles[i].ID
		err = db.Debug().Model(&models2.Comment{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed comments table: %v", err)
		}
	}

}
