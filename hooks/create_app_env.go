package hooks

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type CreateAppEnvHook struct {
}

func (c *CreateAppEnvHook) Name() string {
	return "create_app_env"
}

func (c *CreateAppEnvHook) Exec() {
	err := os.MkdirAll("/data", fs.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("/data/appenv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	env := os.Getenv("env")
	fmt.Fprintln(f, "env=", env)
	region := os.Getenv("region")
	fmt.Fprintln(f, "region=", region)
	zone := os.Getenv("zone")
	fmt.Fprintln(f, "zone=", zone)
}
