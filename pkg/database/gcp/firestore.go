package gcp

import (
	"log"
	"os"
	"text/template"

	"github.com/ThomasCardin/ci-cd/pkg/database"
)

const (
	FIRESTORE                         = "firestore"
	FIRESTORE_TERRAFORM_TEMPLATE_NAME = "firestore.tf.tmpl"
	FIRESTORE_TERRAFORM_TEMPLATE_PATH = "assets/templates/database/gcp/firestore.tf.tmpl"
)

type Firestore struct {
	ProjectID string `yaml:"projectID"`
	Location  string `yaml:"location"`
}

var _ database.Database = (*Firestore)(nil)

func (firestore *Firestore) Parse(data map[string]interface{}) {
	*firestore = Firestore{
		ProjectID: data["projectID"].(string),
		Location:  data["location"].(string),
	}
}

func (firestore *Firestore) ApplyToTerraform(outputFile *os.File) error {
	tmpl, err := template.New(FIRESTORE_TERRAFORM_TEMPLATE_NAME).ParseFiles(FIRESTORE_TERRAFORM_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = tmpl.Execute(outputFile, firestore)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Successfully applied Firestore to Terraform")
	return nil
}
