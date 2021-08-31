package main

type (
	Plugin struct{}

	Config struct{}
)

func (p Plugin) Exec() error {
	return nil
}
