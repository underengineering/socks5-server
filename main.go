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

	store := socks5.StaticCredentials{
		os.Getenv("PROXY_USER"): os.Getenv("PROXY_PASS"),
	}

	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithAuthMethods([]socks5.Authenticator{socks5.UserPassAuthenticator{Credentials: store}}),
		socks5.WithCredential(store),
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	log.Println("Listening on", host+":"+port)

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "127.0.0.1:8005"); err != nil {
		panic(err)
	}
}
