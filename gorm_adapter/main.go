package main

import (
	"fmt"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type Resource struct {
	Id     string
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

	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", user, pass)

	a, err := gormadapter.NewAdapter("mysql", dsn)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	e, err := casbin.NewEnforcer("model.conf", a)

	if err != nil {
		log.Fatalf("could not create enforcer: %v", err)
	}

	e.LoadPolicy()
	e.AddPolicy("admin", "data1", "read|write")

	data1 := Resource{
		Id:     "data1",
		Owners: []string{"sekthor", "testuser"},
	}

	ok, err := e.Enforce("admin", data1, "read|write")

	if err != nil {
		log.Fatalf("could not enforce policy: %v", err)
	}

	fmt.Printf("authorization result: %t", ok)
}
