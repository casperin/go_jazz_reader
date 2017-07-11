package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var domain = viper.GetString("domain")

type Content map[string]interface{}

func renderErrorTpl(w http.ResponseWriter, status int, wrong error) error {
	tpl, err := template.ParseFiles(
		"cmd/jazzd/handlers/tpl-public.html",
		"cmd/jazzd/handlers/error.html",
	)
	if err != nil {
		return err
	}
	log.Printf("%d:\n", status)
	log.Println(wrong)
	return tpl.ExecuteTemplate(w, "html", Content{
		"status": status,
		"error":  wrong,
		"domain": domain,
	})
}

func mustRenderErrorTpl(w http.ResponseWriter, status int, err error) {
	must(renderErrorTpl(w, status, err))
}

func renderPublicTpl(name string, w http.ResponseWriter, content Content) error {
	tpl, err := template.ParseFiles(
		"cmd/jazzd/handlers/tpl-public.html",
		"cmd/jazzd/handlers/"+name+".html",
	)
	if err != nil {
		return err
	}
	content["domain"] = domain
	return tpl.ExecuteTemplate(w, "html", content)
}

func mustRenderPublicTpl(name string, w http.ResponseWriter, content Content) {
	must(renderPublicTpl(name, w, content))
}

func renderLoggedInTpl(name string, w http.ResponseWriter, content Content) error {
	tpl, err := template.ParseFiles(
		"cmd/jazzd/handlers/tpl-logged-in.html",
		"cmd/jazzd/handlers/"+name+".html",
	)
	if err != nil {
		return err
	}
	content["domain"] = domain
	return tpl.ExecuteTemplate(w, "html", content)
}

func mustRenderLoggedInTpl(name string, w http.ResponseWriter, content Content) {
	must(renderLoggedInTpl(name, w, content))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
