package pages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Person struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func InputJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		id := r.FormValue("id")

		// jika input di form tidak ada nama atau ID, check url
		if name == "" && id == "" {
			name = r.URL.Query().Get("Name")
			id = r.URL.Query().Get("ID")
		}

		p := Person{Name: name, ID: id}

		// jika ga ada json, buat baru. jika ada, lansung tambahkan
		var people []Person
		if _, err := os.Stat("people.json"); os.IsNotExist(err) { // kalau belum ada file jsonnya
			people = []Person{}
		} else {
			data, err := ioutil.ReadFile("people.json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := json.Unmarshal(data, &people); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// terima data dari POST
		people = append(people, p)
		newData, err := json.Marshal(people)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// rewrite json ke data baru
		if err := ioutil.WriteFile("people.json", newData, 0644); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Data saved!")
		return
	}

	fmt.Fprintln(w, `
	<html>
	<head>
	<style>
	body {
		display: flex;
		justify-content: center;
		align-items: center;
		font-family: 'Arial', sans-serif;
		background-color: #f2f2f2;
	}
	form {
		background-color: #ffffff;
		max-width: 400px;
		margin: 0 auto;
		padding: 20px;
		border-radius: 8px;
		box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.3);
	}
	input[type="text"] {
		width: 100%;
		padding: 10px;
		border: none;
		border-radius: 4px;
		margin: 6px 0;
		box-sizing: border-box;
		font-size: 16px;
	}
	input[type="submit"] {
		background-color: #4CAF50;
		color: white;
		font-size: 16px;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		padding: 12px 20px;
	}
	</style>
	</head>
	<body>
	<form method="POST">
	<label for="name">Name:</label><br>
	<input type="text" id="name" name="name" placeholder="Enter your name"><br>
	<label for="id">ID:</label><br>
	<input type="text" id="id" name="id" placeholder="Enter your ID"><br>
	<input type="submit" value="Submit">
	</form>
	</body>
	</html>
	`)
}
