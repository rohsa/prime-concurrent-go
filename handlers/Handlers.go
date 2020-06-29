package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/rohsa/prime-concurrent-go/rendoring"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var viewTemplate = filepath.Join("view")

var primeNumbers []int

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	rendoring.RenderTemplate(w, viewTemplate, nil)
}

func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	if r.Method == http.MethodGet {
		input := r.URL.Path[len("/prime/"):]
		number, err := strconv.Atoi(input)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusUnprocessableEntity)
			return
		}
		if number == 0 {
			http.Error(w, "Number should be greater than 0", http.StatusUnprocessableEntity)
			return
		}

		// Concurrent approach
		numberToEvaluate := make(chan int, 10)
		go addNumbersToEvaluate(numberToEvaluate, number)

		numberOfWorkers := 10
		var wg sync.WaitGroup
		var m sync.Mutex
		for i := 0; i < numberOfWorkers; i ++ {
			wg.Add(1)
			go addPrimeNumbers(numberToEvaluate, &m, &wg)
		}
		wg.Wait()

		// Synchronous approach

		//for i := 2; i < number ; i ++ {
		//	if isPrime(i) {
		//		primeNumbers = append(primeNumbers, i)
		//	}
		//}

		responseBytes, err := json.Marshal(primeNumbers)
		if err != nil {
			http.Error(w, "error writing data", http.StatusInternalServerError)
		}

		_, err = fmt.Fprintf(w, "%s\n", responseBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		primeNumbers = primeNumbers[:0]
		fmt.Printf("%s\n", time.Since(currentTime))
	}
}

//func isPrime(n int) bool {
//	is := true
//	for i := 2; i <= int(math.Floor(float64(n) / 2)); i++ {
//		if n%i == 0 {
//			is = false
//			break
//		}
//	}
//	return is
//}

func addNumbersToEvaluate(ch chan int, input int)  {
	for i := 2; i <= input; i++ {
		ch <- i
	}
	close(ch)
}

func addPrimeNumbers(ch chan int, m *sync.Mutex, wg *sync.WaitGroup) {
	for value := range ch {
		isPrime := true
		for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
			if value%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			m.Lock()
			primeNumbers = append(primeNumbers, value)
			m.Unlock()
		}
	}
	wg.Done()
}