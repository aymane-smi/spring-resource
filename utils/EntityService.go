package utils

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	//"github.com/aymane-smi/spring-resource/structs"
)

type Entity struct {
	Name   string
	TypeId int
}

func GenerateEntity(entitynName string, Type int) (bool, error) {
	entity := &Entity{Name: strings.Title(entitynName), TypeId: Type}
	tmpl, _ := template.ParseFiles("static/entity.tmpl")
	os.Mkdir("generated", os.ModeAppend)
	file, errFile := os.Create(entity.Name + ".java")
	if errFile != nil {
		fmt.Errorf("can't work create a new file")
	}
	defer file.Close()
	tmpl.Execute(file, entity)
	return true, nil
}
