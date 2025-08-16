package utils

import (
	"fmt"

	"github.com/Sumant-Dusane/nginx-deployer/dtos"
)

func FormatBoldText(text string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", text)
}

func PrintTableHeader(headers ...dtos.TableData) {
	for i, header := range headers {
		if i > 0 {
			fmt.Printf(" | ")
		}
		var span int
		if header.Span != nil {
			span = *header.Span
		} else {
			span = -20
		}
		fmt.Printf("%-*s", span, FormatBoldText(header.Name))
	}
	fmt.Println()
}

func PrintTableRow(rows ...dtos.TableData) {
	for i, row := range rows {
		if i > 0 {
			fmt.Printf(" | ")
		}
		var span int
		if row.Span != nil {
			span = *row.Span
		} else {
			span = -12
		}
		fmt.Printf("%-*s", span, row.Name)
	}
	fmt.Println()
}
