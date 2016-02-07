package config

import "errors"

// Config comment
type Config struct {
	P []*parameter
}

type parameter struct {
	Key   string
	Value string
}

func Build(i ...interface{}) (Config, error) {
	var err error
	c := Config{}
	for _, v := range i {
		if _, ok := v.(map[string]string); ok {
			c.P = Parameters(v.(map[string]string))
		} else {
			err = errors.New("empty type")
		}
	}
	return c, err
}

// Parameters comment
func Parameters(m map[string]string) []*parameter {
	var p []*parameter
	for k, v := range m {
		p = append(p, Parameter(k, v))
	}
	return p
}

// Parameter comment
func Parameter(k string, v string) *parameter {
	return &parameter{Key: k, Value: v}
}
