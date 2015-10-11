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

  router, err := rest.MakeRouter(
    rest.Get("/stats", GetStat),
  )
  if err != nil {
      log.Fatal(err)
  }
  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetStat(w rest.ResponseWriter, r *rest.Request) {
  query := r.URL.Query()
  user_id := query.Get("user_id")
  repo_id := query.Get("repo_id")
  client := github.NewClient(nil)
  orgs, _, _ := client.Repositories.Get(user_id, repo_id)

  fmt.Println(user_id)
  w.WriteJson(orgs)
}