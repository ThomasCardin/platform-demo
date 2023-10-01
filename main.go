package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ThomasCardin/ci-cd/pkg/database/aws"
	"github.com/ThomasCardin/ci-cd/pkg/database/azure"
	"github.com/ThomasCardin/ci-cd/pkg/database/gcp"
	"gopkg.in/yaml.v3"
)

const (
	PATH = "test/ci/infra.yml"
)

type Params struct {
	Namespace string                   `yaml:"namespace"`
	Databases map[string][]interface{} `yaml:"databases"`
}

/*
TODO: Add command line arguments
-> For testing
-> Trigger the creation of databases only, etc..
*/
func main() {
	data, err := os.ReadFile(PATH)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	params := &Params{}
	err = yaml.Unmarshal(data, params)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Create output directory and subdirectories + file.tf
	namespace := filepath.Join("test", params.Namespace)
	if _, err := os.Stat(namespace); os.IsNotExist(err) {
		err := os.MkdirAll(namespace, 0755)
		if err != nil {
			log.Fatalf("Error creating directory : %s", err)
		}
		log.Printf("Directory created: %s", namespace)
	} else {
		log.Printf("Directory already exist: %s", namespace)
	}

	// This creates (or opens if exists) a single file to append all templates
	filePath := filepath.Join(namespace, "infra-"+params.Namespace+".tf")
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			log.Fatalf("Error deleting file : %v", err)
		}
	}

	outputFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer outputFile.Close()

	for dbType, dbDetails := range params.Databases {
		switch dbType {
		case aws.DYNAMODB:
			for _, instance := range dbDetails {
				var dynamodb aws.DynamoDB
				dynamodb.Parse(instance.(map[string]interface{}))
				dynamodb.ApplyToTerraform(outputFile)
			}
		case aws.RDS:
			for _, instance := range dbDetails {
				var rds aws.Rds
				rds.Parse(instance.(map[string]interface{}))
				rds.ApplyToTerraform(outputFile)
			}
		case azure.SQLSERVER:
			for _, instance := range dbDetails {
				var sqlserver azure.SQLServer
				sqlserver.Parse(instance.(map[string]interface{}))
				sqlserver.ApplyToTerraform(outputFile)
			}
		case azure.COSMOSDB:
			for _, instance := range dbDetails {
				var cosmosdb azure.CosmosDB
				cosmosdb.Parse(instance.(map[string]interface{}))
				cosmosdb.ApplyToTerraform(outputFile)
			}
		case gcp.BIGTABLE:
			for _, instance := range dbDetails {
				var bigtable gcp.Bigtable
				bigtable.Parse(instance.(map[string]interface{}))
				bigtable.ApplyToTerraform(outputFile)
			}
		case gcp.FIRESTORE:
			for _, instance := range dbDetails {
				var firestore gcp.Firestore
				firestore.Parse(instance.(map[string]interface{}))
				firestore.ApplyToTerraform(outputFile)
			}
		}
	}
}
