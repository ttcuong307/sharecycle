package database

import (
	"fmt"
	"net/url"
	"sharecycle/configs"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBV1 struct {
	DB *gorm.DB
}

func NewDB(config *configs.Config) *DBV1 {
	options := url.Values{
		"interpolateParams": []string{"true"},
		"collation":         []string{"utf8mb4_bin"},
		"parseTime":         []string{"true"},
		"loc":               []string{"UTC"},
	}

	dsn := fmt.Sprintf(configs.DnsFormat,
		config.DBConfigs.UserName,
		config.DBConfigs.Password,
		config.DBConfigs.Host,
		config.DBConfigs.Port,
		config.DBNames.V1,
		options.Encode())

	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true, // Don't ping DB
	})

	if err != nil {
		logrus.Fatalln(err)
	}

	// TODO : Add gorm logger
	// db.Logger = NewGormLogger(logger.Info)
	logrus.Println("connected mysql")

	return &DBV1{db}
}
