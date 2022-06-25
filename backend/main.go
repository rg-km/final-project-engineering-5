package main

import (
	"FinalProject/handler"
	"FinalProject/impl/module"
	"FinalProject/utility"
	"log"
	"sync"
)

func main() {
	db := utility.ConnectDB()
	
	if err := utility.MigrationDB(db); err != nil {
		log.Panicln("Error when do migration:", err)
	}

	mu := &sync.Mutex{}
	dataModule := module.NewDataModuleImpl(db, mu)
	serviceModule := module.NewServiceModuleImpl(dataModule)

	handler.StartHandler(serviceModule)
}
