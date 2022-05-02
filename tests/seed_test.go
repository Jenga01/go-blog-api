package tests

import (
	"first/Config"
	models2 "first/Model"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

func TestMain(m *testing.M) {

	err = godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	TestDbDriver := os.Getenv("TestDbDriver")
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))

	Config.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Println("Opening status:", err)
	} else {
		fmt.Println("server opened")
	}
}

func refreshUserTable() error {
	err := Config.DB.DropTableIfExists(&models2.User{}).Error
	if err != nil {
		return err
	}
	err = Config.DB.AutoMigrate(&models2.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func refreshCommentsTable() error {
	err := Config.DB.DropTableIfExists(&models2.Comment{}).Error
	if err != nil {
		return err
	}
	err = Config.DB.AutoMigrate(&models2.Comment{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models2.User, error) {

	err := refreshUserTable()
	if err != nil {
		return models2.User{}, err
	}

	user := models2.User{
		Name:     "Vardenis1",
		Email:    "test@mail.com",
		Password: "password",
	}

	err = Config.DB.Model(&models2.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedOneArticle() (models2.Article, error) {

	err := refreshUserArticlesCommentsTable()
	if err != nil {
		return models2.Article{}, err
	}
	user := models2.User{
		Name:     "Vardenis Pavardenis",
		Email:    "test@gmail.com",
		Password: "password",
	}
	article := models2.Article{
		ID:       1,
		Title:    "Title",
		Content:  "Content",
		AuthorID: user.ID,
	}
	return article, nil
}

func refreshUserArticlesCommentsTable() error {

	err := Config.DB.DropTableIfExists(&models2.Comment{}, &models2.Article{}, &models2.User{}).Error
	if err != nil {
		return err
	}
	err = Config.DB.AutoMigrate(&models2.User{}, &models2.Article{}, &models2.Comment{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

/*func seedArticlesTable() ([]models.User, []models.Article, error) {

}
*/
func seedOneUserAndOneArticle() (models2.Article, error) {

	err := refreshUserArticlesCommentsTable()
	if err != nil {
		return models2.Article{}, err
	}
	user := models2.User{
		Name:     "Vardenis Pavardenis",
		Email:    "test@gmail.com",
		Password: "password",
	}
	err = Config.DB.Model(&models2.User{}).Create(&user).Error
	if err != nil {
		return models2.Article{}, err
	}
	article := models2.Article{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = Config.DB.Model(&models2.Article{}).Create(&article).Error
	if err != nil {
		return models2.Article{}, err
	}
	return article, nil
}
func seedComments() ([]models2.Comment, error) {

	err := refreshCommentsTable()
	article, _ := seedOneArticle()
	if err != nil {
		return nil, err
	}

	comments := []models2.Comment{
		models2.Comment{
			Title:     "Title 1",
			Content:   "Content content content",
			ArticleID: article.ID,
		},
		models2.Comment{
			Title:     "Title 2",
			Content:   "Content2 content2 content2",
			ArticleID: article.ID,
		},
	}

	for i, _ := range comments {
		err := Config.DB.Model(&models2.Comment{}).Create(&comments[i]).Error
		if err != nil {
			return nil, err
		}
	}
	return comments, nil
}

func seedUsersArticlesComments() ([]models2.User, []models2.Article, []models2.Comment, error) {

	var err error

	if err != nil {
		return []models2.User{}, []models2.Article{}, []models2.Comment{}, nil
	}
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
	for i, _ := range users {
		err = Config.DB.Model(&models2.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		articles[i].AuthorID = users[i].ID

		err = Config.DB.Model(&models2.Article{}).Create(&articles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
		comments[i].ArticleID = articles[i].ID
		err = Config.DB.Debug().Model(&models2.Comment{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed comments table: %v", err)
		}
	}
	return users, articles, comments, nil
}
