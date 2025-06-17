// Copyright (c) 2025 SDSLabs
// SPDX-License-Identifier: MIT

package log

import (
	"errors"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const (
	logDir  = "logs"
	logFile = "./" + logDir + "/accounts.log"
)

func initLogger(isProd bool) *zerolog.Logger {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(logDir, os.ModePerm)
		if mkdirErr != nil {
			panic(errors.New("unable to create log directory: " + mkdirErr.Error()))
		}
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(errors.New("unable to open log file: " + err.Error()))
	}

	var l zerolog.Logger
	if isProd {
		l = zerolog.New(file).With().Timestamp().Logger()
	} else {
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Logger()
	}

	return &l
}

func LoggerMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.With().Logger()

		t1 := time.Now()

		// Process request
		defer func() {
			t2 := time.Now()

			// Recover and record stack traces in case of a panic
			if rec := recover(); rec != nil {
				log.Error().
					Str("type", "error").
					Timestamp().
					Interface("recover_info", rec).
					Bytes("debug_stack", debug.Stack()).
					Msg("log system error")
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			// log end request
			log.Info().
				Str("type", "access").
				Timestamp().
				Fields(map[string]interface{}{
					"remote_ip":  c.ClientIP(),
					"url":        c.Request.URL.Path,
					"proto":      c.Request.Proto,
					"method":     c.Request.Method,
					"user_agent": c.Request.Header.Get("User-Agent"),
					"status":     c.Writer.Status(),
					"latency_ms": float64(t2.Sub(t1).Nanoseconds()) / 1000000.0,
					// "bytes_in":   c.Request.Header.Get("Content-Length"),
					// "bytes_out":  c.Writer.Size(),
				}).
				Msg("incoming_request")
		}()

		c.Next()
	}
}

var Logger = initLogger(false)
