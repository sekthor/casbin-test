package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")

	if err != nil {
		log.Fatalf("could not create enforcer: %v", err)
	}

	ok, err := e.Enforce("sekthor", "data1", "read")

	if err != nil {
		log.Fatalf("could not enforce policy: %v", err)
	}

	fmt.Printf("authorization result: %t", ok)
}
