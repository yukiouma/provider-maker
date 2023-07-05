package providermaker

import (
	"bytes"
	"go/ast"
	"go/format"
	"os"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Provider struct {
	Imports map[string]Import
	Funcs   []Func
}

type Render struct {
	Imports []string
	Pkg     string
	Funcs   []string
}

func NewProvider() *Provider {
	return &Provider{
		Imports: make(map[string]Import),
		Funcs:   make([]Func, 0),
	}
}

func (p *Provider) Generate(dst string) error {
	dst, err := toAbs(dst)
	if err != nil {
		return err
	}
	sb := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(sb, p.ToRender(getDerivePackageName(dst)))
	s := sb.Bytes()
	if err != nil {
		return err
	}
	s, err = format.Source(s)
	if err != nil {
		return err
	}
	dst = path.Join(dst, "provider.go")
	newFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = newFile.Write(s)
	return err
}

func (pr *Provider) ToRender(pkgName string) Render {
	r := Render{
		Pkg:     pkgName,
		Imports: make([]string, 0, len(pr.Imports)),
		Funcs:   make([]string, 0, len(pr.Funcs)),
	}
	for _, p := range pr.Imports {
		r.Imports = append(r.Imports, p.Render())
	}
	r.Imports = append(r.Imports, "\"github.com/google/wire\"")
	for _, f := range pr.Funcs {
		r.Funcs = append(r.Funcs, f.Render())
	}
	return r
}

func (pr *Provider) Analyse(dir string) error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedFiles,
	}
	pkgs, err := packages.Load(cfg, dir)
	if err != nil {
		return err
	}
	for _, p := range pkgs {
		for _, s := range p.Syntax {
			for _, decl := range s.Decls {
				f, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}
				if f.Doc == nil {
					continue
				}
				if !hasProviderTag(f.Doc.List...) {
					continue
				}
				pr.AddImport(Import{Alias: p.Name, Path: p.PkgPath})
				pr.AddFunc(Func{PkgAlias: p.Name, Name: f.Name.Name})
			}
		}
	}
	return nil
}

func (pr *Provider) AddImport(i Import) {
	pr.Imports[i.Path] = i
}

func (pr *Provider) AddFunc(f Func) {
	pr.Funcs = append(pr.Funcs, f)
}

func getDerivePackageName(path string) string {
	slice := strings.Split(path, "/")
	if size := len(slice); size > 0 {
		return slice[size-1]
	}
	return ""
}

func hasProviderTag(comments ...*ast.Comment) bool {
	for _, c := range comments {
		if c.Text == "// +provider" {
			return true
		}
	}
	return false
}
