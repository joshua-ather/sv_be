package migrations

import (
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
)

type Migrate struct {
	db *gorm.DB
}

func migrate(m Migrate, migrateMethodName string) {
	// Get the reflect value of the method
	x := reflect.ValueOf(m).MethodByName(migrateMethodName)
	// Exit if the method doesn't exist
	if !x.IsValid() {
		log.Fatal("No method called ", migrateMethodName)
	}
	// Execute the method
	log.Println("Migrating", migrateMethodName, "...")
	x.Call(nil)
	log.Println("Migrate", migrateMethodName, "succeed")
}

func Execute(db *gorm.DB, migrateMethodNames ...string) {
	m := Migrate{db}

	migrateType := reflect.TypeOf(m)

	// Execute all Migration if no method name is given
	if len(migrateMethodNames) == 0 {
		log.Println("Running all migration...")
		// We are looping over the method on a Migration struct
		for i := 0; i < migrateType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := migrateType.Method(i)
			// Execute Migration
			migrate(m, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range migrateMethodNames {
		migrate(m, item)
	}
}
