package html

import (
	"html/template"
	"net/http"
	"path"
)

type Config struct {
	BaseDir      string
	BaseTemplate string
}

type Option func(*Config)

func BaseDir(dir string) Option {
	return func(c *Config) {
		c.BaseDir = dir
	}
}

func BaseTemplate(name string) Option {
	return func(c *Config) {
		c.BaseTemplate = name
	}
}

func Generate(w http.ResponseWriter, data interface{}, fileNames []string, options ...Option) {
	c := Config{
		BaseDir:      "templates",
		BaseTemplate: "",
	}

	for _, option := range options {
		option(&c)
	}

	for i, fileName := range fileNames {
		fileNames[i] = path.Join(c.BaseDir, fileName)
	}
	t := template.Must(template.ParseFiles(fileNames...))
	if c.BaseTemplate != "" {
		t.ExecuteTemplate(w, c.BaseTemplate, data)
	} else {
		t.Execute(w, data)
	}
}
