package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jkarage/vuta/internal/vuta"
)

var (
	// build = "develop"
	url = flag.String("url", "", "URL of the file to download")
)

func run() error {
	flag.Parse()

	now := time.Now()

	client := vuta.NewDownload(*url)
	client.Download()

	// cLength, _, _, err := vuta.Header()
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
	// 	partName := fmt.Sprintf(parts+"%v", i+1)
	// 	go client.DownloadChunks(partName, &wg, start, int(end))
	// 	start = int(end) + 1
	// 	end = int64(start) + (size - 1)
	// }

	// end = cLength - 1
	// fmt.Println("sending the last goroutine", start, end)
	// partName := fmt.Sprintf(parts+"%v", threads)
	// go client.DownloadChunks(partName, &wg, start, int(end))

	// // defer file.Close()
	// wg.Wait()

	fmt.Println(time.Since(now))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}

}
