package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"math"
	"math/rand"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f := fib()

		res := &response{Message: "Hello World"}

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			res.EnvVars = append(res.EnvVars, pair[0]+"="+pair[1])
		}
		fmt.Println("Environment variables added to results")

		for i := 1; i <= 90; i++ {
			res.Fib = append(res.Fib, f())
		}
		fmt.Println("Fibonacci sequence calculated")

		
		res.MonteCarlo = monteCarloPi(100000000)
		fmt.Println("Monte Carlo Pi calculation completed")

		// Beautify the JSON output
		out, _ := json.MarshalIndent(res, "", "  ")

		// Normally this would be application/json, but we don't want to prompt downloads
		w.Header().Set("Content-Type", "text/plain")

		io.WriteString(w, string(out))

		fmt.Println("Result ready")
	})
	http.ListenAndServe(":8080", nil)
}

type response struct {
	Message string   `json:"message"`
	EnvVars []string `json:"env"`
	Fib     []int    `json:"fib"`
	MonteCarlo float64 `json:"pi"`
}


func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func inCircle(x, y float64) bool {
	return math.Sqrt(x*x+y*y) <= 1.0
}

func monteCarloPi(iterations int) float64 {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	var h int
	for i := 0; i <= iterations; i++ {
		if inCircle(r.Float64(), r.Float64()) {
			h++
		}
	}
	pi := 4 * float64(h) / float64(iterations)
	return pi
}
