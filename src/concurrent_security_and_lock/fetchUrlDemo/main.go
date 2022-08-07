package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
)

func fetchUrlDemo() error {

	var g = new(errgroup.Group)
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}

	for _, url := range urls {
		url := url

		g.Go(func() error {

			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("GET error: %v", err)

			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("status error: %v", resp.StatusCode)
			}

			data, _ := io.ReadAll(resp.Body)
			fmt.Println(string(data))
			fmt.Printf("获取%s成功\n", url)

			return err
			//return err // 如何将错误返回呢？
		})

	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}

	return nil
	// 如何获取goroutine中可能出现的错误呢？
}

func ExampleGroupJustErrors() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

func main() {
	fetchUrlDemo()
}
