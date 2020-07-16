package main

// arm build: CGO_ENABLED=0 GOOS=linux GOARCH=arm go build $0

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)


func myPost(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body)
	s1 := strings.Fields(string(s))
	s2 := s1[1:]
	fmt.Printf("run command: %s  args: %s\n",s1[0],s2)
	cmd := exec.Command(string(s1[0]),s2...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, string(out))
}

func main() {

	fmt.Println("web server listen on :8080\n")
	http.HandleFunc("/shell", myPost)
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
