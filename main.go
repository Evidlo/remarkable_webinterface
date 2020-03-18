package main

import (
	// "net/http"
	// "github.com/gobuffalo/packr"
	"flag"
	"fmt"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"os"
	// "github.com/unidoc/unipdf/creator"
	"github.com/jung-kurt/gofpdf"
	"github.com/evidlo/remarkable_webinterface/boilerplate"
	"github.com/evidlo/remarkable_webinterface/rm"
)

var rgb_colors = map[rm.Rm_Color][3]int{
	rm.Rm_Color__Black: {0, 0, 0},
	rm.Rm_Color__Gray: {127, 127, 127},
	rm.Rm_Color__White: {255, 255, 255},
}

func mapgen(in_lo, in_hi, out_lo, out_hi float32) (func(input float32) float64) {
	return func(input float32) (float64) {
		return float64((input - in_lo) * (out_hi - out_lo) / (in_hi - in_lo))
	}
}

func main() {

	// ----- Parse options -----

	debug := flag.Bool("debug", false, "enable debug output")
	input := flag.String("input", "", "input .rm")
	// input_pdf := flag.String("input", "", "input .pdf")
	output := flag.String("output", "", ".pdf output path")

	flag.Parse()

	if *input == "" {
		fmt.Println("input required")
		os.Exit(1)
	}
	if *output == "" {
		fmt.Println("output required")
		os.Exit(1)
	}

	if *debug {
		fmt.Println("Debugging enabled")
	}

	// ----- Serve files -----

	// box := packr.NewBox("html/")

	// http.Handle("/", http.FileServer(box))
	// http.ListenAndServe(":3000", nil)

	// ----- Create PDF -----

	pdf := gofpdf.New("P", "mm", "letter", "")
	pdf.AddPage()
	pdf.TransformBegin()
	pdf.TransformScale(14.95, 14.95, 0, 0)
	pdf.SetLineJoinStyle("round")
	pdf.SetLineCapStyle("round")

	// increase visual difference of line widths
	mapping := mapgen(1.875, 2.125, 1, 8)
	// mapping_paint := mapgen(1.875, 2.125, 1, 8)

	f, err := os.Open(*input)
	s := kaitai.NewStream(f)
	var r rm.Rm
	err = r.Read(s, &r, &r)
	boilerplate.Check(err)

	for _, layer := range r.Layers {
		line_loop:for _, line := range layer.Lines {
			fmt.Println("color:", rgb_colors[line.Color])
			fmt.Println("brush:", line.BrushType)

			// map color to RGB array
			rgb := rgb_colors[line.Color]
			pdf.SetDrawColor(rgb[0], rgb[1], rgb[2])


			switch line.BrushType {
			case rm.Rm_Brush__Paintbrush:
				pdf.SetLineWidth(mapping(line.BrushSize))
				pdf.SetAlpha(1, "Normal")
			case rm.Rm_Brush__Mechanical:
				pdf.SetLineWidth(mapping(line.BrushSize))
				pdf.SetAlpha(1, "Normal")
			case rm.Rm_Brush__Pencil:
				pdf.SetLineWidth(mapping(line.BrushSize))
				pdf.SetAlpha(1, "Normal")
			case rm.Rm_Brush__Ballpoint:
				pdf.SetLineWidth(mapping(line.BrushSize))
				pdf.SetAlpha(1, "Normal")
			case rm.Rm_Brush__Marker:
				pdf.SetLineWidth(mapping(line.BrushSize))
				pdf.SetAlpha(1, "Normal")
			case rm.Rm_Brush__Fineliner:
				pdf.SetLineWidth(mapping(line.BrushSize))
				pdf.SetAlpha(1, "Normal")
			case rm.Rm_Brush__Highlighter:
				pdf.SetLineWidth(30)
				pdf.SetDrawColor(220, 220, 10)
				pdf.SetAlpha(0.33, "Normal")
			case rm.Rm_Brush__Eraser:
				continue line_loop
			case rm.Rm_Brush__Eraserselect:
				continue line_loop
			default:
				pdf.SetLineWidth(0)
			}

			pdf.MoveTo(float64(line.Points[0].X), float64(line.Points[0].Y))
			for _, point := range line.Points {
				if line.BrushType == rm.Rm_Brush__Ballpoint {
					fmt.Println("pressure:", point.Pressure)
				}

				pdf.LineTo(float64(point.X), float64(point.Y))
			}
			pdf.DrawPath("S")
		}
	}

	pdf.TransformEnd()

	err = pdf.OutputFileAndClose(*output)
	boilerplate.Check(err)
}
