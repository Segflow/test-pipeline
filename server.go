package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

func main() {

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		m, _ := mem.VirtualMemory()
		free := fmt.Sprintf("%v", m.Available)
		w.Write([]byte(free))
	})

	http.HandleFunc("/fact", func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("n")

		n, _ := strconv.Atoi(s)
		res := factRecursive(n)

		fmt.Fprintf(w, "result: %d", res)
	})

	http.ListenAndServe("0.0.0.0:8888", nil)
}
