package main

// некоторые импорты нужны для проверки
import (
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

func sum(a, b int) int {
	return a + b*2
}

func main() {

}
