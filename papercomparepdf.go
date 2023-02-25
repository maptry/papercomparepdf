package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/phpdave11/gofpdf"
)

type PaperSize struct {
	Source string
	Name   string
	Width  int // mm
	Height int // mm
}

const fontSize = 14

func (ps PaperSize) SmallerThan(ps2 PaperSize) bool {
	// returns true if and only if the width and the height is less
	return ps.Width < ps2.Width && ps.Height < ps2.Height
}

/*
	size explanation: https://unsharpen.com/paper-sizes/
*/

var PaperSizes = map[string]PaperSize{
	"A5":      PaperSize{"ISO", "A5", 148, 210},
	"A6":      PaperSize{"ISO", "A6", 105, 148},
	"A7":      PaperSize{"ISO", "A7", 74, 105},
	"A8":      PaperSize{"ISO", "A8", 52, 74},
	"A9":      PaperSize{"ISO", "A9", 37, 52},
	"A10":     PaperSize{"ISO", "A10", 26, 37},
	"B6":      PaperSize{"ISO", "B6", 125, 176},
	"B7":      PaperSize{"ISO", "B7", 88, 125},
	"B8":      PaperSize{"ISO", "B8", 62, 88},
	"B9":      PaperSize{"ISO", "B9", 44, 62},
	"B10":     PaperSize{"ISO", "B10", 31, 44},
	"Invoice": PaperSize{"US", "Invoice", 140, 216},
	//"Executive":   PaperSize{"US", "Executive", 184, 267},
	"Field Notes":           PaperSize{"Other", "Field Notes", 90, 140},
	"Moleskine Extra Small": PaperSize{"Other", "Moleskine Extra Small", 65, 105},
	"Moleskine Pocket":      PaperSize{"Other", "Moleskine Pocket", 90, 140},
}
var PaperSizesPrinter = map[string]PaperSize{
	"A4":     PaperSize{"ISO", "A4", 210, 297},
	"Letter": PaperSize{"US", "Letter", 216, 279},
	"Legal":  PaperSize{"US", "Legal", 203, 330},
}

func joinKeys(m map[string]PaperSize) string {
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return strings.Join(keys, "|")
}

func usage() {
	fmt.Printf("Usage: %s [-p %s] PaperSize...\n", os.Args[0], joinKeys(PaperSizesPrinter))
	fmt.Printf("PaperSize can be one of %v\n", joinKeys(PaperSizes))
}

func drawVerticalRuler(pdf *gofpdf.Fpdf, s string) {
	const fontSize = 10
	pdf.SetFont("Times", "", fontSize)
	pdf.SetLineWidth(0.1)
	pdf.MoveTo(30, 0)
	pdf.LineTo(30, float64(PaperSizesPrinter[s].Height))
	pdf.DrawPath("D")
	for i := 0; i <= PaperSizesPrinter[s].Height; i = i + 10 {
		pdf.MoveTo(27, float64(i))
		pdf.LineTo(30, float64(i))
		pdf.DrawPath("D")
		pdf.MoveTo(26, float64(i))
		pdf.CellFormat(1, 5, strconv.Itoa(i/10), "", 0, "RM", false, 0, "")
	}
	for i := 0; i <= PaperSizesPrinter[s].Height; i = i + 10 {
		y := PaperSizesPrinter[s].Height - i
		pdf.MoveTo(30, float64(y))
		pdf.LineTo(33, float64(y))
		pdf.DrawPath("D")
		pdf.MoveTo(34, float64(y))
		pdf.CellFormat(1, 5, strconv.Itoa(i/10), "", 0, "LT", false, 0, "")
	}
	pdf.SetLineWidth(0.05)
	for i := 0; i <= PaperSizesPrinter[s].Height; i = i + 1 {
		pdf.MoveTo(29, float64(i))
		pdf.LineTo(30, float64(i))
		pdf.DrawPath("D")
	}
	for i := PaperSizesPrinter[s].Height; i >= 0; i = i - 1 {
		pdf.MoveTo(30, float64(i))
		pdf.LineTo(31, float64(i))
		pdf.DrawPath("D")
	}

}

