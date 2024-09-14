package wrapunwrap

import (
	"encoding/json"
	"fmt"
	"os"
)

type loadConfigError struct {
	msg string
	err error
}

func (e *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", e.msg, e.err.Error())
}

func (e *loadConfigError) Unwrap() error {
	return e.err
}

type Config struct{}

func LoadConfig(confgFilePath string) (*Config, error) {
	var cfg *Config
	data, err := os.ReadFile(confgFilePath)
	if err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("read file `%s`", confgFilePath), err: err}
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, &loadConfigError{msg: fmt.Sprintf("parse config file `%s`", confgFilePath), err: err}
	}
	return cfg, nil
}
