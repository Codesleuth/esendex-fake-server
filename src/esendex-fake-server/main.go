package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/a-palchikov/goji"
  "github.com/a-palchikov/goji/web"
)


var accountsRepository = []Account {
  {"254c2115-88e0-4fd4-97da-f9778e659932", "EX0012345"},
}


func hello(c web.C, w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func getAccounts(c web.C, w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "GET Accounts %d", len(accountsRepository))
}

func getAccount(c web.C, w http.ResponseWriter, r *http.Request) {
  accountId := c.URLParams["accountId"]

  for i := len(accountsRepository) - 1; i >= 0; i-- {
    if accountsRepository[i].ID == accountId {
      w.Header().Set("Content-Type", "application/json")
      accountsRepository[i].Write(w)
      return
    }
  }

  http.Error(w, "Account not found.", 404)
}

func postAccount(c web.C, w http.ResponseWriter, r *http.Request) {
  var account Account

  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&account)

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  accountsRepository = append(accountsRepository, account)

  url := fmt.Sprintf("/accounts/%s", account.ID)
  http.Redirect(w, r, url, http.StatusCreated)
}

func main() {
  goji.Get("/accounts", getAccounts)
  goji.Get("/accounts/:accountId", getAccount)

  goji.Post("/$/accounts", postAccount)

  goji.Serve()
}