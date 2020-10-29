package main // import "github.com/oiyoup/logger"

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// MyError ...
type MyError struct{}

func (r MyError) Error() string {
	return "hello"
}

var (
	logger = log.Logger.Level(zerolog.InfoLevel)
)

func init() {
	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).Level(zerolog.DebugLevel)
}

func main() {
	logger.Fatal().Err(new(MyError)).Msg("Fail")
}
