package pages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func View(w http.ResponseWriter, r *http.Request) {
	// read file dari people.json
	data, err := ioutil.ReadFile("people.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// jadikan variable
	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// loop variablenya, buat ngisi data
	var cards string
	for _, p := range people {
		card := fmt.Sprintf(`
		<div style="width: 400px; margin: 5 auto; padding: 20px; border-radius: 8px; box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.3);">
			<p><strong>Name:</strong> %s</p>
			<p><strong>ID:</strong> %s</p>
		</div>
		`, p.Name, p.ID)
		cards += card
	}

	fmt.Fprintf(w, `
	<html>
	<head>
	<style>
	body {
		display: flex;
		flex-direction: column;
		align-items: center;
		font-family: 'Arial', sans-serif;
		background-color: #f2f2f2;
	}
	</style>
	</head>
	<body>
	<h1>People:</h1>
	%s
	</body>
	</html>
	`, cards)
}
