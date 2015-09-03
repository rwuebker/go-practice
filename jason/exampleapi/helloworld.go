package main

import "net/http"
import "encoding/json"
import "fmt"

type name struct {
	first, last string
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:8080", nil)
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJSONResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "hello, %s!\n", string(response))
}

/*func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello,"))
}*/

func getJSONResponse() ([]byte, error) {
	names := make(map[string]string)
	names["Jason"] = "Kautz"

	return json.Marshal(names)
}
