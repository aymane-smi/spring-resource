package service

import (
	"errors"
	"os"
	"strings"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func GenerateEntity(shared structs.Shared, path string, templType int) (bool, error) {
	var tmpl *template.Template
	var folderName, extension string

	if templType == 1 {
		tmpl, _ = template.ParseFiles("static/entity.tmpl")
		folderName = "java"
		extension = ".java"
	} else {
		tmpl, _ = template.ParseFiles("static/kotlin/entity.tmpl")
		folderName = "kotlin"
		extension = ".kt"
	}

	pomToPath := strings.ReplaceAll(shared.SharedPom.GroupId, ".", "/") + "/"
	if !utils.GenerateTree(path + "/src/main/" + folderName + "/" + pomToPath + shared.SharedPom.ArtifactId + "/Models") {
		return false, errors.New("can't create folder or subfolder")
	}
	file, errFile := os.Create(path + "/src/main/" + folderName + "/" + pomToPath + shared.SharedPom.ArtifactId + "/Models/" + shared.SharedEntity.Name + extension)
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, shared)
	return true, nil
}
