package providermaker

import "strings"

type Func struct {
	Name     string
	PkgAlias string
}

func (f Func) Render() string {
	slice := make([]string, 0, 2)
	if len(f.PkgAlias) != 0 {
		slice = append(slice, f.PkgAlias)
	}
	slice = append(slice, f.Name)
	return strings.Join(slice, ".")
}
