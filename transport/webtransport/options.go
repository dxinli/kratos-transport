package webtransport

import (
	"crypto/tls"
	"time"

	"github.com/go-kratos/kratos/v2/encoding"

	"github.com/quic-go/quic-go"
)

type ServerOption func(*Server)

func WithTLSConfig(c *tls.Config) ServerOption {
	return func(o *Server) {
		o.tlsConf = c
	}
}

func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.Addr = addr
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxIdleTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		if s.Server.QUICConfig == nil {
			s.Server.QUICConfig = &quic.Config{}
		}
		s.Server.QUICConfig.MaxIdleTimeout = timeout
	}
}

func WithKeepAlivePeriod(timeout time.Duration) ServerOption {
	return func(s *Server) {
		if s.Server.QUICConfig == nil {
			s.Server.QUICConfig = &quic.Config{}
		}
		s.Server.QUICConfig.KeepAlivePeriod = timeout
	}
}

func WithPath(path string) ServerOption {
	return func(s *Server) {
		s.path = path
	}
}

func WithConnectHandle(h ConnectHandler) ServerOption {
	return func(s *Server) {
		s.connectHandler = h
	}
}

func WithCodec(c string) ServerOption {
	return func(s *Server) {
		s.codec = encoding.GetCodec(c)
	}
}

////////////////////////////////////////////////////////////////////////////////

type ClientOption func(*Client)

func WithClientTLSConfig(c *tls.Config) ClientOption {
	return func(o *Client) {
		o.tlsConf = c
	}
}

func WithEndpoint(url string) ClientOption {
	return func(o *Client) {
		o.url = url
	}
}

func WithClientCodec(c string) ClientOption {
	return func(o *Client) {
		o.codec = encoding.GetCodec(c)
	}
}

func WithClientMaxIdleTimeout(timeout time.Duration) ClientOption {
	return func(s *Client) {
		if s.transport.QUICConfig == nil {
			s.transport.QUICConfig = &quic.Config{}
		}
		s.transport.QUICConfig.MaxIdleTimeout = timeout
	}
}

func WithClientKeepAlivePeriod(timeout time.Duration) ClientOption {
	return func(s *Client) {
		if s.transport.QUICConfig == nil {
			s.transport.QUICConfig = &quic.Config{}
		}
		s.transport.QUICConfig.KeepAlivePeriod = timeout
	}
}
