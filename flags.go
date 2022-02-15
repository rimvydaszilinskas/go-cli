package cli

import (
	"flag"
	"fmt"
)

type ValidationFunction = func(interface{}) error

type Flag struct {
	Name        string
	Type        string
	Description string
	Validation  *ValidationFunction
	Default     interface{}
	Value       interface{}
}

func (f *Flag) validate() error {
	if err := f.validateValueType(f.Value); err != nil {
		return err
	}

	if f.Validation != nil {
		validationFunction := *f.Validation
		return validationFunction(f.Value)
	}
	return nil
}

func (f *Flag) init() error {
	if f.Default != nil {
		if err := f.validateValueType(f.Default); err != nil {
			return err
		}
	} else {
		switch f.Type {
		case "int":
			f.Default = -1
		case "float":
			f.Default = float64(0)
		case "string":
			f.Default = ""
		case "bool":
			f.Default = false
		}
	}

	switch f.Type {
	case "bool":
		f.Value = flag.Bool(f.Name, f.Default.(bool), f.Description)
	case "int":
		f.Value = flag.Int(f.Name, f.Default.(int), f.Description)
	case "string":
		f.Value = flag.String(f.Name, f.Default.(string), f.Description)
	case "float":
		f.Value = flag.Float64(f.Name, f.Default.(float64), f.Description)
	default:
		return fmt.Errorf("invalid flag type %s", f.Type)
	}
	return nil
}

func (f *Flag) GetBool() bool {
	return *f.Value.(*bool)
}

func (f *Flag) GetInt() int {
	return *f.Value.(*int)
}

func (f *Flag) GetString() string {
	return *f.Value.(*string)
}

func (f *Flag) GetFloat() float64 {
	return *f.Value.(*float64)
}

func (f *Flag) IsInt() bool {
	return f.Type == "int"
}

func (f *Flag) IsBool() bool {
	return f.Type == "bool"
}

func (f *Flag) IsFloat() bool {
	return f.Type == "float"
}

func (f *Flag) IsString() bool {
	return f.Type == "string"
}

func (f *Flag) validateValueType(value interface{}) error {
	errorMsgFormat := "[%s] invalid value %s for %s"
	switch value.(type) {
	case int:
		if f.Type != "int" {
			return fmt.Errorf(errorMsgFormat, f.Name, value, f.Type)
		}
	case float64:
		if f.Type != "float" {
			return fmt.Errorf(errorMsgFormat, f.Name, value, f.Type)
		}
	case string:
		if f.Type != "string" {
			return fmt.Errorf(errorMsgFormat, f.Name, value, f.Type)
		}
	case bool:
		if f.Type != "bool" {
			return fmt.Errorf(errorMsgFormat, f.Name, value, f.Type)
		}
	}
	return nil
}
