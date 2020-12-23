package dal

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// TODO cnx settings
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

func GetQueryBounds(first *int, after *string) (int64, int, error) {
	// prep query - query start
	start := int64(0)
	limit := 0
	if after != nil {
		decoded, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return 0, 0, err
		}
		start, err = strconv.ParseInt(string(decoded), 10, 0)
		if err != nil {
			return 0, 0, err
		}
	}

	// prep query  - limit
	if first != nil {
		limit = *first
	}
	if limit > 50 {
		logrus.Warn("Limit requested exceeds maximum 50")
		limit = 50
	}

	return start, limit, nil
}
