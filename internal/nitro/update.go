package nitro

import "errors"

func Update(name string) (*Action, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &Action{
		Type:       "exec",
		UseSyscall: false,
		Args:       []string{"exec", name, "--", "sudo", "apt", "update"},
	}, nil
}

func Upgrade(name string) (*Action, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &Action{
		Type:       "exec",
		UseSyscall: false,
		Args:       []string{"exec", name, "--", "sudo", "apt", "upgrade", "-y"},
	}, nil
}
