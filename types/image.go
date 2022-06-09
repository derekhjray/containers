package types

import "time"

type Image struct {
	ID        string
	Digest    string
	Size      int64
	Env       []string
	User      string
	Labels    map[string]string
	RepoTags  []string
	Layers    []string
	CreatedAt time.Time
}

type ImageHistory struct {
	ID   string
	Cmd  string
	Size int64
}
