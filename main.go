package main

/*
This project is created by mauricio castillo
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

//Recipe kitchen structure
type Recipe struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Preparation string `json:"preparation"`
	Updated     string `json:"updated"`
}

//Recipes is a array of structure
type Recipes []Recipe

// returnAllRecipe this function return a array all Recipe
func returnAllRecipe(w *http.ResponseWriter, r *http.Request) {

	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//enableCors(&w)
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
		var id string
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

// findRecipe this function return a array all Recipe in based
/*
	{
		"name": "desayuno"
	}
*/
func findRecipe(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	type Message struct {
		Name string `json:"name"`
	}
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var arrayRecipes Recipes
	// Connect to the "company_db" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
	if err != nil {
		log.Println(err)
	}

	// Select Statement.
	var sql = "SELECT id, title, ingredients, preparation, updated_at FROM recipe_object "
	sql += "WHERE LOWER(title) LIKE LOWER('%" + msg.Name + "%') OR LOWER(ingredients) LIKE LOWER('%" + msg.Name + "%'); "
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
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

	log.Println("Endpoint Hit: findRecipe")
	json.NewEncoder(w).Encode(arrayRecipes)
}

// insertRecipe this function insert a new Recipe in based
/*
{
	"Title":"Arroz llanero",
	"Ingredients":"Rancheras, Huevos,Arroz, leche",
	"Preparation":"Dile a ungringo que llame a su empleada latina y lo prepare"
}
*/
func insertRecipe(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	decoder := json.NewDecoder(r.Body)
	var objectRepice Recipe
	err := decoder.Decode(&objectRepice)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

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

/*This function delete a object type recipe of database
{
	"Id": 448890587979382785,
}
*/
func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} /*This function delete a object type recipe of database
	{
		"Id": 448890587979382785,
	}
	*/

	// Unmarshal
	var msg Recipe
	
	log.Println(">>> ")
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
	var sql = "DELETE FROM recipe_object WHERE id = " + msg.ID + "; "
	log.Println(">>> sql <<< ", sql)
	if _, err := db.Exec(sql); err != nil {
		log.Println(err)
	}

	log.Println("Endpoint Hit: deleteRecipe")

	json.NewEncoder(w).Encode(msg)
}

/*This function update a object type recipe of database
{
    "Id": 448904999043235841,
    "Title":"Desayuno para campeones",
	"Ingredients":"Un capeon y un pancahco",
	"Preparation":"Dale de comer al campeon"
}
*/
func updateRecipe(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	var sql = "UPDATE recipe_object "
	sql += "SET title = '" + msg.Title + "', ingredients = '" + msg.Ingredients + "', preparation = '" + msg.Preparation + "', updated_at = NOW() "
	sql += "WHERE id = " + msg.ID + "; "

	if _, err := db.Exec(sql); err != nil {
		log.Println(err)
	}

	log.Println("Endpoint Hit: updateRecipe")

	json.NewEncoder(w).Encode(msg)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allRepice", returnAllRecipe)
	http.HandleFunc("/insertRecipe", insertRecipe)
	http.HandleFunc("/deleteRecipe", deleteRecipe)
	http.HandleFunc("/updateRecipe", updateRecipe)
	http.HandleFunc("/findRecipe", findRecipe)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
