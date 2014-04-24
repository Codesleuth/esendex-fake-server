package main

import (
  "io"
  "fmt"
)

type Account struct {
  ID        string  `param:"id"`
  Reference string  `param:"reference"`
}

func (account Account) Write(w io.Writer) {
  fmt.Fprintf(w, "{\"id\":\"%s\",\"reference\":\"%s\"}", account.ID, account.Reference)
}