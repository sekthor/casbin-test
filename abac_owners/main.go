package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

type Resource struct {
	Name   string
	Owners []string
}

func (r Resource) IsOwnedBy(sub string) bool {
	for _, owner := range r.Owners {
		if owner == sub {
			return true
		}
	}
	return false
}

func main() {
	e, err := casbin.NewEnforcer("model.conf")

	if err != nil {
		log.Fatalf("could not create enforcer: %v", err)
	}

	data1 := Resource{
		Name:   "resource1",
		Owners: []string{"sekthor", "testuser"},
	}

	ok, err := e.Enforce("sekthor", data1, "read")

	if err != nil {
		log.Fatalf("could not enforce policy: %v", err)
	}

	fmt.Printf("authorization result: %t", ok)

}
