package main
import (
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
)

func main() {
  router := httprouter.New()
  router.ServeFiles("/static/*filepath",
    http.Dir("/home/champuser/CapstoneDemo/restful-exercises/src/github.com/gmcyber/CapstoneDemo/chapter2/static"))
  log.Fatal(http.ListenAndServe(":8000", router))
}
