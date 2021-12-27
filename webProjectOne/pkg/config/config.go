package config

import "html/template"

// AppConfig
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}