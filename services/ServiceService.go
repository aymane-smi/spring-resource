package service

import (
	"errors"
	"os"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateService(shared structs.Shared) (bool, error) {
	tmpl, _ := template.ParseFiles("static/service.tmpl")
	if !utils.GenerateTree("generated/Services") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create("generated/Services/" + shared.SharedEntity.Name + "Service.java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}

func GenerateServiceImpl(shared structs.Shared) (bool, error) {
	tmpl, _ := template.ParseFiles("static/serviceImpl.tmpl")
	if !utils.GenerateTree("generated/Services/Impl") {
		println("can't create subfolder in service impl")
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create("generated/Services/Impl/" + shared.SharedEntity.Name + "ServiceImpl.java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
