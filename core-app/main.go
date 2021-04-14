package main

// note: if you change this to "github.com/silbinarywolf/go-internal-example/other-service/internal/app"
// compilation will fail, because only "other-service" can import things from the "internal" folder
//
// The error returned is:
// package github.com/silbinarywolf/go-internal-example/core-app
// 		main.go:5:8: use of internal package github.com/silbinarywolf/go-internal-example/other-service/internal/app not allowed
import "github.com/silbinarywolf/go-internal-example/core-app/internal/app"

func main() {
	app.Start()
}
