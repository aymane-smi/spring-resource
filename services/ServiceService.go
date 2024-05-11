package service

import (
	"errors"
	"os"
	"strings"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateService(shared structs.Shared, path string) (bool, error) {
	tmpl, _ := template.ParseFiles("static/service.tmpl")
	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"
	if !utils.GenerateTree(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Services") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Services/" + shared.SharedEntity.Name + "Service.java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}

func GenerateServiceImpl(shared structs.Shared, path string) (bool, error) {
	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"
	tmpl, _ := template.ParseFiles("static/serviceImpl.tmpl")
	if !utils.GenerateTree(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Services/Impl") {
		println("can't create subfolder in service impl")
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Services/Impl/" + shared.SharedEntity.Name + "ServiceImpl.java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
