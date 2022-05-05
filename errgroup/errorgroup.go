package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func FetchUrmDemo() error {
	var g errgroup.Group
	var urls = []string{
		"http://pk.go.dev",
		"http://www.liwenzhou.com",
		"http://www.baiduxxx.com",
	}
	for _, url := range urls {
		g.Go(func() error {

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("获取%s成功", url)
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("所有的goroutine均成功")
	return nil
}

func main() {
	FetchUrmDemo()
}
