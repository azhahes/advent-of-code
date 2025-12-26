package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type TemplateData struct {
	Day string
}

// must be "02", "03", "10", ...
func padDay(day string) string {
	if len(day) == 1 {
		return "0" + day
	}
	return day
}

func writeTemplate(tmplFile, outFile string, data TemplateData) {
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatalf("Template parse error: %v", err)
	}

	f, err := os.Create(outFile)
	if err != nil {
		log.Fatalf("File create error: %v", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Template execute error: %v", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run generate.go <day-number>")
	}

	day := padDay(os.Args[1])
	year := time.Now().Year()
	dayDir := filepath.Join(fmt.Sprintf("%d", year), "day"+day)
	// create folder
	err := os.MkdirAll(dayDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Cannot create dir: %v", err)
	}

	data := TemplateData{Day: day}

	// generate dayXX.go
	writeTemplate("tmpl/day.go.tmpl",
		filepath.Join(dayDir, "day"+day+".go"), data)

	// generate dayXX_test.go
	writeTemplate("tmpl/day_test.go.tmpl",
		filepath.Join(dayDir, "day"+day+"_test.go"), data)

	// create empty input.txt
	os.WriteFile(filepath.Join(dayDir, "input.txt"), []byte{}, 0o644)

	os.WriteFile(filepath.Join(dayDir, "sample-input.txt"), []byte{}, 0o644)

	fmt.Printf("Generated files for day %s 🎉\n", day)
}
