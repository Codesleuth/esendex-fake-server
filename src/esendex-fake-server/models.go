package main

import (
  "io"
  "fmt"
  "encoding/json"
)

type Account struct {
  ID        string  `json:"id"`
  Reference string  `json:"reference"`
}

func (account Account) Write(w io.Writer) {
  resultJson, _ := json.Marshal(account)
  fmt.Fprint(w, string(resultJson))
}