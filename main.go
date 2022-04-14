package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// La structure Task

type Task struct {
	Description string
	Done        bool
	ID          string
}

// La variable globale tasks

var tasks = []Task{
	{
		Description: "Faire les courses",
		Done:        false,
		ID:          "0",
	},
	{
		Description: "Payer les factures",
		Done:        false,
		ID:          "1",
	},
}

// Début de ma fonction main

func main() {

	// Début de la fonction list

	list := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		for _, task := range tasks {
			if task.Done == false {
				data := ([]byte("ID " + task.ID + ", task: " + task.Description + "\n"))
				w.Write([]byte(data))
			}
		}
	}

	// Fin de la fonction list

	// Début de la fonction done

	done := func(w http.ResponseWriter, re *http.Request) {
		switch re.Method {
		case http.MethodGet:
			for _, task := range tasks {
				if task.Done == true {
					data := ([]byte("ID " + task.ID + ", task : " + task.Description + "\n"))
					w.Write([]byte(data))
				}
			}

		}

	}
	// Fin de la fonction done

	//Début de la fonction add

	add := func(w http.ResponseWriter, re *http.Request) {
		if re.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(re.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			http.Error(
				w,
				"can't read body", http.StatusBadRequest)
			return
		}
		desc := string(body)
		tasks = append(tasks, Task{desc, false, "2"})
		w.WriteHeader(http.StatusOK)
	}

	// Fin de la fonction add

	// Appel des 3 Handlefunc

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	//  Appel ListenAndServe

	http.ListenAndServe("localhost:8081", nil)
}

// Fin de la fonction main
