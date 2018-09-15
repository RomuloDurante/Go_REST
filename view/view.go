package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/RomuloDurante/WordHunter/restApi_GO/model"
)

// createTemplate ...
func createTemplate() *template.Template {
	var basePath = "view/templates"

	tpl := template.Must(template.ParseGlob(basePath + "/*.html"))

	return tpl
}

// HandleView ...
func HandleView(w http.ResponseWriter, r *http.Request) {
	requestFile := r.URL.Path[1:]

	tpl := createTemplate()
	t := tpl.Lookup(requestFile + ".html")

	//***************************************/ TODO: use the model layer to get data
	user := model.GetData()
	fmt.Println(user)
	//***************************************/
	if t != nil {
		err := t.Execute(w, user)
		if err != nil {
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
