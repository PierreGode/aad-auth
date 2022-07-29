package passwd

import (
	"github.com/ubuntu/aad-auth/internal/cache"
	"gopkg.in/yaml.v3"
)

// SetCacheOption set opts everytime we open a cache.
// This is not compatible with parallel testing as it needs to change a global state.
func SetCacheOption(opts ...cache.Option) {
	testopts = opts
}

type publicPasswd struct {
	Name   string
	Passwd string
	UID    uint
	GID    uint
	Gecos  string
	Dir    string
	Shell  string
}

// MarshalYAML use a public object to Marhsal to a yaml format.
func (p Passwd) MarshalYAML() (interface{}, error) {
	return publicPasswd{
		Name:   p.name,
		Passwd: p.passwd,
		UID:    p.uid,
		GID:    p.gid,
		Gecos:  p.gecos,
		Dir:    p.dir,
		Shell:  p.shell,
	}, nil
}

// UnmarshalYAML use a public object to Unmarhsal to.
func (p *Passwd) UnmarshalYAML(value *yaml.Node) error {
	o := publicPasswd{}
	err := value.Decode(&o)
	if err != nil {
		return err
	}

	*p = Passwd{
		name:   o.Name,
		passwd: o.Passwd,
		uid:    o.UID,
		gid:    o.GID,
		gecos:  o.Gecos,
		dir:    o.Dir,
		shell:  o.Shell,
	}
	return nil
}