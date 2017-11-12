package controller

import(
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

type User struct {
  Login string
  Id int
}

type Pulls struct {
  Id int
  RequestedReviewers []User
}

type PullsResponse struct {
  Collection []Pulls
}

func RepositoryControlLoop () {
  for true {
    CheckForNewPR()
  }
}

func CheckForNewPR () {
  fmt.Println("Checking repo for new Pulls")

  res, err := http.Get("https://api.github.com/repos/llvieira/GHCLi/pulls")
  if err != nil {
    fmt.Println("An error occured: ", err)
    return
  }

  defer res.Body.Close()

  ReadBody(res)
}

func ReadBody (res *http.Response) {
  if res.StatusCode == http.StatusOK {
    bodyBytes, err2 := ioutil.ReadAll(res.Body)
    if err2 != nil {
      fmt.Println("An error occured: ", err2)
      return
    }

    pulls := make([]Pulls, 0)
    fmt.Println(&pulls)

    json.Unmarshal(bodyBytes, &pulls)

    fmt.Println(pulls)
  }
}

