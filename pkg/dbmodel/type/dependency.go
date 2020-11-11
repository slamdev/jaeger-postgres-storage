package mtype

type Dependency struct {
	Parent    string `json:"parent"`
	Child     string `json:"child"`
	CallCount int64  `json:"call_count"`
	Source    string `json:"source"`
}

func (d *Dependency) TypeName() string {
	return "dependency"
}
