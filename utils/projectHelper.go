package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aymane-smi/spring-resource/structs"
)

// check if the given path is for java project (maven/gradle)
// return boolean in case of success
// also return int value(0,1,2)
// 0 for non java project
// 1 for maven
// 2 for gradle
func IsJavaProject(path string) (bool, int) {
	maven, errMaven := os.Stat(path + "/pom.xml")
	gradle, errGradle := os.Stat(path)
	if errMaven != nil || errGradle != nil {
		return false, 0
	} else if maven != nil {
		return true, 1
	} else if gradle != nil {
		return true, 2
	}
	return false, 0
}

// generate Pom structer with given groupId and artifactId from the given path
func GenerateProjectInfoMaven(path string) structs.Pom {
	argsGroup := []string{"help:evaluate", "-Dexpression=project.groupId", "-q", "-f", path, "-DforceStdout"}
	argsArtifact := []string{"help:evaluate", "-Dexpression=project.artifactId", "-q", "-f", path, "-DforceStdout"}
	outputGroup, errGroup := exec.Command("mvn", argsGroup...).Output()
	outputArtifact, errArtifact := exec.Command("mvn", argsArtifact...).Output()
	if errGroup != nil || errArtifact != nil {
		fmt.Println(errArtifact, errGroup)
	}
	return structs.Pom{GroupId: string(outputGroup), ArtifactId: string(outputArtifact)}
}
