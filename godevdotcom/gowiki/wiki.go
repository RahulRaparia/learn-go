package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body []byte
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) 
// }

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}


func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// HARD CODED HTML
// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title, Body: []byte("Error Occured while Loading Page")}
// 	}
// 	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
// 		"<form action=\"/save/%s\" method=\"POST\">"+
// 		"<textarea name=\"body\">%s</textarea><br>"+
// 		"<input type =\"submit\" value=\"Save\">"+
// 		"</form>",	
// 		p.Title, p.Title, p.Body)
// }

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title, Body: []byte("Start Typing here to create a new page")}
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
    renderTemplate(w, "view", p)
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])	
	}
}

func renderTemplate(w http.ResponseWriter,tmpl string, p *Page){
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main () {
	//saving test pages
	// saveTestPages("TestPage2", "This is a sample Pagee.")
	

	http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

