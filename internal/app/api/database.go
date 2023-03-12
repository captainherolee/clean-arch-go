package api

import (
	"os"
	"sync"

	"github.com/captainherolee/go-rest/core"
	"github.com/jinzhu/gorm"
)

type singletonDatabase struct {
	Database *gorm.DB
}

var (
	instanceDatabase *singletonDatabase
	onceDatabase     sync.Once
)

func getDatabase() *gorm.DB {
	return getDbInstance().Database
}

func getDbInstance() *singletonDatabase {
	onceDatabase.Do(func() {

		instanceDatabase = &singletonDatabase{}
		dbString := os.Getenv("DATA_SOURCE_NAME")
		core.Logger().Infoln(dbString)

		var err error
		instanceDatabase.Database, err = gorm.Open("postgres", dbString)
		if err != nil {
			core.Logger().Fatal(err)
		}
		initDatabase(instanceDatabase.Database)
	})
	return instanceDatabase
}

func initDatabase(db *gorm.DB) {
	db.DB().SetMaxIdleConns(0)
	// db.SetLogger(core.Logger())
	db.LogMode(true)
}
