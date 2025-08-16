package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	allEntries := Scan(cmd, args)

	var input int
	var domain string

	fmt.Printf("Enter the Sr. No: ")
	fmt.Scanf("%d", &input)

	program := allEntries[input-1]

	fmt.Printf("Domain Name or IP Address: ")
	fmt.Scanf("%s", &domain)

	directory := program.Source + "/" + program.Name
	isDeploymentStatic := program.Source != "docker"

	DeployNginx(directory, program.Port, isDeploymentStatic, domain)
}
