package loadBalancer

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from backend %s\n", port)
	})

	log.Println("Backend running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	http.Client{}
}
