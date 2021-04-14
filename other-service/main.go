package main

// note: if you change this to "github.com/silbinarywolf/go-internal-example/core-app/internal/app"
// compilation will fail, because only "core-app" can import things from the "internal" folder
import "github.com/silbinarywolf/go-internal-example/other-service/internal/app"

func main() {
	app.Start()
}
