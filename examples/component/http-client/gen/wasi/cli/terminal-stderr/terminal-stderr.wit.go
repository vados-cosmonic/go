// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminalstderr represents the imported interface "wasi:cli/terminal-stderr@0.2.0".
package terminalstderr

import (
	terminaloutput "github.com/wasmCloud/go/examples/component/http-client/gen/wasi/cli/terminal-output"
	"go.bytecodealliance.org/cm"
)

// TerminalOutput represents the imported type alias "wasi:cli/terminal-stderr@0.2.0#terminal-output".
//
// See [terminaloutput.TerminalOutput] for more information.
type TerminalOutput = terminaloutput.TerminalOutput

// GetTerminalStderr represents the imported function "get-terminal-stderr".
//
//	get-terminal-stderr: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStderr() (result cm.Option[TerminalOutput]) {
	wasmimport_GetTerminalStderr(&result)
	return
}
