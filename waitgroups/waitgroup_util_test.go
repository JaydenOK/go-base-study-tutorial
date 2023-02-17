package waitgroups

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWaitGroup(t *testing.T) {

	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	wg := NewWaitGroup(3)

	for _, url := range urls {
		wg.BlockAdd()
		go func(url string) {
			defer wg.Done()
			fmt.Printf("%s: checking\n", url)
			res, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error: %v", err)
			} else {
				defer res.Body.Close()
				fmt.Printf("%s: result: %v\n", url, err)
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("Finished")

}
