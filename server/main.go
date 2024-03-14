package main

import (
	"flag"
	"fmt"
	

	"github.com/MBI-88/fiber-grpc-crud/server/config"
	"github.com/MBI-88/fiber-grpc-crud/server/models"
	
)

func server() {
	flag.Parse()
	arg := flag.Arg(0)
	env := config.Config{}.GetEnvVar() // cargar variables de entorno
	
	switch arg {
	case "start":
		models.DialDb(env.Dsn, "./logs/error.log") // conectar con la base de datos (sqlite)

	case "migrate":
		models.Migrate(env.Dsn)

	default:
		fmt.Println("[-] No action taken")
	}

}

func main() {
	server()
}
