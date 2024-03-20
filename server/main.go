package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pbuser "grpc/user"
	"server/api/user"
	"server/config"
	"server/middleware"
	"server/models"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func server() {
	flag.Parse()
	arg := flag.Arg(0)
	env := config.Config{}.GetEnvVar() // cargar variables de entorno

	switch arg {
	case "start":
		models.DialDb(env.Dsn, "./logs/error.log") // conectar con la base de datos (sqlite)
		middleware.SetToken(env.Key)               // cargando el token en memoria

		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", env.Host, env.Port))
		if err != nil {
			log.Fatalf("[-] Failed to listen: %s\n", err)
		}

		/*
			creds, err := credentials.NewClientTLSFromFile("./certs/server_cert.pem", "./certs/server_key.pem")
			if err != nil {
				panic(err)
			}
		*/

		opts := []grpc.ServerOption{
			grpc.Creds(insecure.NewCredentials()),
			grpc.ChainUnaryInterceptor(
				auth.UnaryServerInterceptor(middleware.ValidateAuthToken),
			),
		}

		grpcUser := grpc.NewServer(opts...)
		apiUser := user.NewUser()
		pbuser.RegisterUserServiceServer(grpcUser, apiUser)

		fmt.Println("\n*********************************")
		fmt.Println("********* GRPC Server ***********")
		fmt.Println("*********************************")
		fmt.Printf("\n[+] Server running at -> %s:%d time: %s\n", env.Host, env.Port, time.Now().Format("15:04:05"))

		if err := grpcUser.Serve(listen); err != nil {
			log.Fatalf("[-] Failed to serve %s\n", err)
		}

	case "migrate":
		models.Migrate(env.Dsn)

	default:
		fmt.Println("[-] No action taken")
	}

}

func main() {
	server()
}
