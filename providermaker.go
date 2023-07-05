package providermaker

import (
	"github.com/yukiouma/provider-maker/internal/providermaker"
)

// in: input dir, which contains the provider tags
//
// out: generated provider file
func Make(in, out string) error {
	iter, err := providermaker.NewDirIterator(in, providermaker.DefaultFilters...)
	if err != nil {
		return err
	}
	p := providermaker.NewProvider()
	for f, ok := iter.Next(); ok; {
		if err := p.Analyse(f); err != nil {
			return err
		}
		f, ok = iter.Next()
	}
	return p.Generate(out)
}
