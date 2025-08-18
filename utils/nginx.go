package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallNginx() {
	cmd := exec.Command("systemctl", "status", "nginx")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Nginx not installed, installing... ", err)
		cmd = exec.Command("apt-get", "update")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error updating apt: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("apt-get", "install", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error installing nginx: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("systemctl", "start", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error starting nginx: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("systemctl", "enable", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error enabling nginx: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("systemctl", "status", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error checking nginx status: ", err)
			os.Exit(1)
		}
		fmt.Println("Nginx installed successfully")
	}

}

func CreateNginxConfig(fileName *string) *os.File {
	var file *os.File
	var fileCreationErr error
	if fileName == nil {
		file, fileCreationErr = os.Create("/etc/nginx/sites-available/default")
	} else {
		file, fileCreationErr = os.Create("/etc/nginx/sites-available/" + *fileName + ".conf")
	}

	if fileCreationErr != nil {
		fmt.Println("Error configuring nginx: ", fileCreationErr)
		os.Exit(1)
	}

	return file
}

func AddStaticDeploymentConfig(file *os.File, domain string, directory string) {
	file.WriteString("server {\n")
	file.WriteString("  listen 80;\n")
	file.WriteString("  listen [::]:80;\n")
	file.WriteString("  root " + directory + ";\n")
	file.WriteString("  server_name " + domain + ";\n")
	file.WriteString("  index index.html index.htm index.nginx-debian.html;\n")
	file.WriteString("  location / {\n")
	file.WriteString("    try_files $uri $uri/ =404;\n")
	file.WriteString("  }\n")
	file.WriteString("}\n")
}

func AddSpaDeploymentConfig(file *os.File, domain string, directory string, port string) {
	file.WriteString("server {\n")
	file.WriteString("  listen 80;\n")
	file.WriteString("  listen [::]:80;\n")
	file.WriteString("  server_name " + domain + ";\n")
	file.WriteString("  location / {\n")
	file.WriteString("    proxy_pass http://localhost:" + port + ";\n")
	file.WriteString("    proxy_pass_header Host $host;\n")
	file.WriteString("  }\n")
	file.WriteString("}\n")
}
