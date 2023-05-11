package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/track/internal/api"
	"github.com/lishimeng/track/internal/etc"
	"github.com/lishimeng/track/internal/process"
	"time"
)

//import _ "github.com/lib/pq"

import _ "github.com/lishimeng/track/ddd/g7"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}

		builder.
			SetWebLogLevel("debug").
			PrintVersion().
			EnableWeb(etc.Config.Web.Listen, api.Route).
			ComponentAfter(process.RunTask)

		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
