package gcp

import (
	"log"
	"os"
	"text/template"

	"github.com/ThomasCardin/ci-cd/pkg/database"
)

const (
	BIGTABLE                         = "bigtable"
	BIGTABLE_TERRAFORM_TEMPLATE_NAME = "bigtable.tf.tmpl"
	BIGTABLE_TERRAFORM_TEMPLATE_PATH = "assets/templates/database/gcp/bigtable.tf.tmpl"
)

type Bigtable struct {
	Name        string `yaml:"name"`
	ClusterID   string `yaml:"clusterID"`
	Zone        string `yaml:"zone"`
	NumNodes    int    `yaml:"numNodes"`
	DisplayName string `yaml:"displayName"`
}

var _ database.Database = (*Bigtable)(nil)

func (bigtable *Bigtable) Parse(data map[string]interface{}) {
	*bigtable = Bigtable{
		Name:        data["name"].(string),
		ClusterID:   data["clusterID"].(string),
		Zone:        data["zone"].(string),
		NumNodes:    int(data["numNodes"].(int)),
		DisplayName: data["displayName"].(string),
	}
}

func (bigtable *Bigtable) ApplyToTerraform(outputFile *os.File) error {
	tmpl, err := template.New(BIGTABLE_TERRAFORM_TEMPLATE_NAME).ParseFiles(BIGTABLE_TERRAFORM_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = tmpl.Execute(outputFile, bigtable)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Successfully applied Bigtable to Terraform")
	return nil
}
