package vm

type Vm interface {
	UnmarshalYAML(unmarshal func(interface{}) error) error
}
