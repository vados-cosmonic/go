// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package outgoinghandler represents the imported interface "wasi:http/outgoing-handler@0.2.0".
package outgoinghandler

import (
	"github.com/wasmCloud/go/examples/component/invoke/gen/wasi/http/types"
	"go.bytecodealliance.org/cm"
)

// OutgoingRequest represents the imported type alias "wasi:http/outgoing-handler@0.2.0#outgoing-request".
//
// See [types.OutgoingRequest] for more information.
type OutgoingRequest = types.OutgoingRequest

// RequestOptions represents the imported type alias "wasi:http/outgoing-handler@0.2.0#request-options".
//
// See [types.RequestOptions] for more information.
type RequestOptions = types.RequestOptions

// FutureIncomingResponse represents the imported type alias "wasi:http/outgoing-handler@0.2.0#future-incoming-response".
//
// See [types.FutureIncomingResponse] for more information.
type FutureIncomingResponse = types.FutureIncomingResponse

// ErrorCode represents the type alias "wasi:http/outgoing-handler@0.2.0#error-code".
//
// See [types.ErrorCode] for more information.
type ErrorCode = types.ErrorCode

// Handle represents the imported function "handle".
//
//	handle: func(request: outgoing-request, options: option<request-options>) -> result<future-incoming-response,
//	error-code>
//
//go:nosplit
func Handle(request OutgoingRequest, options cm.Option[RequestOptions]) (result cm.Result[ErrorCodeShape, FutureIncomingResponse, ErrorCode]) {
	request0 := cm.Reinterpret[uint32](request)
	options0, options1 := lower_OptionRequestOptions(options)
	wasmimport_Handle((uint32)(request0), (uint32)(options0), (uint32)(options1), &result)
	return
}
