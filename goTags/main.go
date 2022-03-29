package main

import (
	"fmt"
	"os"
	"reflect"
)

type Config struct {
	Email    string `env:"EMAIL"`
	Password string `env:"PASSWORD"`
	Port     string `env:"PORT"`
}

const tagName = "env"

func LoadConfig(q *Config) {
	v := reflect.ValueOf(q).Elem()
	if v.Kind() == reflect.Struct {
		val := reflect.TypeOf(q).Elem()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			tag := field.Tag.Get(tagName)
			fmt.Printf("field : %v | tagName : %v\n", field.Name, tag)
			envVal := os.Getenv(tag)
			reflect.ValueOf(q).Elem().FieldByName(field.Name).Set(reflect.ValueOf(envVal))
		}
	}
}

func main() {
	var cfg Config

	LoadConfig(&cfg)
	fmt.Println(cfg)
}
