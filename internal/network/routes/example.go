package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type exampleImpl struct{}

func registerExample(r chi.Router) {
	impl := &exampleImpl{}

	r.Get("/example", impl.sayExample)
}

func (impl *exampleImpl) sayExample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's an example route. Seems everything is working fine!"))
}
