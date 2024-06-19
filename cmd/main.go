package main

import ("app/config"
		"app/pkg/database"
		"log"
		"fmt"
)



func main() {

	log.Println("Starting procces...")

	cfg :=config.NewConfig()

	log.Println("Succesfully loaded config!")

	conn, err :=database.Conn(cfg)

		if err != nil{
			log.Println("Error on connecting to database:", err)
		}

		log.Println("Succesfully connected to database")

		fmt.Println(conn)

}