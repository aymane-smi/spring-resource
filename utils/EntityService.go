package utils

import (
	"errors"
	"html/template"
	"os"
	"strings"

	"github.com/aymane-smi/spring-resource/structs"
)

func GenerateEntity(entitynName string, Type int) (bool, error) {
	entity := &structs.Entity{Name: strings.Title(entitynName), TypeId: Type}
	pom := &structs.Pom{GroupId: "com.example", ArtifactId: "demo"}
	shared := &structs.Shared{SharedPom: *pom, SharedEntity: *entity}
	tmpl, _ := template.ParseFiles("static/entity.tmpl")
	if !generateTree("generated/Models") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create("generated/Models/" + entity.Name + ".java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}

func generateTree(path string) bool {
	//prototype function
	//works only on array of two ["path", "subfolder1"]
	slices := strings.Split(path, "/")
	if err := os.MkdirAll(slices[0], 0755); err != nil {
		return false
	} else if err := os.MkdirAll(slices[0]+"/"+slices[1], 0755); err != nil {
		return false
	}
	os.Chmod(slices[0]+"/"+slices[1], 0755|os.ModeSetuid)
	return true
}
