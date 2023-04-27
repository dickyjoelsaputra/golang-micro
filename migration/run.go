package main

import (
	"fmt"
	"golang-micro/configs"
	"golang-micro/entity"
	"golang-micro/utils"
	"reflect"
	"time"
)

func main() {
	Migrate(
		&entity.Laptop{},
	)
}

// Migrate ..
func Migrate(data ...interface{}) {
	if utils.Getenv("DB_POSTGRES_MIGRATE", "true") == "true" && utils.Getenv("APP_ENV", "dev") != "production" {
		fmt.Println("Migrate DB")
		configs.Dbcon.AutoMigrate(data...)
		for _, v := range data {
			fmt.Printf("Migrate Table %s at %s\n", getType(v), time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
