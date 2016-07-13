package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		visitors int
		err      error
	)
	if buf, err := ioutil.ReadFile("./hwcounter"); err == nil {
		if n, err := strconv.Atoi(string(buf)); err == nil {
			visitors = n
		}
	}
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	visitors += 1
	fmt.Println("Handle", r.Method, r.URL)
	fmt.Fprintf(w, "Hello world from my Go program! You are %s visitor!\n", ordinal(visitors))
	if err := ioutil.WriteFile("./hwcounter", []byte(fmt.Sprintf("%d", visitors)), 0644); err != nil {
		fmt.Println("ERROR:", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
