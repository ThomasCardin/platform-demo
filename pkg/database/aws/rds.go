package aws

import (
	"github.com/ThomasCardin/ci-cd/pkg/database"
)

type RDS struct {
	DbName           string `yaml:"dbName"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	InstanceClass    string `yaml:"instanceClass"`
	AllocatedStorage int    `yaml:"allocatedStorage"`
}

var _ database.Database = (*RDS)(nil)

func (rds *RDS) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type d RDS
	aux := &d{}
	if err := unmarshal(aux); err != nil {
		return err
	}
	*rds = RDS(*aux)
	return nil
}
