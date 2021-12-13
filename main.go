package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/mert-erol/api_example/models"
)

func main() {
	apiRoot := "/api"

	//simple routing

	// list all data from table user
	http.HandleFunc(apiRoot+"/list", func(w http.ResponseWriter, r *http.Request) {
		db := conn()
		defer db.Close()
		response, err := db.Query("SELECT * FROM user")

		output, err := json.Marshal(response)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// Getting datas of specific user
	http.HandleFunc(apiRoot+"/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		db := conn()
		defer db.Close()

		path := r.URL.Path // Getting url path without gorilla

		u, _ := url.Parse(path) // parsing url
		id := getFirstParam(u.Path)

		response, err := db.Query("SELECT * FROM user WHERE user_id = " + id)

		output, err := json.Marshal(response)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	//creating new user
	http.HandleFunc(apiRoot+"/create_user", func(w http.ResponseWriter, r *http.Request) {
		db := conn()
		defer db.Close()

		sql := "INSERT INTO user(FirstName, LastName, Age) VALUES('Mert', 'Erol', '25')"
		result, err := db.Exec(sql)
		checkError(err)

		affected, err := result.RowsAffected()
		fmt.Printf("%d rows added.", affected)

	})

	//deleting user
	http.HandleFunc(apiRoot+"/delete_user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		db := conn()
		defer db.Close()

		path := r.URL.Path // Getting url path without gorilla

		u, _ := url.Parse(path) // parsing url
		id := getFirstParam(u.Path)

		sql := "DELETE FROM cities WHERE id = " + id
		result, err := db.Exec(sql)
		checkError(err)
		affectedRows, err := result.RowsAffected()
		checkError(err)
		fmt.Printf("%d kayıt silindi.", affectedRows)

	})

}

func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:111@tcp(127.0.0.1:3306)/api_example")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// I use this function for parse the url and get the id without gorilla
func getFirstParam(path string) (ps string) {

	// ignore first '/' and when it hits the second '/'
	// get whatever is after it as a parameter
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			ps = path[i+1:]
		}
	}
	return
}
