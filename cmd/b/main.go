package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("lishimeng",
		buildscript.Application{
			Name:    "track",
			AppPath: "cmd/track",
			HasUI:   false,
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
