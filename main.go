package main

import (
	"fmt"
	"os"

	"github.com/aymane-smi/spring-resource/cmd"
	service "github.com/aymane-smi/spring-resource/services"
	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	newFigure := figure.NewFigure("Spring Resource", "", true)
	newFigure.Print()
	var path string
	shared := &structs.Shared{}
	entityType, errEntity := cmd.Execute(1)
	repositoryType, errRepository := cmd.Execute(2)
	shared.SharedEntity.TypeId = entityType.(int)
	shared.SharedEntity.RepoType = repositoryType.(string)
	if errRepository != nil || errEntity != nil {
		fmt.Println("error during list generation")
	}
	fmt.Print("Entity name:")
	fmt.Scanf("%s", &shared.SharedEntity.Name)
	fmt.Printf("project path:")
	fmt.Scanf("%s", &path)
	isJava, typeOfManager := utils.IsJavaProject(path)
	if !isJava {
		fmt.Println("the path is not for java project")
		os.Exit(1)
	}
	if typeOfManager == 1 {

	} else {
		shared.SharedPom = utils.GenerateProjectInfoGradle(path)
	}
	service.GenerateRepository(*shared, path)
	service.GenerateEntity(*shared, path)
	service.GenerateService(*shared, path)
	service.GenerateServiceImpl(*shared, path)
	fmt.Println("\nall file generated✅ please check the folder =>", path)
}
