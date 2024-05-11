package service

import (
	"errors"
	"os"
	"strings"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateEntity(shared structs.Shared, path string) (bool, error) {
	tmpl, _ := template.ParseFiles("static/entity.tmpl")
	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"
	if !utils.GenerateTree(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Models") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/java/" + pomToPath + shared.SharedPom.ArtifactId + "/Models/" + shared.SharedEntity.Name + ".java")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
