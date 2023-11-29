package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/jkarage/vuta"
)

var build = "v1.0.0"

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

		return fmt.Errorf("parsing config: %w", err)
	}

	client := vuta.NewApp(cfg.Web.URL)

	file, err := os.Create(path.Base(cfg.Web.URL))
	if err != nil {
		return err
	}

	client.Download(file)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}
}
