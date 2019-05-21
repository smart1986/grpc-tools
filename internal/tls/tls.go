package tls

import (
	"google.golang.org/grpc/metadata"
	"net/http"
	"regexp"
)

const (
	forwardedHeader = "Forwarded"
	httpsProto      = "proto=https"
)

var httpsProtoPattern = regexp.MustCompile(httpsProto)

func AddHTTPSMarker(header http.Header) {
	header.Add(forwardedHeader, httpsProto)
}

func IsTLSRPC(md metadata.MD) bool {
	values := md.Get(forwardedHeader)
	for _, value := range values {
		if value == httpsProto {
			return true
		}
	}

	return false
}

func IsTLSRequest(header http.Header) bool {
	value := header.Get(forwardedHeader)
	return httpsProtoPattern.MatchString(value)
}
