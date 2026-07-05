package cache

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

type CachingServer struct {
	cacheDir string
	listener net.Listener
}

func CreateCachingServer() CachingServer {
	cacheDir := createCachDir()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(fmt.Sprintf("Failed to create caching listener\n%s", err.Error()))
	}
	result := CachingServer{
		cacheDir: cacheDir,
		listener: listener,
	}

	return result
}

func (c CachingServer) handler(w http.ResponseWriter, req *http.Request) {
	url, err := url.QueryUnescape(req.URL.Query().Get("url"))
	if err != nil {
		fmt.Println("Malformed url query parameter", req.URL.Query().Get("url"), err.Error())
		w.WriteHeader(400)
		w.Write([]byte("Malformed url query parameter"))
		return
	}

	file, err := c.getFileFromCache(url)
	if err == nil {
		w.WriteHeader(200)
		w.Write(file)
		return
	}

	res, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch", url)
		w.WriteHeader(500)
		w.Write([]byte("Failed to fetch"))
		return
	}

	if res.StatusCode != 200 {
		fmt.Println("Failed to fetch", url, "Status code", res.StatusCode)
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Failed to read response body", err.Error())
			responseBody = []byte("Failed to read response body")
		}
		w.WriteHeader(res.StatusCode)
		w.Write(responseBody)
		return
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body")
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Failed to read response body\n%s", err.Error())))
		return
	}

	err = c.storeFileInCache(url, responseBody)
	if err != nil {
		fmt.Printf("Failed to save %s in cache\n%s\n", url, err.Error())
	}

	w.WriteHeader(500)
	w.Write(responseBody)
}

func (c CachingServer) GetUrl() string {
	return fmt.Sprintf(`http://%s`, c.listener.Addr().String())
}

func (c CachingServer) Run() {
	fmt.Println("Starting caching server")
	http.Serve(c.listener, http.HandlerFunc(c.handler))
}

func (c CachingServer) Close() {
	fmt.Println("Closing caching server")
	c.listener.Close()
}
