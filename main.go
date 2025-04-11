package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
	"strconv"
)

var c int

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte(strconv.Itoa(c)))
	case "POST":
		r.ParseForm()
		res := r.Form.Get("count")
		resi, err := strconv.Atoi(res)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
		} else {
			c += resi
		}
	}
}

func main() {
	http.HandleFunc("/count", handler)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
