package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Params struct {
	Namespace string `yaml:"namespace"`
}

func main() {
	data, err := ioutil.ReadFile(filepath.Join("ci", "infra.yml"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	params := &Params{}
	err = yaml.Unmarshal(data, params)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	projectName := params.ProjectName
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		err := os.MkdirAll(projectName, 0755)
		if err != nil {
			log.Fatalf("Erreur lors de la création du répertoire : %v", err)
		}
		log.Printf("Répertoire %v créé avec succès.", projectName)
	} else {
		log.Printf("Le répertoire %v existe déjà.", projectName)
	}

	templatePath := filepath.Join("templates", "aws", "dynamodb.tf.tmpl")
	tmpl, err := template.New(templatePath).ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file, err := os.Create(filepath.Join(projectName, "dynamodb.tf"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, params)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
