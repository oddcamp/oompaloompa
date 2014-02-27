package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func logHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v: %v %v", r.Method, r.URL.Path, r.PostFormValue("project"))
		next.ServeHTTP(w, r)
	}
}

func logHandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return logHandler(next)
}

type Project struct {
	Name, Path string
}
type Configuration struct {
	Projects []Project
}

type Payload struct {
	Repository struct {
		Name string
	}
}

var configuration = &Configuration{}
var port = flag.Int("port", 3000, "listen on port")

func main() {
	flag.Parse()
	log.Println("Oompa Loompa warming up...")
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	http.Handle("/deploy", logHandlerFunc(deploy))
	http.Handle("/", logHandlerFunc(home))

	portString := strconv.Itoa(*port)
	log.Println("Listening on port " + portString + "...")
	http.ListenAndServe(":"+portString, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Oompa Loompa reporting for duty!"))
}

func projectIndex(t string, vs []Project) int {
	for i, v := range vs {
		if v.Name == t {
			return i
		}
	}
	return -1
}

func deploy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("Payload received.")

		payload := &Payload{}
		var decoder = json.NewDecoder(strings.NewReader(r.PostFormValue("payload")))
		decoder.Decode(&payload)

		index := projectIndex(payload.Repository.Name, configuration.Projects)

		if index < 0 {
			log.Println("Invalid payload.")
			return
		}

		log.Println("Deploying '" + configuration.Projects[index].Name + "'...")
		err := os.Chdir(configuration.Projects[index].Path)
		if err != nil {
			log.Println(err.Error())
			return
		}

		out, err := exec.Command("git", "pull").Output()
		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(string(out[:]))
	}
}
