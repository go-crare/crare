package middleware

import (
	"github.com/3JoB/unsafeConvert"

	"gopkg.in/crare.v1"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger(logger crare.Logger) crare.HandlerFunc {
	return func(c crare.Context) error {
		data, _ := c.Bot().Json().MarshalIndent(c.Update(), "", "  ")
		logger.Println(unsafeConvert.StringSlice(data))
		return c.Next()
	}
}
