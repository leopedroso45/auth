package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB *gorm.DB
}

var (
	//Server now implements the modelInterface, so he can define its methods
	Model modelInterface = &Server{}
)

type modelInterface interface {
	//db initialization
	Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error)

	//user methods
	ValidateEmail(string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)

}

func (s *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	s.DB.Debug().AutoMigrate(
		&User{},
	)
	return s.DB, nil
}
