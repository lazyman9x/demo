package main

import (
	"fmt"

	"github.com/lazyman9x/demo/v1/config"
	"github.com/lazyman9x/demo/v1/database"
	"github.com/lazyman9x/demo/v1/mode"
	"github.com/lazyman9x/demo/v1/model"
	"github.com/lazyman9x/demo/v1/router"
	"github.com/lazyman9x/demo/v1/runner"
)

var (
	// Version the version of Gotify.
	Version = "unknown"
	// Commit the git commit hash of this version.
	Commit = "unknown"
	// BuildDate the date on which this binary was build.
	BuildDate = "unknown"
	// Mode the build mode
	Mode = mode.Dev
)

func main() {
	vInfo := &model.VersionInfo{Version: Version, Commit: Commit, BuildDate: BuildDate}
	fmt.Print("START APPLICATION !!!")
	conf := config.Get()
	db, _ := database.New(conf.Database.Dialect, conf.Database.Connection, conf.DefaultUser.Name, conf.DefaultUser.Pass, conf.PassStrength, true)
	fmt.Print("Start application!!!")
	engine, _ := router.Create(db, vInfo, conf)
	runner.Run(engine, conf)
}
