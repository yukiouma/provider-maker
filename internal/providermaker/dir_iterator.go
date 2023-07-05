package providermaker

import (
	"os"
	"path/filepath"
)

type DirIterator struct {
	dch     chan string
	errch   chan error
	workend bool
	filters []FilterFunc
}

func NewDirIterator(dir string, filters ...FilterFunc) (*DirIterator, error) {
	dir, err = toAbs(dir)
	if err != nil {
		return nil, err
	}
	f := &DirIterator{
		dch:     make(chan string, 10),
		errch:   make(chan error, 10),
		filters: filters,
	}
	go func() {
		if err := filepath.WalkDir(dir, f.visit); err != nil {
			f.errch <- err
		}
		f.workend = true
	}()
	go func() {
		for {
			if f.workend && len(f.dch) == 0 {
				close(f.dch)
				break
			}
		}
	}()
	return f, nil
}

func (di *DirIterator) visit(path string, info os.DirEntry, err error) error {
	if err != nil {
		di.errch <- err
		return err
	}
	if info.IsDir() {
		for _, f := range di.filters {
			if !f(path) {
				return nil
			}
		}
		di.dch <- path
	}
	return nil
}

func (di *DirIterator) Next() (string, bool) {
	p, ok := <-di.dch
	return p, ok
}
