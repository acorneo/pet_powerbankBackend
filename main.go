package main

import (
	"acorneo/powerbanks/utils"
	"log"
)

func main() {
	db, err := utils.OpenConnection()
	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}
	defer utils.CloseConnection(db)

	err = utils.InsertExcuse(db, "Никогда не говори Никита.")
	if err != nil {
		log.Println("Failed to insert excuse: ", err)
	}
}
