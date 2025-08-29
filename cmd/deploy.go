package cmd

import (
	"fmt"
	"os"

	"github.com/Sumant-Dusane/nginx-deployer/utils"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Create and deploy nginx config",
	Long:  `Create and deploy nginx config for a given domain and source`,
	Run:   Deploy,
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("directory", "d", "", "Directory of the deployment")
	deployCmd.Flags().StringP("port", "p", "", "Port of the deployment")
	deployCmd.Flags().StringP("domain", "a", "", "Domain name or IP address of the deployment")
	deployCmd.Flags().BoolP("static", "i", false, "Is the deployment static")
}

func Deploy(cmd *cobra.Command, args []string) {
	directory, _ := cmd.Flags().GetString("directory")
	port, _ := cmd.Flags().GetString("port")
	domain, _ := cmd.Flags().GetString("domain")
	isDeploymentStatic, _ := cmd.Flags().GetBool("static")

	if directory == "" || port == "" || domain == "" {
		fmt.Println("Kindly provide all the required arguments. Use --help to see the available options.")
		os.Exit(1)
	}

	DeployNginx(directory, port, isDeploymentStatic, domain)
}

func DeployNginx(directory string, port string, isDeploymentStatic bool, domain string) {
	var confFileName *string
	fmt.Printf("Enter conf file name (default is default): ")
	fmt.Scanf("%s", confFileName)

	utils.InstallNginx()
	file := utils.CreateNginxConfig(confFileName)

	if isDeploymentStatic {
		utils.AddStaticDeploymentConfig(file, domain, directory)
	} else {
		utils.AddSpaDeploymentConfig(file, domain, directory, port)
	}

	fmt.Println("Successfully deployed on http://" + domain)
	fmt.Print("Would you like to deploy on https://" + domain + "? (y/n): ")
	var deployOnHttps string
	fmt.Scanf("%s", &deployOnHttps)

	if deployOnHttps == "y" {
		utils.InstallCertbot()
		utils.RunCertbotForHttps(domain)
	}
}
