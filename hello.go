package main

import (
  "github.com/google/go-github/github"
  "github.com/ant0ine/go-json-rest/rest"
  "fmt"
  "strings"
  "log"
  "net/http"
)

func main() {
  api := rest.NewApi()

  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
    rest.Get("/stats/#id", GetStat),
  )
  if err != nil {
      log.Fatal(err)
  }
  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetStat(w rest.ResponseWriter, r *rest.Request) {
  id := r.PathParam("id")
  chunks := strings.Split(id, "-")
  client := github.NewClient(nil)
  orgs, _, _ := client.Repositories.Get(chunks[0], chunks[1])

  fmt.Println(orgs)
  w.WriteJson(orgs)
}