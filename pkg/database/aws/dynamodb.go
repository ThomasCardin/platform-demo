package aws

import (
	"github.com/ThomasCardin/ci-cd/pkg/database"
)

type DynamoDB struct {
	TableName     string `yaml:"tableName"`
	HashKeyName   string `yaml:"hashKeyName"`
	HashKeyType   string `yaml:"hashKeyType"`
	ReadCapacity  int    `yaml:"readCapacity"`
	WriteCapacity int    `yaml:"writeCapacity"`
}

var _ database.Database = (*DynamoDB)(nil)

func (dynamoDb *DynamoDB) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type d DynamoDB
	aux := &d{}
	if err := unmarshal(aux); err != nil {
		return err
	}
	*dynamoDb = DynamoDB(*aux)
	return nil
}
