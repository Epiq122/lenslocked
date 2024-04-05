package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the danger zones!</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">robgleason122@gmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, ` <h1>🤔 Frequently Asked Questions 🤔</h1>
	<ul>
			<li>
					<b>Is there a free version?</b>
					<p>Yep! Our free version is so good that you'll wonder why anyone would ever need the paid one. Just kidding, the paid version has unicorns 🦄</p>
			</li>
			<li>
					<b>What are your support hours?</b>
					<p>Our support team is fueled by coffee ☕ and memes 😂, so they're available 24/7. But beware, if you ask for help during a full moon, you might get a werewolf instead. 🐺</p>
			</li>
			<li>
					<b>How do I contact support?</b>
					<p>Summon them by shouting "Supporto Patronum!" into a mirror three times. Or just email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a></p>
			</li>
	</ul>
`)

}

type Router struct{}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		homeHandler(w, req)
	case "/contact":
		contactHandler(w, req)
	case "/faq":
		faqHandler(w, req)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)

}
