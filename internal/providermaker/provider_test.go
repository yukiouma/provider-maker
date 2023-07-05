package providermaker

import "testing"

func TestProvider(t *testing.T) {
	p := NewProvider()
	p.Analyse("/root/playground/golang/misc-playground/provider-maker/examples/demo/usecase")
	p.Analyse("/root/playground/golang/misc-playground/provider-maker/examples/demo/adapter/datasource")
	if err := p.Generate("/root/playground/golang/misc-playground/provider-maker/examples/demo/infra/constructor"); err != nil {
		t.Fatal(err)
	}
}
