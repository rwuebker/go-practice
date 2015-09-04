package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestOurFunc(t *testing.T) {
	expected := "whats up"
	actual := funcToTest()
	if actual != expected {
		t.Error("Test failed")
	}

}

func TestHelloworld(t *testing.T) {

	person := `{"FirstName":"Jack", "LastName":"Black"}`

	usersUrl := "http://localhost:8080/helloWorld"

	reader := strings.NewReader(person)

	request, err := http.NewRequest("POST", usersUrl, reader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error("Bad Request expected: %d", res.StatusCode)
	} else {
		fmt.Printf("this is response: %+v\n", res.Body)
	}

	decoder := json.NewDecoder(res.Body)
	defer r.Body.Close()

	u := new(Person)
	if err := decoder.Decode(u); err != nil {
		t.Errorf("error occurred %s", err)
	}

	if u.FirstName != "Jack" {
		t.Errorf("expected Jack, got %s", u.FirstName)
		fmt.Printf("%s - %d ", u.FirstName, 10)
	}
}
