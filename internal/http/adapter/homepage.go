package adapter

import (
	"go-webapp/internal/template"
	"go-webapp/internal/template/page"
	"net/http"
)

type HomePageAdapter struct {
	renderer template.Renderer
}

func (h HomePageAdapter) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	p := page.Home{
		Text: "Hello world!",
	}

	err := h.renderer.Render(p, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NewHomePageAdapter(r template.Renderer) HomePageAdapter {
	return HomePageAdapter{
		renderer: r,
	}
}
