// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1/extauth/extauth.proto

package extauth

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *GatewayExtauth) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayExtauth)
	if !ok {
		that2, ok := that.(GatewayExtauth)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetExtauthzRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauthzRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauthzRef(), target.GetExtauthzRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHttpService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHttpService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHttpService(), target.GetHttpService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestTimeout(), target.GetRequestTimeout()) {
			return false
		}
	}

	if m.GetFailureModeAllow() != target.GetFailureModeAllow() {
		return false
	}

	if h, ok := interface{}(m.GetRequestBody()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequestBody()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequestBody(), target.GetRequestBody()) {
			return false
		}
	}

	if m.GetClearRouteCache() != target.GetClearRouteCache() {
		return false
	}

	if m.GetStatusOnError() != target.GetStatusOnError() {
		return false
	}

	if m.GetTransportApiVersion() != target.GetTransportApiVersion() {
		return false
	}

	if strings.Compare(m.GetStatPrefix(), target.GetStatPrefix()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *HttpService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpService)
	if !ok {
		that2, ok := that.(HttpService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetPathPrefix(), target.GetPathPrefix()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequest(), target.GetRequest()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *BufferSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BufferSettings)
	if !ok {
		that2, ok := that.(BufferSettings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetMaxRequestBytes() != target.GetMaxRequestBytes() {
		return false
	}

	if m.GetAllowPartialMessage() != target.GetAllowPartialMessage() {
		return false
	}

	if m.GetPackAsBytes() != target.GetPackAsBytes() {
		return false
	}

	return true
}

// Equal function
func (m *RouteExtauth) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteExtauth)
	if !ok {
		that2, ok := that.(RouteExtauth)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Spec.(type) {

	case *RouteExtauth_Disable:
		if _, ok := target.Spec.(*RouteExtauth_Disable); !ok {
			return false
		}

		if m.GetDisable() != target.GetDisable() {
			return false
		}

	case *RouteExtauth_ConfigRef:
		if _, ok := target.Spec.(*RouteExtauth_ConfigRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetConfigRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetConfigRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetConfigRef(), target.GetConfigRef()) {
				return false
			}
		}

	case *RouteExtauth_CustomAuth:
		if _, ok := target.Spec.(*RouteExtauth_CustomAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCustomAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCustomAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCustomAuth(), target.GetCustomAuth()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Spec != target.Spec {
			return false
		}
	}

	return true
}

// Equal function
func (m *CustomAuth) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CustomAuth)
	if !ok {
		that2, ok := that.(CustomAuth)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetContextExtensions()) != len(target.GetContextExtensions()) {
		return false
	}
	for k, v := range m.GetContextExtensions() {

		if strings.Compare(v, target.GetContextExtensions()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *HttpService_Request) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpService_Request)
	if !ok {
		that2, ok := that.(HttpService_Request)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetAllowedHeaders()) != len(target.GetAllowedHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowedHeaders() {

		if strings.Compare(v, target.GetAllowedHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetHeadersToAdd()) != len(target.GetHeadersToAdd()) {
		return false
	}
	for k, v := range m.GetHeadersToAdd() {

		if strings.Compare(v, target.GetHeadersToAdd()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *HttpService_Response) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpService_Response)
	if !ok {
		that2, ok := that.(HttpService_Response)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetAllowedUpstreamHeaders()) != len(target.GetAllowedUpstreamHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowedUpstreamHeaders() {

		if strings.Compare(v, target.GetAllowedUpstreamHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAllowedClientHeaders()) != len(target.GetAllowedClientHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowedClientHeaders() {

		if strings.Compare(v, target.GetAllowedClientHeaders()[idx]) != 0 {
			return false
		}

	}

	return true
}
