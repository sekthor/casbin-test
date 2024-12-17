package main

import (
	"embed"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

//go:embed model.conf
var assets embed.FS

func main() {
	f, err := assets.ReadFile("model.conf")
	if err != nil {
		log.Fatal(err)
	}

	a := fileadapter.NewAdapter("policy.csv")

	m, err := model.NewModelFromString(string(f))
	if err != nil {
		log.Fatal(err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatal(err)
	}

	ok, err := e.Enforce("sekthor", "data1", "read")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("evaluation result:", ok)
}
