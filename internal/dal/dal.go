package dal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormlog "gorm.io/gorm/logger"
)

var db *gorm.DB

type DB struct {
	DB *gorm.DB
}

// Init creates a db cnx with specified config
func Init() (*DB, error) {
	logrus.Info("Initializing datasource ...")
	var err error
	cnx := os.Getenv("UNAME") + ":" + os.Getenv("PASS")
	schema, exist := os.LookupEnv("SCHEMA")
	if !exist {
		schema = "stacke"
	}
	orm := &DB{DB: nil}
	dataSourceName := cnx + "@tcp(" + os.Getenv("SERVER") + ":3306)/" + schema + "?parseTime=true"

	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: gormlog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      gormlog.Info,
				Colorful:      true,
			},
		)})

	if err != nil {
		fmt.Println(err)
		logrus.Fatal("Failed to connect database")
	}
	orm.DB = db
	return orm, err
}
