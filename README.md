# prime-concurrent-go
This project implements a web UI which takes a number as inout from the user. Implements an ajax `GET` call which passes this number and gets all the prime numbers until the inoput number and displays it on the UI.

It implements the go wroker pools and the go routine to implement it concurrently.

# What it does
The project:
1. Takes a number as inout from the user
2. Makes a GET ajax call passing that number and receiving all the number prime numbers till that number as response
3. Implements the Golang/Go Programming as the server side language which finds all the prime numbers concurrently using the worker pool (using go wait groups) and go routine
3. Shows the response on the UI
# Running the project
1. Please make sure the project follows the correct path order (`github.com/rohsa/prime-concurrent-go`)
2. Please make sure you're not using/blocking the port `8080` on your localhost
2. `cd` to the project directory and run `go build main.go`
3. Open your web browser and hit the url `www.localhost:8080`
