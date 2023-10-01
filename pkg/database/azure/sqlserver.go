package azure

import (
	"log"
	"os"
	"text/template"

	"github.com/ThomasCardin/ci-cd/pkg/database"
)

const (
	SQLSERVER_TERRAFORM_TEMPLATE_NAME = "sqlserver.tf.tmpl"
	SQLSERVER_TERRAFORM_TEMPLATE_PATH = "assets/templates/database/azure/sqlserver.tf.tmpl"
)

type SQLServer struct {
	Name              string `yaml:"name"`
	ResourceGroupName string `yaml:"resourceGroupName"`
	Location          string `yaml:"location"`
	AdminLogin        string `yaml:"adminLogin"`
	AdminPassword     string `yaml:"adminPassword"`
}

var _ database.Database = (*SQLServer)(nil)

func (sqlserver *SQLServer) Parse(data map[string]interface{}) {
	*sqlserver = SQLServer{
		Name:              data["name"].(string),
		ResourceGroupName: data["resourceGroupName"].(string),
		Location:          data["location"].(string),
		AdminLogin:        data["adminLogin"].(string),
		AdminPassword:     data["adminPassword"].(string),
	}
}

func (sqlserver *SQLServer) ApplyToTerraform(outputFile *os.File) error {
	tmpl, err := template.New(SQLSERVER_TERRAFORM_TEMPLATE_NAME).ParseFiles(SQLSERVER_TERRAFORM_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = tmpl.Execute(outputFile, sqlserver)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Successfully applied SQLServer to Terraform")
	return nil
}
