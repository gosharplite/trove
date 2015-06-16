package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

type flags struct {
	url  url.URL
	path string
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()*2 + 1)

	f, err := getFlags()
	if err != nil {
		log.Fatalf("flags parsing fail: %v", err)
	}

	fs := http.FileServer(http.Dir(f.path))
	http.Handle("/", http.StripPrefix("/", fs))

	err = http.ListenAndServe(getPort(f.url), nil)
	if err != nil {
		log.Fatalf("ListenAndServe: ", err)
	}
}

func getFlags() (flags, error) {

	u := flag.String("url", "http://localhost:8080", "server")

	p := flag.String("path", "/trove_files", "file folder")

	flag.Parse()

	ur, err := url.Parse(*u)
	if err != nil {
		log.Printf("url parse err: %v", err)
		return flags{}, err
	}

	return flags{*ur, *p}, nil
}

func getPort(u url.URL) string {

	r := u.Host

	if n := strings.Index(r, ":"); n != -1 {
		r = r[n:]
	} else {
		r = ":8080"
	}

	return r
}
