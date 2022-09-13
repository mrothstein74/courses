package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>" )
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:mrothsteinbills@gmail.com\">mrothsteinbills@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
		<ul>
			<li>Q: Is there a free version?</li>
			<li>A: Yes! We offer a free trial for 30 days on any paid plans.</li>
			<br>
			<li>Q: What are your support hours?</li>
			<li>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</li>
			<br>
			<li>Q: How can I conttact support?</li>
			<li>A: Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a></li>
		</ul>
		`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}