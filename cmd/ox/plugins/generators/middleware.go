package generators

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gobuffalo/flect"
)

//go:embed middleware.go.tmpl
var middlewareTemplate string

type Middleware struct{}

func (md Middleware) Name() string {
	return "generators/middleware"
}

func (md Middleware) InvocationName() string {
	return "middleware"
}

func (md Middleware) Generate(ctx context.Context, root string, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("middleware requires a name")
	}

	tmpl, err := template.New("middleware").Parse(middlewareTemplate)
	if err != nil {
		return err
	}

	path := filepath.Join(root, "app", "middleware", flect.Underscore(args[2])+".go")
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, map[string]interface{}{
		"Name": flect.Pascalize(args[2]),
	})

	if err != nil {
		return err
	}

	return nil
}
