package service

import (
	"errors"
	"os"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateException(shared structs.Shared) (bool, error) {
	tmpl, _ := template.ParseFiles("static/exception.tmpl")
	if !utils.GenerateTree("generated/exceptions") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create("generated/exceptions/" + shared.SharedException.Name + ".java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
