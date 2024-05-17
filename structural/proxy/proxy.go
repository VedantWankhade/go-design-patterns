package proxy

import (
	"fmt"
	"slices"
)

type server interface {
	serveHTTP(string, string)
}

type application struct{}

func (a *application) serverHTTP(url, method string) {
	fmt.Printf("Serving %s request at %v\n", method, url)
}

type nginx struct {
	a              application
	methodsAllowed []string
}

func (n *nginx) serveHTTP(url, method string) {
	if !slices.Contains(n.methodsAllowed, method) {
		fmt.Println("NOT ALLOWED")
		return
	}
	n.a.serverHTTP(url, method)
}
