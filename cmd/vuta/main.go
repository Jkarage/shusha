package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/jkarage/vuta"
)

var build = "develop"

func run() error {
	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout time.Duration `conf:"default:5s,flag:read-timeout"`
			IdleTimeout time.Duration `conf:"default:120s,flag:idle-timeout"`
			Verbose     bool          `conf:"default:false,flag:v"`
			URL         string        `conf:"required,flag:url"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "Vuta is a file downloader tool, utilizing concurrency in golang",
		},
	}

	help, err := conf.Parse("Vuta", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		return fmt.Errorf("%w", err)
	}

	client := vuta.NewApp(cfg.Web.URL)
	// parts := path.Base(cfg.Web.URL)

	client.Download()

	// cLength, _, _, err := vuta.Header(cfg.Web.URL)
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

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}
}
