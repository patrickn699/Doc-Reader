package main

import (
	//"html/template"
	"log"
	"net/http"
)

// Serve static assets like /styles.css, /script.js directly from ./static
func serveStaticFiles() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/style.css", fs)
	http.Handle("/script.js", fs)
	// Add more if needed (e.g., fonts, images)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodGet {
		// tmpl, err := template.ParseFiles("static/login.html")
		// if err != nil {
		// 	http.Error(w, "Error loading template", http.StatusInternalServerError)
		// 	return
		// }
	http.ServeFile(w, r, "static/login.html")
		// tmpl.Execute(w, nil)
		
	// } else if r.Method == http.MethodPost {
	// 	username := r.FormValue("username")
	// 	password := r.FormValue("password")
	// 	if username == "admin" && password == "password" {
	// 		http.Redirect(w, r, "/home", http.StatusSeeOther)
	// 	} else {
	// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	// 	}
	// }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func rootRedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	// Serve static files
	//serveStaticFiles()

	// Route handlers
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)

	// Redirect root / to /login
	//http.HandleFunc("/", rootRedirectHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
