package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynammicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynammicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynammicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynammicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynammicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// secureHeaders → servemux → application handler
	return standardMiddleware.Then(mux)
}
