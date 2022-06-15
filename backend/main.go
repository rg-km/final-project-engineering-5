package main

import (
	"FinalProject/handler"
	"FinalProject/impl/module"
	"FinalProject/utility"
	"log"
)

func main() {
	db := utility.ConnectDB()
	
	if err := utility.MigrationDB(db); err != nil {
		log.Panicln("Error when do migration:", err)
	}

	dataModule := module.NewDataModuleImpl(db)
	serviceModule := module.NewServiceModuleImpl(dataModule)

	handler.StartHandler(serviceModule)
}
