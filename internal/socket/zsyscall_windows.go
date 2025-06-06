//go:build windows

// Code generated by 'go generate' using "github.com/observeinc/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package socket

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	return e
}

var (
	modws2_32 = windows.NewLazySystemDLL("ws2_32.dll")

	procbind        = modws2_32.NewProc("bind")
	procgetpeername = modws2_32.NewProc("getpeername")
	procgetsockname = modws2_32.NewProc("getsockname")
)

func bind(s windows.Handle, name unsafe.Pointer, namelen int32) (err error) {
	r1, _, e1 := syscall.SyscallN(procbind.Addr(), uintptr(s), uintptr(name), uintptr(namelen))
	if r1 == socketError {
		err = errnoErr(e1)
	}
	return
}

func getpeername(s windows.Handle, name unsafe.Pointer, namelen *int32) (err error) {
	r1, _, e1 := syscall.SyscallN(procgetpeername.Addr(), uintptr(s), uintptr(name), uintptr(unsafe.Pointer(namelen)))
	if r1 == socketError {
		err = errnoErr(e1)
	}
	return
}

func getsockname(s windows.Handle, name unsafe.Pointer, namelen *int32) (err error) {
	r1, _, e1 := syscall.SyscallN(procgetsockname.Addr(), uintptr(s), uintptr(name), uintptr(unsafe.Pointer(namelen)))
	if r1 == socketError {
		err = errnoErr(e1)
	}
	return
}
