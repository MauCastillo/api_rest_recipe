package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strconv"

	_ "github.com/lib/pq"
)

//Recipe kitchen structure
type Recipe struct {
	ID          int64  `json:"Id"`
	Title       string `json:"Title"`
	Ingredients string `json:"Ingredients"`
	Preparation string `json:"Preparation"`
	Updated     string `json:"Updated"`
}

//Recipes is a array of structure
type Recipes []Recipe

func returnAllRecipe(w http.ResponseWriter, r *http.Request) {
	var arrayRecipes Recipes
	// Connect to the "company_db" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	// Select Statement.
	var sql = "SELECT id, title, ingredients, preparation, updated_at FROM recipe_object;"
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var title string
		var ingredients string
		var preparation string
		var updatedAt string
		if err := rows.Scan(&id, &title, &ingredients, &preparation, &updatedAt); err != nil {
			log.Println(err)
		}
		recipe := Recipe{ID: id, Title: title, Ingredients: ingredients, Preparation: preparation, Updated: updatedAt}
		arrayRecipes = append(arrayRecipes, recipe)
	}

	log.Println("Endpoint Hit: returnAllRecipe")
	json.NewEncoder(w).Encode(arrayRecipes)
}

func insertRecipe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var objectRepice Recipe
	err := decoder.Decode(&objectRepice)

	if err != nil {
		log.Println(err)
	}

	log.Println(objectRepice.Title)
	// Connect to the "company_db" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
	if err != nil {
		log.Println("error connecting to the database: ", err)
	}

	var sql = "INSERT INTO recipe_object (title, ingredients, preparation, updated_at) VALUES "
	sql += "('" + objectRepice.Title + "', '" + objectRepice.Ingredients + "', '" + objectRepice.Preparation + "', NOW());"

	if _, err := db.Exec(sql); err != nil {
		log.Println(err)
	}

	log.Println("Endpoint Hit: insertRecipe")

	json.NewEncoder(w).Encode(objectRepice)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Recipe
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Connect to the "company_db" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
	if err != nil {
		log.Println("error connecting to the database: ", err)
	}
	// Insert a row into the "recipe" table.
	var sql = "DELETE FROM recipe_object WHERE id = " + strconv.FormatInt(msg.ID,10) + "; "
	if _, err := db.Exec(sql); err != nil {
		log.Println(err)
	}

	log.Println("Endpoint Hit: deleteRecipe")

	json.NewEncoder(w).Encode(msg)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allRepice", returnAllRecipe)
	http.HandleFunc("/insertRecipe", insertRecipe)
	http.HandleFunc("/deleteRecipe", deleteRecipe)

}

func main() {
	start := time.Now()
	address := "0.0.0.0:8000"
	handleRequests()
	log.Println("Starting server on address", address, " Date: ", start.Format(time.UnixDate))
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
