package main

import (
	"fmt"

	"github.com/aymane-smi/spring-resource/utils"
)

func main() {
	// shared := &structs.Shared{
	// 	SharedEntity: structs.Entity{
	// 		Name:     "User",
	// 		TypeId:   1,
	// 		RepoType: "JpaRepository",
	// 	},
	// 	SharedPom: structs.Pom{
	// 		GroupId:    "com.example",
	// 		ArtifactId: "demo",
	// 	},
	// }
	// service.GenerateRepository(*shared)
	// service.GenerateEntity(*shared)
	// service.GenerateService(*shared)
	// service.GenerateServiceImpl(*shared)
	fmt.Println(utils.IsJavaProject("/Users/macbookair/Desktop/YouStream/demo/pom.xml"))
	//fmt.Println(utils.GenerateProjectInfo("/Users/macbookair/Desktop/YouStream/demo/pom.xml"))
}
