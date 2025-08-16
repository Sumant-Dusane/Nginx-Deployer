package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sumant-Dusane/nginx-deployer/docker"
	"github.com/Sumant-Dusane/nginx-deployer/dtos"
)

func GetDeployablePrograms() []dtos.ProgramDto {
	varEntries := GetVarFolder()
	optEntries := GetOptFolder()
	dockerEntries := GetDockerContainers()

	allEntries := append(varEntries, optEntries...)
	allEntries = append(allEntries, dockerEntries...)

	return allEntries
}

func GetVarFolder() []dtos.ProgramDto {
	dir := "/var/www/"
	varEntries, _ := os.ReadDir(dir)

	if len(varEntries) < 1 {
		return nil
	}

	var varEntriesArr []dtos.ProgramDto
	for _, e := range varEntries {
		varEntriesArr = append(varEntriesArr, dtos.ProgramDto{
			Id:     "-",
			Name:   e.Name(),
			Source: dir,
			Port:   "-",
		})
	}

	return varEntriesArr
}

func GetOptFolder() []dtos.ProgramDto {
	dir := "/opt/"
	optEntries, _ := os.ReadDir(dir)

	if len(optEntries) < 1 {
		return nil
	}

	var optEntriesArr []dtos.ProgramDto
	for _, e := range optEntries {
		optEntriesArr = append(optEntriesArr, dtos.ProgramDto{
			Id:     "-",
			Name:   e.Name(),
			Source: dir,
			Port:   "-",
		})
	}

	return optEntriesArr
}

func GetDockerContainers() []dtos.ProgramDto {
	containers, err := docker.GetDockerContainers()

	if err != nil {
		return nil
	}

	var dockerEntriesArr []dtos.ProgramDto
	for _, container := range containers {
		if container.Status != "running" {
			ports := []string{}
			for _, port := range container.Ports {
				ports = append(ports, fmt.Sprintf("%d:%d", port.PublicPort, port.PrivatePort))
			}
			containerId := container.ID[:12]
			containerName := strings.ReplaceAll(strings.Join(container.Names, ","), "/", "")
			containerPort := strings.Join(ports, ",")
			dockerEntriesArr = append(dockerEntriesArr, dtos.ProgramDto{
				Id:     containerId,
				Name:   containerName,
				Port:   containerPort,
				Source: "docker",
			})
		}
	}

	return dockerEntriesArr
}
