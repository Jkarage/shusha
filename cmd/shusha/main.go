package main

import (
	"fmt"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/jkarage/vuta"
)

func run() error {
	var cfg = struct {
		Version string `conf:"default:0.0.1"`
		URL     string `conf:"flag,required"`
		Threads int    `conf:"default:1"`
	}{}

	_, err := conf.Parse("", &cfg)
	if err != nil {
		return err
	}

	client := vuta.NewDownload(cfg.URL)
	client.Download()

	// cLength, name, _, err := client.Header()
	// if err != nil {
	// 	return err
	// }

	// fmt.Println("CONTENT_LENGTH: ", cLength)

	// threads := runtime.NumCPU()
	// size := cLength / int64(threads)

	// fmt.Println("THREADS: ", threads)
	// fmt.Println("SIZE: ", size)

	// start := 0
	// end := size - 1
	// var wg sync.WaitGroup

	// wg.Add(runtime.NumCPU())

	// for i := 0; i < threads-1; i++ {
	// 	fmt.Println("send goroutine ", i, start, end)
	// 	partName := fmt.Sprintf(name+"%v", i+1)
	// 	go client.DownloadChunks(partName, &wg, start, int(end))
	// 	start = int(end) + 1
	// 	end = int64(start) + (size - 1)
	// }

	// end = cLength - 1
	// fmt.Println("sending the last goroutine", start, end)
	// partName := fmt.Sprintf(name+"%v", threads)
	// go client.DownloadChunks(partName, &wg, start, int(end))

	// // defer file.Close()
	// wg.Wait()

	return nil
}

func main() {
	now := time.Now()
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Time taken: ", time.Since(now).Seconds())

}
