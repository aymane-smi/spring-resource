package service

import (
	"errors"
	"os"
	"strings"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateRepository(shared structs.Shared, path string) (bool, error) {
	tmpl, _ := template.ParseFiles("static/repository.tmpl")
	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"
	if !utils.GenerateTree(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Repositories") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Repositories/" + shared.SharedEntity.Name + "Repository.java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
