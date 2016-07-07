package immortal

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
)

type Daemon struct {
	owner   *user.User
	Pidfile string
	Log     string
	Env     map[string]string
	Cmd     string
	Cwd     string
	signals map[string]string
	Status  chan error
	Pid     chan int
}

func New(u *user.User, c, p, l *string) (*Daemon, error) {
	if *c != "" {
		yml_file, err := ioutil.ReadFile(*c)
		if err != nil {
			return nil, err
		}

		var D Daemon

		if err := yaml.Unmarshal(yml_file, &D); err != nil {
			return nil, err
		}

		return &D, nil
	}

	return &Daemon{
		owner:   u,
		Pidfile: *p,
		Log:     *l,
		Status:  make(chan error, 1),
		Pid:     make(chan int, 1),
	}, nil
}