package main
import(
	"fmt"
	"net/http"
	"time"
	"sync"
)
type Result struct{
	URL string
	StatusCode int
	Err error
}
func main(){
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/status/404",
		"https://httpbin.org/delay/1",
	}
	var wg sync.WaitGroup
	results := make(chan Result)
	for _, url := range urls {
		wg.Add(1)
		go ping(url, results, &wg)
	}
	go func(){
		wg.Wait()
		close(results)
	}()
	fmt.Println("--- Pring Results ---")
	for res := range results {
		if res.Err != nil {
			fmt.Printf("[ERROR] %-30s -> %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("[SUCCESS] %-28s -> Status: %d\n", res.URL, res.StatusCode)
		}
	}
	fmt.Println("All target URLs processed.")
}
func ping(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{
		Timeout: 5*time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		results <- Result{URL: url, StatusCode: 0, Err: err}
	}
	defer resp.Body.Close()
	results <-Result{URL: url, StatusCode: resp.StatusCode, Err: nil}
}