package main

import (
	"fmt"
	"os"

	"github.com/aymane-smi/spring-resource/cmd"
	service "github.com/aymane-smi/spring-resource/services"
	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
)

func main() {
	var path string
	shared := &structs.Shared{
		SharedEntity: structs.Entity{
			Name:     "",
			TypeId:   1,
			RepoType: "JpaRepository",
		},
		SharedPom: structs.Pom{
			GroupId:    "com.example",
			ArtifactId: "demo",
		},
	}
	//fmt.Println(utils.IsJavaProject("/Users/macbookair/Desktop/YouStream/demo/pom.xml"))
	//fmt.Println(utils.GenerateProjectInfoMaven("/Users/macbookair/Desktop/YouStream/demo/pom.xml"))
	entityType, errEntity := cmd.Execute(1)
	repositoryType, errRepository := cmd.Execute(2)
	shared.SharedEntity.TypeId = entityType.(int)
	shared.SharedEntity.RepoType = repositoryType.(string)
	if errRepository != nil || errEntity != nil {
		fmt.Println("error during list generation")
	}
	fmt.Printf("project path:")
	fmt.Scanf("%s", path)
	isJava, _ := utils.IsJavaProject(path + "/pom.xml")
	if !isJava {
		fmt.Println("the path is not for java project")
		os.Exit(1)
	}
	shared.SharedPom = utils.GenerateProjectInfoMaven(path + "/pom.xml")
	service.GenerateRepository(*shared, path)
	service.GenerateEntity(*shared, path)
	service.GenerateService(*shared)
	service.GenerateServiceImpl(*shared)
}
