package main

import (
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

func main() {
	host := "0.0.0.0"
	if host_env, ok := os.LookupEnv("PROXY_HOST"); ok {
		host = host_env
	}

	port := "8080"
	if port_env, ok := os.LookupEnv("PROXY_PORT"); ok {
		port = port_env
	}

	options := []socks5.Option{
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	}

	username := os.Getenv("PROXY_USER")
	password := os.Getenv("PROXY_PASS")
	if username != "" {
		log.Println("Using username/password auth")
		store := socks5.StaticCredentials{
			username: password,
		}

		options = append(options, socks5.WithAuthMethods([]socks5.Authenticator{socks5.UserPassAuthenticator{Credentials: store}}))
		options = append(options, socks5.WithCredential(store))
	} else {
		log.Println("Using no auth")
		options = append(options, socks5.WithAuthMethods([]socks5.Authenticator{socks5.NoAuthAuthenticator{}}))
	}

	// Create a SOCKS5 server
	server := socks5.NewServer(
		options...,
	)

	log.Println("Listening on", host+":"+port)
	if err := server.ListenAndServe("tcp", host+":"+port); err != nil {
		panic(err)
	}
}
