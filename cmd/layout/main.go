package main

import (
	"fmt"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout"
)

func main(){
	res := layout.NewApp()
	fmt.Println("run app", res)
}