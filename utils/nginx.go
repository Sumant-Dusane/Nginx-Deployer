package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallNginx() {
	cmd := exec.Command("nginx", "-v")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Nginx not installed, installing... ", err)
		cmd = exec.Command("sudo", "apt-get", "update")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error updating apt: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("sudo", "apt-get", "install", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error installing nginx: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("sudo", "systemctl", "start", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error starting nginx: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("sudo", "systemctl", "enable", "nginx")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error enabling nginx: ", err)
			os.Exit(1)
		}
		cmd = exec.Command("sudo", "systemctl", "status", "nginx")
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
		fmt.Println("Error creating file: ", fileCreationErr)
		os.Exit(1)
	}

	return file
}

func AddStaticDeploymentConfig(file *os.File, domain string, directory string) {
	file.WriteString("server {")
	file.WriteString("  listen 80;")
	file.WriteString("  listen [::]:80;")
	file.WriteString("  root " + directory + ";")
	file.WriteString("  server_name " + domain + ";")
	file.WriteString("  index index.html index.htm index.nginx-debian.html;")
	file.WriteString("  location / {")
	file.WriteString("    try_files $uri $uri/ =404;")
	file.WriteString("  }")
	file.WriteString("}")
}

func AddSpaDeploymentConfig(file *os.File, domain string, directory string, port string) {
	file.WriteString("server {")
	file.WriteString("  listen 80;")
	file.WriteString("  listen [::]:80;")
	file.WriteString("  server_name " + domain + ";")
	file.WriteString("  location / {")
	file.WriteString("    proxy_pass http://localhost:" + port + ";")
	file.WriteString("    proxy_pass_header Host $host;")
	file.WriteString("  }")
	file.WriteString("}")
}
