package nitro

import "errors"

func Redis(name string) (*Action, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &Action{
		Type:       "exec",
		UseSyscall: true,
		Args:       []string{"exec", name, "--", "redis-cli"},
	}, nil
}
