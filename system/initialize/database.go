package initialize

import (
	"errors"
	"reflect"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
)

// LoadAllDatabaseConnection load All Database Connection
func LoadAllDatabaseConnection(h *handler.Handler) error {
	cfg, err := h.GetConfig()
	if err != nil {
		return err
	}

	if cfg != nil {
		dbs := cfg.Databases
		e := reflect.ValueOf(&dbs).Elem()
		for i := 0; i < e.NumField(); i++ {
			// varName := e.Type().Field(i).Name
			// varType := e.Type().Field(i).Type
			dbConfig := e.Field(i).Interface()
			// fmt.Printf("%v %v %v\n", varName, varType, dbConfig)
			if dbConfig != nil {
				err := LoadDatabaseConnection(dbConfig.(config.Database), h)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// LoadDatabaseConnection load Database Connection using GORM
func LoadDatabaseConnection(dbConfig config.Database, h *handler.Handler) error {
	if h != nil {
		connString := dbConfig.Username + ":" + dbConfig.Password + "@(" + dbConfig.HostName + ")/" + dbConfig.DBName + "?" + dbConfig.Config
		if dbConfig.Driver == "sqlite3" {
			connString = dbConfig.HostName
		}

		dbCon, err := openDBConnection(dbConfig.Driver, connString)
		if err != nil {
			// h.GetLogger().Errorf("%s startup exception: [%s]", dbConfig.Driver, err.Error())
			// os.Exit(0)
			return err
		}
		dbClient, err := dbCon.DB()
		if err != nil {
			return err
		}
		dbClient.SetMaxIdleConns(dbConfig.MaxIdleConns)
		dbClient.SetMaxOpenConns(dbConfig.MaxOpenConns)
		if dbConfig.LogMode {
			dbCon.Logger = dbCon.Logger.LogMode(logger.Info)
		}

		err = dbClient.Ping()
		if err != nil {
			// h.GetLogger().Errorf("%s Ping error: [%s]", dbConfig.Driver, err.Error())
			return err
		}

		h.SetGormDB(dbConfig.ConnectionName, dbCon)
	}

	return nil
}

func openDBConnection(driverName, dataSourceName string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	if driverName == "postgres" {
		db, err = gorm.Open(postgres.Open(dataSourceName+" dbname=postgres"), &gorm.Config{})
	} else if driverName == "mysql" {
		db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	} else if driverName == "sqlite3" {
		db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	} else if driverName == "sqlserver" {
		db, err = gorm.Open(sqlserver.Open(dataSourceName), &gorm.Config{})
	} else {
		return nil, errors.New("database dialect is not supported")
	}
	if err != nil {
		return nil, err
	}
	return db, err
}
