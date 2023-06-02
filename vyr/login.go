package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	
    "github.com/julienschmidt/httprouter"
)


var users = map[string]string{"a": "a", "j": "j"}
var store = sessions.NewCookieStore([]byte("secret_key"))

func loginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "login.html", nil)
	}
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	pwd, exists := users[username]
	if exists {
		session, _ := store.Get(r, "session.id")
		if pwd == password {
			session.Values["authenticated"] = true 
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
//		w.Write([]byte("Login successfully!"))
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	session, _ := store.Get(r, "session.id")
	session.Values["authenticated"] = false 
	session.Save(r, w)
//	w.Write([]byte("Logout Successful"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Dash(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Welcome!"))
}

func isLoggedIn(r *http.Request) bool {
	session, _ := store.Get(r, "session.id")
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated != false {
		return true
	}
	return false
}