func drawHorizontalRuler(pdf *gofpdf.Fpdf, s string) {
	pdf.SetLineWidth(0.1)
	pdf.MoveTo(0, 30)
	pdf.LineTo(float64(PaperSizesPrinter[s].Width), 30)
	pdf.DrawPath("D")
	for i := 0; i <= PaperSizesPrinter[s].Width; i = i + 10 {
		pdf.MoveTo(float64(i), 27)
		pdf.LineTo(float64(i), 30)
		pdf.DrawPath("D")
		pdf.MoveTo(float64(i), 21)
		pdf.CellFormat(1, 5, strconv.Itoa(i/10), "", 0, "LB", false, 0, "")
	}
	for i := PaperSizesPrinter[s].Width; i >= 0; i = i - 10 {
		x := PaperSizesPrinter[s].Width - i
		pdf.MoveTo(float64(x), 30)
		pdf.LineTo(float64(x), 33)
		pdf.DrawPath("D")
		pdf.MoveTo(float64(x), 34)
		pdf.CellFormat(1, 5, strconv.Itoa(i/10), "", 0, "LT", false, 0, "")
	}
	pdf.SetLineWidth(0.05)
	for i := 0; i <= PaperSizesPrinter[s].Width; i = i + 1 {
		pdf.MoveTo(float64(i), 29)
		pdf.LineTo(float64(i), 30)
		pdf.DrawPath("D")
	}
	for i := PaperSizesPrinter[s].Width; i >= 0; i = i - 1 {
		pdf.MoveTo(float64(i), 30)
		pdf.LineTo(float64(i), 31)
		pdf.DrawPath("D")
	}
}

func drawPaperSize(pdf *gofpdf.Fpdf, s string) {
	const originX = 50
	const originY = 50
	pdf.SetLineWidth(0.3)
	pdf.MoveTo(originX, originY)
	pdf.LineTo(originX+float64(PaperSizes[s].Width), originY)
	pdf.MoveTo(originX+float64(PaperSizes[s].Width), originY)
	pdf.LineTo(originX+float64(PaperSizes[s].Width), originY+float64(PaperSizes[s].Height))
	pdf.MoveTo(originX+float64(PaperSizes[s].Width), originY+float64(PaperSizes[s].Height))
	pdf.LineTo(originX, originY+float64(PaperSizes[s].Height))
	pdf.MoveTo(originX, originY+float64(PaperSizes[s].Height))
	pdf.LineTo(originX, originY)
	pdf.DrawPath("D")
	pdf.MoveTo(originX+float64(PaperSizes[s].Width), originY+float64(PaperSizes[s].Height)+0)
	pdf.SetFont("Times", "", 12)
	output := fmt.Sprintf("%s (%d x %d)", s, PaperSizes[s].Width, PaperSizes[s].Height)
	pdf.CellFormat(1, 5, output, "", 0, "RM", false, 0, "")
}

func drawPrintNotice(pdf *gofpdf.Fpdf, PaperPrinter string) {
	pdf.MoveTo(50, 40)
	pdf.SetFont("Times", "", 14)
	output := fmt.Sprintf("Please print this page without any scaling on %s paper! (Units are mm)", PaperPrinter)
	pdf.CellFormat(1, 5, output, "", 0, "LT", false, 0, "")
}

func main() {
	var PaperPrinter string
	flag.StringVar(&PaperPrinter, "p", "A4", "Specify the paper size of your printer. Default is A4")
	flag.Usage = usage
	flag.Parse()
	if _, ok := PaperSizesPrinter[PaperPrinter]; !ok {
		fmt.Printf("paper size \"%s\" choosen for printer is unknown/not allowed\n", PaperPrinter)
		os.Exit(1)
	}
	sizes := flag.Args()
	for _, s := range sizes {
		if _, ok := PaperSizes[s]; !ok {
			fmt.Printf("paper size \"%s\" is unknown/not allowed", s)
			os.Exit(1)
		}
	}
	// Initialize the graphic context on a pdf document
	pdf := gofpdf.New("P", "mm", PaperPrinter, "")
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()
	drawPrintNotice(pdf, PaperPrinter)
	drawVerticalRuler(pdf, PaperPrinter)
	drawHorizontalRuler(pdf, PaperPrinter)
	for _, s := range sizes {
		drawPaperSize(pdf, s)
	}
	pdf.OutputFileAndClose("ouput.pdf")
}
