package utils

import (
	"bytes"
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
	gradle, errGradle := os.Stat(path + "/build.gradle")
	if errMaven != nil && errGradle != nil {
		return false, 0
	} else if maven != nil {
		return true, 1
	} else if gradle != nil {
		return true, 2
	}
	return false, 0
}

// generate Pom structer with given groupId and artifactId from the given path of Maven pom
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

// generate Pom structer with given groupId and artifactId from the given path of Gradle settings
func GenerateProjectInfoGradle(path string) structs.Pom {
	var gradle bytes.Buffer
	fmt.Fprintf(&gradle, "-q -p %s properties", path)
	groupCmds := []string{
		gradle.String(),
		"'^group:'",
		"'{print $2}'",
	}
	nameCmds := []string{
		gradle.String(),
		"'^name:'",
		"'{print $2}'",
	}
	outputName, errName := pipeHelper(nameCmds...)
	outputGroup, errGroup := pipeHelper(groupCmds...)
	if errName != nil || errGroup != nil {
		fmt.Println("erro ==>", errName, errGroup)
		os.Exit(1)
	}
	return structs.Pom{
		GroupId:    string(outputGroup),
		ArtifactId: string(outputName),
	}
}

// write a helper for multiple pipe commands
// assuming that len(commands) = 3
func pipeHelper(commands ...string) ([]byte, error) {
	//init
	gradle := exec.Command("gradle", commands[0])
	grep := exec.Command("grep", commands[1])
	awk := exec.Command("awk", commands[2])
	//swap rw
	grep.Stdin, _ = gradle.StdoutPipe()
	awk.Stdin, _ = grep.StdoutPipe()
	awk.Stdout = os.Stdout
	//run and wait
	_ = awk.Start()
	_ = grep.Start()
	_ = gradle.Run()
	_ = grep.Wait()
	_ = awk.Wait()
	return awk.Output()
}
