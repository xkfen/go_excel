package main

import (
	"fmt"
	"gcoresys/common/mysql"
	"gexcel/db/config"
	"flag"
	"gcoresys/common"
	"gexcel/model"
)

func main() {
	common.DefineDbMigrateCommonFlag()
	env := common.DefineRunTimeCommonFlag()
	action := flag.Lookup("action").Value.String()
	switch action {
	case "create":
		doCreate(env)
	case "drop":
		doDrop(env)
	case "migrate":
		doMigrate(env)
	}
}

func doCreate(env string) {
	fmt.Println("do create")
	dbConfig := config.GetUserInfoDbConfig(env)
	mysql.CreateDB(dbConfig)
}

func doDrop(env string) {
	fmt.Println("do drop")
	dbConfig := config.GetUserInfoDbConfig(env)
	mysql.DropDB(dbConfig)
}

func doMigrate(env string) {
	fmt.Println("do migrate")
	config.GetUserInfoDbConfig(env)
	db := config.GetDb()
	db.AutoMigrate(
					&model.UserInfo{},
	               )
}
