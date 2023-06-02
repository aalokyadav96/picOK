package main
import (
	"html/template"
	"net/http"
  	"log"
	
    "github.com/julienschmidt/httprouter"
)

const PORT = "localhost:4000"
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	HandleRoutes()
	return
}

func HandleRoutes() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/dash", Dash)

	router.GET("/new", NewPhotoGet)
	router.POST("/upload", NewPhotoPost)
	router.GET("/post/:postid", ShowPost)
	router.GET("/delete/:postid", DeletePost)/*
	router.GET("/user/:name", UserProfile)
	
	router.GET("/register", Register)
	router.POST("/register", Register)
*/	
	router.GET("/login", loginHandler)
	router.POST("/login", loginHandler)
	router.GET("/logout", logoutHandler)


	router.NotFound = http.FileServer(http.Dir(""))
	router.ServeFiles("/usrimg/*filepath", http.Dir("usrimg"))
	router.ServeFiles("/files/*filepath", http.Dir("uploads"))
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	
	log.Println("Starting erver on ", PORT)
	err := http.ListenAndServe(PORT, router)
//err := http.ListenAndServe(GetPort(), router)
 	if err != nil {
		log.Fatal("error starting http server : ", router)
 	}

}

