package cmd

import (
	"fmt"

	"github.com/Sumant-Dusane/nginx-deployer/dtos"
	"github.com/Sumant-Dusane/nginx-deployer/utils"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan live applications",
	Long:  `Scan live applications on docker, scans opt, scans var/www and help to assign endpoints, ports, other config made in nginx with post actions like ssl setup using certbot, letsencrypt, etc.`,
	Run:   ScanCmd,
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func ScanCmd(cmd *cobra.Command, args []string) {
	allEntries := Scan(cmd, args)
	printTable(allEntries)
}

func Scan(cmd *cobra.Command, args []string) []dtos.ProgramDto {
	allEntries := utils.GetDeployablePrograms()
	printTable(allEntries)
	return allEntries
}

func printTable(allEntries []dtos.ProgramDto) {
	var headerNameSpan int = 60
	var rowNameSpan int = headerNameSpan - 8

	var tableHeaders = []dtos.TableData{
		{Name: "Sr. No."},
		{Name: "ID"},
		{Name: "NAME", Span: &headerNameSpan},
		{Name: "PORT"},
		{Name: "SOURCE"},
	}

	fmt.Println()
	fmt.Println()

	utils.PrintTableHeader(tableHeaders...)
	for i, e := range allEntries {
		var tableRows = []dtos.TableData{
			{Name: fmt.Sprintf("%d", i+1)},
			{Name: e.Id},
			{Name: e.Name, Span: &rowNameSpan},
			{Name: e.Port},
			{Name: e.Source},
		}

		utils.PrintTableRow(tableRows...)
	}

	fmt.Println()
	fmt.Println()
}
