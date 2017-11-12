package main

import(
  "./src/controller"
  "fmt"
)

func main () {
  fmt.Println("Starting server...")
  controller.RepositoryControlLoop()
}
