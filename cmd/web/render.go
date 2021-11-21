package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSFRToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

// go:embed templates

var templateFS embed.FS

// this will use to add any info we need from templateData struct
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

func (app *application) RenderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	// build our template to render
	templateToRender := fmt.Sprintf("templates/%s.page.html", page)

	//Now lets check if templateCache exist
	_, templateInMap := app.templateCache[templateToRender]

	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			return err
		}
	}
	// build out template data 
	if td == nil {
		td = &templateData{}
	}
	td = app.addDefaultData(td, r)

   //Execute the template 
   err = t.Execute(w, td)
   if err != nil {
	   app.errorLog.Println(err)
	   panic(err)
   }
	return nil
}

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var err error
	var t *template.Template

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.html", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.html", page)).Funcs(functions).ParseFS(templateFS, "template/base.layout.html", strings.Join(partials, ","), templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return nil, err
		}
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.html", page)).Funcs(functions).ParseFS(templateFS, "template/base.layout.html", templateToRender)
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
