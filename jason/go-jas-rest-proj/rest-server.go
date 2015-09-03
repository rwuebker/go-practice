package main
 
import (
    "fmt"
    "net/http"
    "github.com/coocood/jas"
)
 
type Hello struct {}
 
func (*Hello) Get (ctx *jas.Context) {
    ctx.Data = "Hello World"
}
 
func main () {
    router := jas.NewRouter(new(Hello))
    router.BasePath = "/v1/"
    fmt.Println(router.HandledPaths(true))
    http.Handle(router.BasePath, router)
    http.ListenAndServe(":8080", nil)
}

