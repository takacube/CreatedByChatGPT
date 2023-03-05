package main

import (
    "net/http"
    "fmt"
	"encoding/json"
)

type loginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", loginHandler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	isValid := validateUser("exampleName", "examplePassword")
	if isValid {
		fmt.Println("User is valid")
	} else {
    fmt.Println("User is invalid")
	}
    fmt.Fprint(w, "Hello, World!")
}



func validateUser(name, password string) bool {
    // 適当に決めた文字列
    correctName := "exampleName"
    correctPassword := "examplePassword"

    if name == correctName && password == correctPassword {
        return true
    }
    return false
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// JSON形式のリクエストボディをパースする
	var reqBody loginRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// パースしたリクエストボディをそのままレスポンスとして返す
	respBody, err := json.Marshal(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}