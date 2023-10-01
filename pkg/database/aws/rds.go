package aws

import (
	"log"
	"os"
	"text/template"

	"github.com/ThomasCardin/ci-cd/pkg/database"
)

const (
	RDS                         = "rds"
	RDS_TERRAFORM_TEMPLATE_NAME = "rds.tf.tmpl"
	RDS_TERRAFORM_TEMPLATE_PATH = "assets/templates/database/aws/rds.tf.tmpl"
)

type Rds struct {
	DbName           string `yaml:"dbName"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	InstanceClass    string `yaml:"instanceClass"`
	AllocatedStorage int    `yaml:"allocatedStorage"`
}

var _ database.Database = (*Rds)(nil)

func (rds *Rds) Parse(data map[string]interface{}) {
	*rds = Rds{
		DbName:           data["dbName"].(string),
		Username:         data["username"].(string),
		Password:         data["password"].(string),
		InstanceClass:    data["instanceClass"].(string),
		AllocatedStorage: int(data["allocatedStorage"].(int64)),
	}
}

func (rds *Rds) ApplyToTerraform(outputFile *os.File) error {
	tmpl, err := template.New(RDS_TERRAFORM_TEMPLATE_NAME).ParseFiles(RDS_TERRAFORM_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = tmpl.Execute(outputFile, rds)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Successfully applied RDS to Terraform")
	return nil
}
