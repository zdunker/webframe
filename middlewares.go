package webframe

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(errMsg))
				c.ErrorResponse(http.StatusInternalServerError, errors.New("Internal Server Error"))
			}
		}()
		c.Next()
	}
}

func trace(msg string) string {
	pcs := make([]uintptr, 32)
	n := runtime.Callers(3, pcs)

	var builder strings.Builder
	builder.WriteString(msg)
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		builder.WriteString("\n\t")
		builder.WriteString(file)
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(line))
	}
	return builder.String()
}
