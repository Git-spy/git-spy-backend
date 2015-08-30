package main

import (
  "github.com/google/go-github/github"
  "github.com/ant0ine/go-json-rest/rest"
  "fmt"
  "log"
  "net/http"
)

func main() {
  api := rest.NewApi()

  api.Use(rest.DefaultDevStack...)

  api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
    username := "zzarcon"
    client := github.NewClient(nil)
    orgs, _, _ := client.Organizations.List(username, nil)

    fmt.Println(orgs)
    w.WriteJson(map[string]string{"Body": "Hello World!"})
  }))

  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}