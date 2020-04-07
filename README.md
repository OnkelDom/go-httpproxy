# go-httpproxy
Lightweight GO Webproxy with Basic Auth

## Go Get
```
go get -u github.com/OnkelDom/go-httpproxy
```

## Go Build
```
go build proxy.go
```

## Startup Commands
```
./proxy -h
usage: proxy [<flags>]

Flags:
  -h, --help  Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":8080"  
              server address
      --auth.username="webproxy"  
              auth username
      --auth.password="webproxy"  
              auth password
```
