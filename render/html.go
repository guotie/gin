// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"net/http"

	"github.com/flosch/pongo2"
)

type (
	HTMLRender struct {
		*pongo2.TemplateSet
	}
	HTMLTemplate struct {
		tpl     *pongo2.Template
		context pongo2.Context
	}
)

var htmlContentType = []string{"text/html; charset=utf-8"}

// LoadTemplate load template from dir
func LoadTemplate(dir string, debug bool) *HTMLRender {
	loader := pongo2.MustNewLocalFileSystemLoader(dir)
	ts := pongo2.NewSet("gin", loader)
	ts.Debug = debug
	return &HTMLRender{ts}
}

// Instance return HTMLTemplate
func (r *HTMLRender) Instance(name string, context map[string]interface{}) (*HTMLTemplate, error) {
	tpl, err := r.FromCache(name)
	if err != nil {
		return nil, err
	}

	return &HTMLTemplate{tpl, pongo2.Context(context)}, nil
}

// Render render pongo2 template
func (r *HTMLTemplate) Render(w http.ResponseWriter) error {
	err := r.tpl.ExecuteWriter(r.context, w)
	return err
}
