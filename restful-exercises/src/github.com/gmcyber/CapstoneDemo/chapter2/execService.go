package main
import (
  "bytes"
  "fmt"
  "log"
  "net/http"
  "os/exec"
  "github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string) string {
  cmd := exec.Command(command, arguments...)
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err := cmd.Start()

  if err != nil {
    log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
  }
  err = cmd.Wait()
  if err != nil {
    log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
  }
  return out.String()

}
func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  fmt.Fprintf(w, getCommandOutput("/usr/local/go/bin/go", "version"))
}

func getFileEcontent(w http.ResponseWriter, r * http.Request, params httprouter.Params) {
  fmt.Fprintf(w, getCommandOutput("/usr/bin/cat", params.ByName("name")))
}

func main() {
  router := httprouter.New()
  router.GET("/api/v1/go-version", goVersion)
  router.GET("/api/v1/show-file/:name", getFileEcontent)
  log.Fatal(http.ListenAndServe(":8000", router))
}
