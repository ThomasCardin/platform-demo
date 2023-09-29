package database

/*
To use this package (databases),
import this one into the other. So you can
use the interface for another struct
*/
type Database interface {
	UnmarshalYAML(unmarshal func(interface{}) error) error
}
