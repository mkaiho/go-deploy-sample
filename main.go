package main

import (
	"github.com/mkaiho/go-deploy-sample/infrastructure"
)

func main() {
	infrastructure.NewServer().Run()
}
