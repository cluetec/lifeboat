package logging

import (
	"context"
	"crypto/tls"
	"github.com/rs/zerolog"
	"net/http/httptrace"
	"time"
)

func GetHttpTracingContext() (context.Context, *zerolog.Event) {
	var start, connect, dns, tlsHandshake time.Time
	traceLoggerEvent := zerolog.Dict()

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			traceLoggerEvent.Dur("dns", time.Since(dns))
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			traceLoggerEvent.Dur("tlsHandshake", time.Since(tlsHandshake))
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			traceLoggerEvent.Dur("connect", time.Since(connect))
		},

		GotFirstResponseByte: func() {
			traceLoggerEvent.Dur("receivedFirstByte:", time.Since(start))
		},
	}

	return httptrace.WithClientTrace(context.Background(), trace), traceLoggerEvent
}
