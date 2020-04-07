package main

import (
    "log"
    "net/http"
    "os"
    "github.com/go-httpproxy/httpproxy"
    "gopkg.in/alecthomas/kingpin.v2"
)

var (
    listenAddress = kingpin.Flag("web.listen-address", "server address").Default(":8080").String()
    authUser = kingpin.Flag("auth.username", "auth username").Default("webproxy").String()
    authPass = kingpin.Flag("auth.password", "auth password").Default("webproxy").String()
)

func OnError(ctx *httpproxy.Context, where string,
	err *httpproxy.Error, opErr error) {
	// Log errors.
	log.Printf("ERROR: %s: %s [%s]", where, err, opErr)
}

func OnAccept(ctx *httpproxy.Context, w http.ResponseWriter,
	r *http.Request) bool {
	// Handle local request has path "/info"
	if r.Method == "GET" && !r.URL.IsAbs() && r.URL.Path == "/info" {
		w.Write([]byte("GO-HTTP Proxy by Team Monitoring to Scrape Prometheus Endpoints"))
		return true
	}
	return false
}

func OnAuth(ctx *httpproxy.Context, authType string, user string, pass string) bool {
	// Auth test user.
	if user == *authUser && pass == *authPass {
		return true
	}
	return false
}

func OnConnect(ctx *httpproxy.Context, host string) (
	ConnectAction httpproxy.ConnectAction, newHost string) {
	// Apply "Man in the Middle" to all ssl connections. Never change host.
	return httpproxy.ConnectMitm, host
}

func OnRequest(ctx *httpproxy.Context, req *http.Request) (
	resp *http.Response) {
	// Log proxying requests.
	log.Printf("INFO: %d %d: %s %s", ctx.SessionNo, ctx.SubSessionNo, req.Method, req.URL.String())
	return
}

func OnResponse(ctx *httpproxy.Context, req *http.Request,
	resp *http.Response) {
	// Add header "Via: go-httpproxy".
	resp.Header.Add("Via", "HTTP Proxy")
}

func main() {
    kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.SetOutput(os.Stdout)
	log.Print("Started v1.0")
	log.Print("Auth User: ", *authUser)
	log.Print("Auth Pass: ", *authPass)
	log.Print("Listening HTTP ", *listenAddress)

	// Create a new proxy with default certificate pair.
	prx, _ := httpproxy.NewProxy()

	// Set proxy handlers.
	prx.OnError = OnError
	prx.OnAccept = OnAccept
	prx.OnAuth = OnAuth
	prx.OnConnect = OnConnect
	prx.OnRequest = OnRequest
	prx.OnResponse = OnResponse
	//prx.MitmChunked = false

    http.ListenAndServe(*listenAddress, prx)
}
