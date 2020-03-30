package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	usage := "Usage: allports [start] [end]"
	if len(os.Args) != 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("[start] must be a number")
		fmt.Println(usage)
		os.Exit(1)
	}
	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("[end] must be a number")
		fmt.Println(usage)
		os.Exit(1)
	}
	ports := []int{}
	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}

	mkServers(ports)
}

func mkServers(ports []int) {
	wg := sync.WaitGroup{}
	for _, port := range ports {
		wg.Add(1)
		go listen(port, &wg)
	}
	wg.Wait()
}

func listen(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	handler := mkHandler(port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		fmt.Println(err)
	}
}

func mkHandler(port int) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", port)
	}
	return f
}
