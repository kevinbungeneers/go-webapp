package server

import "log/slog"

type Option func(s *server)

func WithListenAddress(address string) Option {
	return func(s *server) {
		if len(address) == 0 {
			return
		}
		s.listenAddress = address
	}
}

func WithLogLevel(l slog.Level) Option {
	return func(s *server) {
		s.logLevel = l
	}
}
