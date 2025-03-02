package initialize

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func NewCasbinEnforcer(db *gorm.DB) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal("Error creating Casbin adapter:", err)
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer("rbac_model.conf", adapter)
	if err != nil {
		log.Fatal("Error creating Casbin Enforcer:", err)
		panic(err)
	}

	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatal("Error loading policy:", err)
		panic(err)
	}

	addPolicies(enforcer)

	log.Println("Casbin has been initialized")
	return enforcer
}

// Add roles and permissions to Casbin
func addPolicies(e *casbin.Enforcer) {
	// admin has full access
	e.AddPolicy("admin", "/*", "GET")
	e.AddPolicy("admin", "/*", "POST")
	e.AddPolicy("admin", "/*", "PUT")
	e.AddPolicy("admin", "/*", "PATCH")
	e.AddPolicy("admin", "/*", "DELETE")

	// student and read and write access to their own data
	e.AddNamedPolicy("p2", "student", "/students/:student_id", "GET", ":student_id")
	e.AddNamedPolicy("p2", "student", "/students/:student_id", "PUT", ":student_id")
	e.AddNamedPolicy("p2", "student", "/students/:student_id", "PATCH", ":student_id")

	e.SavePolicy()
	log.Println("Policies have been added to the database!")
}
