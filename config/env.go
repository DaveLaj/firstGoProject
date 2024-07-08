package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type ErrCannotGetEnv struct {
	EnvName string
	Err     error
}

func (e *ErrCannotGetEnv) Error() string {
	return fmt.Sprintf("cannot get env %s: %s", e.EnvName, e.Err)
}

func (e *ErrCannotGetEnv) Unwrap() error {
	return e.Err
}

func LoadEnv(env string) error {
	err := godotenv.Load(env)
	if err != nil {
		return &ErrCannotGetEnv{EnvName: env, Err: err}
	}
	return nil
}
