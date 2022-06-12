package types

import "time"

type Container struct {
	ID         string
	Name       string
	Image      string
	ImageID    string
	State      string
	Pid        int
	PidMode    string
	NetMode    string
	Envs       []string
	Labels     map[string]string
	User       string
	Privileged bool
	Memory     int64
	CPUs       string
	IPAddress  string
	Created    time.Time
	Started    time.Time
	Finished   time.Time
}
