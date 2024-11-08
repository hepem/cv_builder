package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	fmt.Println("Building CV ...")
	languages := [2]string{"es", "en"}
	for _, lang := range languages {
		fmt.Printf("for %s version ...\n", lang)
		cvData, err := loadCV("data/cv_data.json", lang)
		if err != nil {
			fmt.Println("Error loading CV data: ", err)
			return
		}

		templatePath := "templates/renderCV_EngineeringResume.tex"

		texFile := fmt.Sprintf("pdfs/Hector_Perez_Munoz_CV_%s.tex", lang)

		err = render(templatePath, cvData, texFile)
		if err != nil {
			fmt.Println("Error rendering LaTeX template:", err)
		}

		err = toPDF(texFile)
		if err != nil {
			fmt.Println("Error compiling LaTeX to PDF:", err)
		}
	}

	fmt.Println("Done!")
}

func loadCV(filePath string, language string) (*CV, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var parsedData map[string]CVData
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		return nil, err
	}

	cvData, ok := parsedData[language]
	if !ok {
		return nil, fmt.Errorf("language %s not found in the data", language)
	}

	cv := CV{
		Keywords:     cvData.Keywords,
		Name:         cvData.Name,
		Birthday:     cvData.Birthday,
		Place:        cvData.Place,
		Email:        cvData.Email,
		Resume:       cvData.Resume,
		Phone:        cvData.Phone,
		Social:       cvData.Social,
		Education:    cvData.Education,
		Experience:   cvData.Experience,
		Projects:     cvData.Projects,
		Technologies: cvData.Technologies,
	}

	return &cv, nil
}

func render(templatePath string, cvData *CV, outputPath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(outputFile, cvData)
	if err != nil {
		return err
	}

	return nil
}

func toPDF(textFile string) error {
	cmd := exec.Command("pdflatex", "-output-directory=pdfs", textFile)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to compile LaTeX: %v", err)
	}
	fmt.Println("PDF successfully generated.")
	return nil
}
