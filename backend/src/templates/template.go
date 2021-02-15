package templates

import (
	"io/ioutil"
	"log"
	"text/template"
)

type TemplateData struct {
	Title string
	Data  interface{}
}

func GetTemplate(path string) *template.Template {
	f, err := Assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	t := template.New(path)
	t, err = t.Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}
	t, err = t.Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func GetLayoutTemplate(name string, paths ...string) *template.Template {
	t := template.New(name)
	for _, path := range paths {
		f, err := Assets.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		t, err = t.Parse(string(b))
		if err != nil {
			log.Fatal(err)
		}
	}
	return t
}
