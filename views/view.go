package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

//Render accepts a reciever from the View class and
//renders each of the templates
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//Get all of the template files in the layouts
//directory so we don't have to import each one
//individually.  Uses file globbing for .gohtml
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

//NewView gathers the layout files/templates and parses them
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

//View struct contains pointer to Template as well as
//a Layout which indicates which template is to be used.
type View struct {
	Template *template.Template
	Layout   string
}
