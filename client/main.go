package main

import (
	"client/cmd"
	"flag"
	"fmt"
	"log"
	"os"

	"strconv"

	pbuser "grpc/user"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	name     *string
	password *string
	website  *string
	dni      *string
	id       *string
	address  *string
	phone    *string
)

func init() {
	name = flag.String("name", "", "set name")
	password = flag.String("password", "", "set password")
	website = flag.String("website", "", "set website")
	dni = flag.String("dni", "", "set dni")
	id = flag.String("id", "", "set id")
	address = flag.String("address", "", "set address")
	phone = flag.String("phone", "", "set phone")

	flag.Usage = func() {
		info := fmt.Sprintln(`
		[*] Variable info:
		\nname=John
		\npassword=john@123*
		\ndni=1234567A
		\naddress=Stree...
		\nphone=+461234568
		\nArguments:\nOptions:\n
		CreateUser: (create)\nUpdateUser: (update)\nDeleteUser: (delete)\nGetUsers: (get)\nExit: (exit)
		`)
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

func client() {
	opt := flag.Arg(0)
	env := cmd.Config{}.GetEnvVar()
	addr := fmt.Sprintf("%s:%d", env.Host, env.Port)
	cmd.SetToken(env.Key)
	/*
		creds, err := credentials.NewClientTLSFromFile("./certs/client_cert.pem", "grpc.example.com")
		if err != nil {
			log.Fatalf("%s", err)
		}
	*/

	settings := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(cmd.UnaryAuthInterceptor),
	}

	conn, err := grpc.Dial(addr, settings...)
	if err != nil {
		log.Fatalf("%s", err)
	}

	u := pbuser.NewUserServiceClient(conn)
	fmt.Printf("[+] GRPC operation %s running...\n",opt)
	switch opt {
	case "create":
		fmt.Println(cmd.CreateUser(
			*name,
			*phone,
			*website,
			*password,
			*dni,
			*address,
			u,
		))

	case "update":
		idInt, err := strconv.Atoi(*id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Println(cmd.UpdateUser(
			uint32(idInt),
			*name,
			*phone,
			*website,
			*password,
			*dni,
			*address,
			u,
		))
	case "delete":
		idInt, err := strconv.Atoi(*id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		fmt.Println(cmd.DeletUser(uint32(idInt), u))

	case "get":
		cmd.GetUsers(u)
	default:
		os.Exit(0)

	}

}

func main() {
	flag.Parse()
	client()

}
