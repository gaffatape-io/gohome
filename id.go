package gohome

import "reflect"

// Identifiable is used to indicate that the instance have an ID method.
type Identifiable interface {
	ID() string
}

// ID returns the ID of the instance.
// If instance implements Identifiable it uses the implementation.
// If not then it uses reflection to get the name of the type.
func ID(i interface{}) string {
	if identifiable, ok := i.(Identifiable); ok {
		return identifiable.ID()
	}

	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Struct {
		return t.Name()
	}

	v := reflect.Indirect(reflect.ValueOf(i))
	t = v.Type()
	id := t.Name()
	if id != "" {
		return id
	}

	panic("unable to find useful id; coding error?")
}
