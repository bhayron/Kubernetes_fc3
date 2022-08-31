package main

//v1
import "net/http"

func main() {
	http.HandleFunc("/", Helo)
	http.ListenAndServe(":88", nil)
}

func Helo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}
