// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package client

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"net/textproto"
	"strings"
	"sync"
)

type clientTracer struct {
	context.Context
	request     *http.Request
	requestBody []byte
	headers     map[string]interface{}
	mtx         sync.Mutex
}

func (ct *clientTracer) getConn(host string) {

}

func (ct *clientTracer) gotConn(info httptrace.GotConnInfo) {}

func (ct *clientTracer) putIdleConn(err error) {
}

func (ct *clientTracer) dnsStart(info httptrace.DNSStartInfo) {}

func (ct *clientTracer) dnsDone(info httptrace.DNSDoneInfo) {
	var buffer strings.Builder
	for _, v := range info.Addrs {
		if buffer.Len() != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(v.String())
	}
}

func (ct *clientTracer) connectStart(network, addr string) {}

func (ct *clientTracer) connectDone(network, addr string, err error) {
}

func (ct *clientTracer) tlsHandshakeStart() {

}

func (ct *clientTracer) tlsHandshakeDone(_ tls.ConnectionState, err error) {
}

func (ct *clientTracer) wroteHeaderField(k string, v []string) {
	if len(v) > 1 {
		ct.headers[k] = v
	} else {
		ct.headers[k] = v[0]
	}
}

func (ct *clientTracer) wroteHeaders() {

}

func (ct *clientTracer) wroteRequest(info httptrace.WroteRequestInfo) {}

func (ct *clientTracer) got100Continue() {

}

func (ct *clientTracer) wait100Continue() {

}

func (ct *clientTracer) gotFirstResponseByte() {

}

func (ct *clientTracer) got1xxResponse(code int, header textproto.MIMEHeader) error {
	return nil
}
