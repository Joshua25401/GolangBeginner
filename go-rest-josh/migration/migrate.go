package migration

import (
	"log"

	"github.com/joshua25401/go-test/database"
	"github.com/joshua25401/go-test/model/entity"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&entity.User{},
	)

	// if there is error
	// do panic()
	if err != nil {
		panic(err.Error())
	}

	// if there is no error
	// do log
	log.Println("Migrating table....")
}
