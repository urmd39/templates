package service

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"templates/infrastructure"

	"github.com/tealeg/xlsx"
)

var (
	Excel_VNLocationOptions     = xlsx.DateTimeOptions{Location: infrastructure.VNLocation, ExcelTimeFormat: "dd-MM-yyyy HH:mm"}
	Excel_DateVNLocationOptions = xlsx.DateTimeOptions{Location: infrastructure.VNLocation, ExcelTimeFormat: "dd-MM-yyyy"}
	DoSpaceContentType          = "application/vnd.ms-excel"
)

func addExcelTitle(sheet *xlsx.Sheet, tableTitle []string, style *xlsx.Style, fontSize int) {
	row := sheet.AddRow()
	for i := 0; i < len(tableTitle); i++ {
		cell := row.AddCell()
		cell.Value = tableTitle[i]
		if style != nil {
			cell.SetStyle(style)
		} else {
			cell.SetStyle(&xlsx.Style{
				Border:         xlsx.Border{Bottom: "thin", Left: "thin", Right: "thin", Top: "thin"},
				Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "center", Vertical: "center", ShrinkToFit: true},
				Font:           xlsx.Font{Bold: true, Name: "Times New Roman", Size: fontSize},
				ApplyFont:      true,
				ApplyAlignment: true,
				ApplyBorder:    true,
			})
		}
	}
}

func setCellCenterStyleNoneBorder(cell *xlsx.Cell, bold bool, fontSize int) {
	cell.SetStyle(&xlsx.Style{
		Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "center", Vertical: "center", ShrinkToFit: true},
		Font:           xlsx.Font{Name: "Times New Roman", Size: fontSize, Bold: bold},
		ApplyFont:      true,
		ApplyAlignment: true,
	})
}

func setCellCenterStyleBorder(cell *xlsx.Cell, bold bool, fontSize int) {
	cell.SetStyle(&xlsx.Style{
		Border:         xlsx.Border{Bottom: "thin", Left: "thin", Right: "thin", Top: "thin"},
		Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "center", Vertical: "center", ShrinkToFit: true},
		Font:           xlsx.Font{Name: "Times New Roman", Size: fontSize, Bold: bold},
		ApplyBorder:    true,
		ApplyFont:      true,
		ApplyAlignment: true,
	})
}

func setCellRightStyleNoneBorder(cell *xlsx.Cell, bold bool, fontSize int) {
	cell.SetStyle(&xlsx.Style{
		Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "right", Vertical: "center", ShrinkToFit: true},
		Font:           xlsx.Font{Name: "Times New Roman", Size: fontSize, Bold: bold},
		ApplyFont:      true,
		ApplyAlignment: true,
	})
}

func setCellRightStyleBorder(cell *xlsx.Cell, bold bool, fontSize int) {
	cell.SetStyle(&xlsx.Style{
		Border:         xlsx.Border{Bottom: "thin", Left: "thin", Right: "thin", Top: "thin"},
		Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "right", Vertical: "center", ShrinkToFit: true},
		Font:           xlsx.Font{Name: "Times New Roman", Size: fontSize, Bold: bold},
		ApplyBorder:    true,
		ApplyFont:      true,
		ApplyAlignment: true,
	})
}

func setCellLeftStyleNoneBorder(cell *xlsx.Cell, bold bool, fontSize int) {
	cell.SetStyle(&xlsx.Style{
		Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "left", Vertical: "center", ShrinkToFit: true},
		Font:           xlsx.Font{Name: "Times New Roman", Size: fontSize, Bold: bold},
		ApplyFont:      true,
		ApplyAlignment: true,
	})
}

func setCellLeftStyleBorder(cell *xlsx.Cell, bold bool, fontSize int) {
	cell.SetStyle(&xlsx.Style{
		Border:         xlsx.Border{Bottom: "thin", Left: "thin", Right: "thin", Top: "thin"},
		Alignment:      xlsx.Alignment{WrapText: true, Horizontal: "left", Vertical: "center", ShrinkToFit: true},
		Font:           xlsx.Font{Name: "Times New Roman", Size: fontSize, Bold: bold},
		ApplyBorder:    true,
		ApplyFont:      true,
		ApplyAlignment: true,
	})
}

func getImageFromUrl(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getExcelURL(fileName string, domainReport, defaultUrl string) string {
	domainUrl, err := url.Parse(domainReport)
	if err != nil || strings.TrimSpace(fileName) == "" {
		infrastructure.ErrLog.Printf("error image domain %s, error %v\n", domainReport, err)
		return defaultUrl
	}
	domainUrl.Path = path.Join(domainUrl.Path, fileName)
	return domainUrl.String()
}
