package main

import (
	"github.com/valyala/fasthttp"
	"html/template"
	"math/rand"
	"os"
)

type date struct {
	Response string
}

func getLeapYear() string {
	var response string
	// get random number between 0-3
	r := rand.Intn(4)

	if r%4 == 0 {
		response = "Leap Day"
	} else {
		response = "February 29"
	}
	return response
}

func getRoute(fileRoute string, d bool, ctx *fasthttp.RequestCtx) {
	t, _ := template.ParseFiles(fileRoute)

	// set content type, 200 status and caching for 50 seconds
	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Cache-Control", "stale-while-revalidate=50")

	if d != false {
		getDate := date{Response: getLeapYear()}
		t.Execute(ctx, getDate)
	} else {
		t.Execute(ctx, nil)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// super simple "mux" based on request route
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/ta":
			getRoute("notes/ta.html", false, ctx)
		case "/ac":
			getRoute("notes/ac.html", false, ctx)
		case "/coc":
			getRoute("notes/coc.html", false, ctx)
		case "/":
			getRoute("s.html", true, ctx)
		default:
			ctx.Error("not found :p", fasthttp.StatusNotFound)
		}
	}

	/*
		mux := http.NewServeMux()
		mux.HandleFunc("/ta", ta)
		mux.HandleFunc("/ac", ac)
		mux.HandleFunc("/coc", coc)
		mux.HandleFunc("/", s)

		staticFiles := http.FileServer(http.Dir("./static"))
		mux.Handle("/static/", http.StripPrefix("/static", staticFiles))

		http.ListenAndServe(":"+port, mux)
	*/
	fasthttp.ListenAndServe(":"+port, m)
}
