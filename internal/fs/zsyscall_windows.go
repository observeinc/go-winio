//go:build windows

// Code generated by 'go generate' using "github.com/observeinc/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package fs

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
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procCreateFileW = modkernel32.NewProc("CreateFileW")
)

func CreateFile(name string, access AccessMask, mode FileShareMode, sa *windows.SecurityAttributes, createmode FileCreationDisposition, attrs FileFlagOrAttribute, templatefile windows.Handle) (handle windows.Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _CreateFile(_p0, access, mode, sa, createmode, attrs, templatefile)
}

func _CreateFile(name *uint16, access AccessMask, mode FileShareMode, sa *windows.SecurityAttributes, createmode FileCreationDisposition, attrs FileFlagOrAttribute, templatefile windows.Handle) (handle windows.Handle, err error) {
	r0, _, e1 := syscall.SyscallN(procCreateFileW.Addr(), uintptr(unsafe.Pointer(name)), uintptr(access), uintptr(mode), uintptr(unsafe.Pointer(sa)), uintptr(createmode), uintptr(attrs), uintptr(templatefile))
	handle = windows.Handle(r0)
	if handle == windows.InvalidHandle {
		err = errnoErr(e1)
	}
	return
}
