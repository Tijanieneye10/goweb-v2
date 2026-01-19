package main

import "net/http"

func main() {
	app := &Application{
		mux: http.NewServeMux(),
	}

	app.mount()
}
