package main

import (
	"fmt"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout"
)

func main() {
	res := layout.NewApp()
	res.Run()
	fmt.Println("run app", res)
}
