package main
import (
  "fmt"
  "net/http"
)

func middleware(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
    fmt.Println("Middleware Before Request")
    handler.ServeHTTP(w, r)
    fmt.Println("Middleware After Request")

  })
}
func mainLogic(w http.ResponseWriter, r *http.Request){
  fmt.Println("mainLogicHandler")
  w.Write([]byte("OK"))
}
func main(){
  mainLogicHandler := http.HandlerFunc(mainLogic)
  http.Handle("/", middleware(mainLogicHandler))
  http.ListenAndServe(":8000", nil)
}
