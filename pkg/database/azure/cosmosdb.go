package azure

import (
	"log"
	"os"
	"text/template"

	"github.com/ThomasCardin/ci-cd/pkg/database"
)

const (
	COSMOSDB                         = "cosmosdb"
	COSMOSDB_TERRAFORM_TEMPLATE_NAME = "cosmosdb.tf.tmpl"
	COSMOSDB_TERRAFORM_TEMPLATE_PATH = "assets/templates/database/azure/cosmosdb.tf.tmpl"
)

type CosmosDB struct {
	Name              string `yaml:"name"`
	ResourceGroupName string `yaml:"resourceGroupName"`
	Location          string `yaml:"location"`
}

var _ database.Database = (*CosmosDB)(nil)

func (cosmosdb *CosmosDB) Parse(data map[string]interface{}) {
	*cosmosdb = CosmosDB{
		Name:              data["name"].(string),
		ResourceGroupName: data["resourceGroupName"].(string),
		Location:          data["location"].(string),
	}
}

func (cosmosdb *CosmosDB) ApplyToTerraform(outputFile *os.File) error {
	tmpl, err := template.New(COSMOSDB_TERRAFORM_TEMPLATE_NAME).ParseFiles(COSMOSDB_TERRAFORM_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = tmpl.Execute(outputFile, cosmosdb)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Successfully applied CosmosDB to Terraform")
	return nil
}
