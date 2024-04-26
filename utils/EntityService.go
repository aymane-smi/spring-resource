package utils

import (
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
	os.Mkdir("generated/Models", 0775)
	file, errFile := os.Create(entity.Name + ".java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
