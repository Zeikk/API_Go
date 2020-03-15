package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct{

	Title string `json:"title"`
	Desc string `json:"desc"`
}

type listArticles struct{
	Articles []Article
}

var Articles []Article

//Fonction d'affichage des différentes pages
func home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Welcome to the Home Page")
	fmt.Println("Arrivée sur la page")
}

func allArticles(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "<h1>Listes des articles </h1>")
	for _, article := range Articles{
		fmt.Fprintf(w, "<p>Titre: %s</p>", article.Title)
		fmt.Fprintf(w, "<p>Description: %s</p>", article.Desc)
	}
}

func affichageParams(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	title := vars["title"]

	var find = false

	for _, v := range Articles{
		if (v.Title == title) {
			find = true
		}
	}

	if find {
		fmt.Fprintf(w, "L'article cherché est le %s", title)
	}else{
		fmt.Fprintf(w, "Le paramètre passé n'est pas un titre d'article")
	}
}

func (list *listArticles) ServeHTTP(w http.ResponseWriter, r *http.Request){

	fmt.Println("Liste des Articles: \n")
	for _, 	art:= range list.Articles {
		fmt.Fprintf(w, "\nTitre de l'article: %s\n", art.Title)
		fmt.Fprintf(w, "Description de l'article: %s\n", art.Desc)
		fmt.Printf("%+v\n", art)
	}
	json.NewEncoder(w).Encode(list)
	
}

func createPage(w http.ResponseWriter, r *http.Request){

	fmt.Printf("Arrivée page création\n")
	fmt.Fprintf(w, "<h1>Création d'article</h1>")
	fmt.Fprintf(w, "<form action='/create' method='POST'>"+
		"<label>Titre </label>"+
		"<input type='text' name='Title' required />"+
		"<br/>"+
		"<br/>"+
		"<label>Description </label>"+
		"<textarea name='Desc' required></textarea>"+
		"<button type='submit'>Envoyer</button>"+
	"</form>")
}

func createArticle(w http.ResponseWriter, r *http.Request){

	fmt.Printf("createArticle\n")
	var article Article 
	json.NewDecoder(r.Body).Decode(&article)
	

	fmt.Fprintf(w, "+%v", article)
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

//Fonction permettant le routage
func handle(art [2]Article){

	router := mux.NewRouter()

	list := listArticles{
		Articles: []Article {
			art[0],
			art[1],
		},
	}

	fmt.Printf("%+v\n", list)
	router.HandleFunc("/", home)
	router.Handle("/list", &list)
	router.HandleFunc("/article", allArticles)
	router.HandleFunc("/create", createArticle).Methods("POST")
	router.HandleFunc("/create", createPage).Methods("GET")
	router.HandleFunc("/article/{title}", affichageParams)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main(){

	articles := [...]Article{
		Article{Title: "Article1", Desc: "Desc1"},
		Article{Title: "Article2", Desc: "Desc2"},
	}

	Articles = []Article{
		Article{Title: "Article1", Desc: "Desc1"},
		Article{Title: "Article2", Desc: "Desc2"},
	}
	handle(articles)
}