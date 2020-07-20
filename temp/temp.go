// Package github.com/xen0ne/riceutils/temp provides functions for applying Pallets to golang templates
package temp

import (
	"io/ioutil"
	"text/template"
	"io"
	"github.com/xen0ne/riceutils/color"
)

// ApplyPallet reads a template from a reader, applies the given pallet to it
// and then writes the filled in template to the writer.
func ApplyPallet(r io.Reader, p color.Pallet, w io.Writer) error {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"hex": color.Hex,
		"rgb": color.Rgb,
	}

	tmpl := template.Must(template.New("test").Funcs(funcs).Parse(string(b)))

	err = tmpl.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}
