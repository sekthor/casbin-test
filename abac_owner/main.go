package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

type Resource struct {
	Name  string
	Owner string
}

func main() {
	e, err := casbin.NewEnforcer("model.conf")

	if err != nil {
		log.Fatalf("could not create enforcer: %v", err)
	}

	data1 := Resource{
		Name:  "resource1",
		Owner: "sekthor",
	}

	ok, err := e.Enforce("sekthor", data1, "read")

	if err != nil {
		log.Fatalf("could not enforce policy: %v", err)
	}

	fmt.Printf("authorization result: %t", ok)

}
