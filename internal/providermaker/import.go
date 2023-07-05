package providermaker

import (
	"fmt"
	"strings"
)

type Import struct {
	Alias string
	Path  string
}

func (i Import) Render() string {
	slice := make([]string, 0, 2)
	// if len(i.Alias) != 0 {
	// 	slice = append(slice, i.Alias)
	// }
	slice = append(slice, qout(i.Path))
	return strings.Join(slice, " ")
}

func qout(source string) string {
	return fmt.Sprintf(`"%s"`, source)
}
