package website

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/SabinaSoderstjerna/sabinatech/internal/handler"
	"github.com/SabinaSoderstjerna/sabinatech/internal/website/aboutpage"
)

func InitAboutHandler() *handler.Handler {
	fmt.Printf("Init AboutHandler")
	aboutData := aboutpage.NewAboutPage(filepath.Join("src", "about.json"))
	aboutData.Education = aboutData.GetEducations(filepath.Join("src", "education"))
	aboutData.VocationalExperience = aboutData.GetExperience(filepath.Join("src", "vocational"))
	aboutData.NonProfitExperience = aboutData.GetExperience(filepath.Join("src", "nonprofit"))
	return &handler.Handler{Pattern: "/about", Page: "about.html", Data: aboutData}
}

func InitIndexHandler() *handler.Handler {
	fmt.Printf("Init IndexHandler")
	return &handler.Handler{Pattern: "/", Page: "index.html"}
}

func InitHTTPMux(handlers []*handler.Handler) *http.ServeMux {
	fmt.Println("Init HTTP server mux")
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	for _, h := range handlers {
		mux.HandleFunc(h.Pattern, h.ServeHTTP)
	}
	return mux
}

func InitHTTPServer(mux *http.ServeMux) *http.Server {
	fmt.Println("Init HTTP server")
	return &http.Server{Addr: "0.0.0.0:8080", Handler: mux}
}
