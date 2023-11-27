package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/jkarage/vuta/external/logger"
	"go.uber.org/zap"
)

var build = "develop"

func main() {
	log, err := logger.New("vuta")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer log.Sync()

	if err := run(log); err != nil {
		fmt.Println(err)
		return

	}
}

func run(log *zap.SugaredLogger) error {

	// -----------------------------------------------------------------------------
	// GOMAXPROCS
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	cfg := struct {
		conf.Version
		Vuta struct {
			ReadTimeout time.Duration `conf:"default:5s"`
			IdleTimeout time.Duration `conf:"default:120s"`
			URL         string        `conf:"required"`
			Destination string        `conf:"default:."`
			APIHost     string        `conf:"default:0.0.0.0:3000"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "Download file, utilizing a concurrency in golang",
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

	return nil

}
