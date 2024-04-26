package service

import (
	"errors"
	"os"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateRepository(shared structs.Shared) (bool, error) {
	tmpl, _ := template.ParseFiles("static/repository.tmpl")
	if !utils.GenerateTree("generated/Repositories") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create("generated/Repositories/" + shared.SharedEntity.Name + "Repository.java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
