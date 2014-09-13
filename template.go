package template_metrics

import (
	"html/template" // ToDo: how can I change this from outside?
	"io"
	"time"
)

type Template struct {
	Original *template.Template
	*Metrics
}

func newTemplate(name string, template *template.Template) *Template {
	return &Template{
		template,
		newMetrics(name),
	}
}

// instrument template.Execute
func (proxy *Template) Execute(wr io.Writer, data interface{}) error {
	var startTime time.Time
	if Enable {
		startTime = time.Now()
	}
	error := proxy.Original.Execute(wr, data)
	if Enable {
		// treat as no base name
		defer proxy.measure(startTime, "")
	}
	return error
}

// instrucment template.ExecuteTemplate
func (proxy *Template) ExecuteTemplate(wr io.Writer, base string, data interface{}) error {
	var startTime time.Time
	if Enable {
		startTime = time.Now()
	}
	error := proxy.Original.ExecuteTemplate(wr, base, data)
	if Enable {
		defer proxy.measure(startTime, base)
	}
	return error
}
