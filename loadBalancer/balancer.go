package loadBalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Balancer struct {
	backends []*url.URL
	current  int
	mutex    sync.Mutex
}

func (lb *Balancer) Next() *url.URL {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	backend := lb.backends[lb.current]
	lb.current = (lb.current + 1) % len(lb.backends)
	return backend
}

func (lb *Balancer) handler(w http.ResponseWriter, r *http.Request) {
	backend := lb.Next()

	proxy := httputil.NewSingleHostReverseProxy(backend)
	proxy.ServeHTTP(w, r)
}

func main() {

	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	var urls []*url.URL
	for _, b := range backends {
		u, err := url.Parse(b)
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, u)
	}

	lb := &Balancer{
		backends: urls,
	}

	http.HandleFunc("/", lb.handler)

	log.Println("Load balancer running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
