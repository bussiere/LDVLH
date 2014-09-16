package main

import (
	"github.com/go-martini/martini"

	"bytes"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/gorilla/schema"
	"html/template"
	"net/http"
)

type Form1 struct {
	toto string
}

type Person struct {
	Name string
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", func(r render.Render) string {
		t := template.New("template1")
		p := Person{Name: "Mary"}
		templateh := "<html><body>Hello world! {{.Name}} <form name='input' action='../first' method='post' ><input type='texte' name='toto'><input type='submit' value='Submit'></form></body></html>"
		t, _ = t.Parse(templateh)
		var doc bytes.Buffer
		err := t.Execute(&doc, p)
		if err != nil {
			fmt.Println("There was an error:", err)
		}
		s := doc.String()
		fmt.Println(s)

		return s

	})
	m.Post("/first", func(req *http.Request, r render.Render) string {
		form := new(Form1)
		decoder := schema.NewDecoder()
		err := req.ParseForm()

		if err != nil {
			// Handle error
		}

		decoder = schema.NewDecoder()
		// r.PostForm is a map of our POST form values
		err = decoder.Decode(form, req.PostForm)

		if err != nil {
			// Handle error
		}
		fmt.Println(form.toto)
		templateh := "Hello world! <form name='input' action='../first' method='post' ><input type='texte' name='toto'><input type='submit' value='Submit'></form>"
		return templateh

	})

	m.Run()
}
