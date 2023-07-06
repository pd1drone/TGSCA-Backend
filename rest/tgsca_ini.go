package rest

import (
	"fmt"
	"tgsca/database"

	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

type TGSCAConfiguration struct {
	TGSCAdb    *sqlx.DB
	FrontEndIP string
	PortNumber string
}

func New() (*TGSCAConfiguration, error) {

	// read config file
	cfg, err := ini.Load("/root/TGSCA-Backend/config.ini")
	if err != nil {
		return nil, fmt.Errorf("Fail to read file: %v", err)
	}

	dbSection := cfg.Section("db")
	user := dbSection.Key("user").String()
	password := dbSection.Key("password").String()
	dbhost := dbSection.Key("dbhost").String()
	dbport := dbSection.Key("dbport").String()
	dbname := dbSection.Key("dbname").String()

	Frontend := cfg.Section("frontend")
	ip := Frontend.Key("ip").String()
	port := Frontend.Key("port").String()

	caddb, err := database.InitializeTGSCADatabase(dbname, user, password, dbhost, dbport)
	if err != nil {
		return nil, err
	}

	return &TGSCAConfiguration{
		caddb,
		ip,
		port,
	}, nil
}
