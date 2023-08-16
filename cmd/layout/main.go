package main

import (
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout"
)

func main() {
	res := layout.NewApp()
	res.Run()
}
