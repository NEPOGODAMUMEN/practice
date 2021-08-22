package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

//var out string = fmt.Sprintf("Number: \t %07d", 45)
//fmt.Println(out)

/*fmt.Println("Type the year you were born: ")
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
fmt.Printf("You will be %d years old at he end of 2021", 2021-input)*/
