package proxy

import "testing"

func TestProxy(t *testing.T) {
	serverApp := &application{}
	serverApp.serverHTTP("/hello", "GET")
	serverApp.serverHTTP("/hello", "DELETE") // works

	nginxProxy := &nginx{a: *serverApp, methodsAllowed: []string{"GET", "POST"}}
	nginxProxy.serveHTTP("/hello", "GET")
	nginxProxy.serveHTTP("/hello", "DELETE") // not allowed
}
