package pkg

import (
	"github.com/RogerDurdn/users/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type Storage interface {
	FindUserById(id int) (*model.User, error)
	FindUserByUserName(userName string) (*model.User, error)
	CreateOrUpdateUser(user *model.User) (*model.User, error)
	DeleteUserById(id int) error
}

type postgresStorage struct {
	db *gorm.DB
}

var config = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "data.",
		SingularTable: false,
	},
}

func NewPostgresStorage() *postgresStorage {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Panicln(err)
	}
	return &postgresStorage{db: db}
}

func (p *postgresStorage) FindUserById(id int) (*model.User, error) {
	var user model.User
	return &user, p.db.First(&user, id).Error
}

func (p *postgresStorage) FindUserByUserName(userName string) (*model.User, error) {
	var user model.User
	return &user, p.db.Where("user_name = ?", userName).First(&user).Error
}

func (p *postgresStorage) CreateOrUpdateUser(user *model.User) (*model.User, error) {
	return user, p.db.Create(user).Error
}

func (p *postgresStorage) DeleteUserById(id int) error {
	return p.db.Delete(model.User{}, id).Error
}
