package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "strings"
)

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

type reqObj struct{
Name string
}

type resObj struct{
Greeting string
}
func phello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    decoder := json.NewDecoder(req.Body)
    var t reqObj  
    err := decoder.Decode(&t)
    if err != nil {
        fmt.Println("Error")
    }
    //fmt.Println(t.Name)

    s := []string{"Hello, ",t.Name}
    g := resObj{strings.Join(s,"")}
    js,err := json.Marshal(g)
    if err != nil{
	fmt.Println("Error")
	return
	}
    rw.Header().Set("Content-Type","application/json")
    rw.Write(js)
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/phello",phello)
    server := http.Server{
            Addr:        "0.0.0.0:8083",
            Handler: mux,
    }

    server.ListenAndServe()
}
