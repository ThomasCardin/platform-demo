package aws

import (
	"log"
	"os"
	"text/template"

	"github.com/ThomasCardin/ci-cd/pkg/database"
)

const (
	DYNAMODB_TERRAFORM_TEMPLATE_NAME = "dynamodb.tf.tmpl"
	DYNAMODB_TERRAFORM_TEMPLATE_PATH = "assets/templates/database/aws/dynamodb.tf.tmpl"
)

type DynamoDB struct {
	TableName     string `yaml:"tableName"`
	HashKeyName   string `yaml:"hashKeyName"`
	HashKeyType   string `yaml:"hashKeyType"`
	ReadCapacity  int    `yaml:"readCapacity"`
	WriteCapacity int    `yaml:"writeCapacity"`
}

var _ database.Database = (*DynamoDB)(nil)

func (dynamodb *DynamoDB) Parse(data map[string]interface{}) {
	*dynamodb = DynamoDB{
		TableName:     data["tableName"].(string),
		HashKeyName:   data["hashKeyName"].(string),
		HashKeyType:   data["hashKeyType"].(string),
		ReadCapacity:  int(data["readCapacity"].(int)),
		WriteCapacity: int(data["writeCapacity"].(int)),
	}
}

func (dynamodb *DynamoDB) ApplyToTerraform(outputFile *os.File) error {
	tmpl, err := template.New(DYNAMODB_TERRAFORM_TEMPLATE_NAME).ParseFiles(DYNAMODB_TERRAFORM_TEMPLATE_PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = tmpl.Execute(outputFile, dynamodb)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Successfully applied DynamoDB to Terraform")
	return nil
}
