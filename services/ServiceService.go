package service

import (
	"errors"
	"os"
	"strings"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateService(shared structs.Shared, path string, templType int) (bool, error) {
	var tmpl *template.Template
	var folderName, extension string

	if templType == 1 {
		tmpl, _ = template.ParseFiles("static/service.tmpl")
		folderName = "java"
		extension = ".java"
	} else {
		tmpl, _ = template.ParseFiles("static/kotlin/service.tmpl")
		folderName = "kotlin"
		extension = ".kt"
	}
	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"

	if !utils.GenerateTree(path + "/src/main/" + folderName + "/" + pomToPath + shared.SharedPom.ArtifactId + "/Services") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/" + folderName + "/" + pomToPath + shared.SharedPom.ArtifactId + "/Services/" + shared.SharedEntity.Name + "Service" + extension)
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}

func GenerateServiceImpl(shared structs.Shared, path string, templType int) (bool, error) {
	var tmpl *template.Template
	var folderName, extension string

	if templType == 1 {
		tmpl, _ = template.New("service").Funcs(template.FuncMap{
			"UnCapitalized": utils.Uncapitalized,
		}).ParseFiles("static/serviceImpl.tmpl")
		folderName = "java"
		extension = ".java"
	} else {
		tmpl, _ = template.ParseFiles("static/kotlin/serviceImpl.tmpl")
		folderName = "kotlin"
		extension = ".kt"
	}
	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"
	if !utils.GenerateTree(path + "/src/main/" + folderName + "/" + pomToPath + shared.SharedPom.ArtifactId + "/Services/Impl") {
		println("can't create subfolder in service impl")
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/" + folderName + "/" + pomToPath + shared.SharedPom.ArtifactId + "/Services/Impl/" + shared.SharedEntity.Name + "ServiceImpl" + extension)
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
