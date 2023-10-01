package database

import (
	"os"
)

type Database interface {
	Parse(data map[string]interface{})
	ApplyToTerraform(outputFile *os.File) error
}
