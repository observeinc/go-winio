//go:build windows

// Code generated by 'go generate' using "github.com/observeinc/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package bindfilter

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
	modbindfltapi = windows.NewLazySystemDLL("bindfltapi.dll")

	procBfGetMappings   = modbindfltapi.NewProc("BfGetMappings")
	procBfRemoveMapping = modbindfltapi.NewProc("BfRemoveMapping")
	procBfSetupFilter   = modbindfltapi.NewProc("BfSetupFilter")
)

func bfGetMappings(flags uint32, jobHandle windows.Handle, virtRootPath *uint16, sid *windows.SID, bufferSize *uint32, outBuffer *byte) (hr error) {
	hr = procBfGetMappings.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procBfGetMappings.Addr(), uintptr(flags), uintptr(jobHandle), uintptr(unsafe.Pointer(virtRootPath)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(bufferSize)), uintptr(unsafe.Pointer(outBuffer)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func bfRemoveMapping(jobHandle windows.Handle, virtRootPath string) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(virtRootPath)
	if hr != nil {
		return
	}
	return _bfRemoveMapping(jobHandle, _p0)
}

func _bfRemoveMapping(jobHandle windows.Handle, virtRootPath *uint16) (hr error) {
	hr = procBfRemoveMapping.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procBfRemoveMapping.Addr(), uintptr(jobHandle), uintptr(unsafe.Pointer(virtRootPath)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func bfSetupFilter(jobHandle windows.Handle, flags uint32, virtRootPath string, virtTargetPath string, virtExceptions **uint16, virtExceptionPathCount uint32) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(virtRootPath)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(virtTargetPath)
	if hr != nil {
		return
	}
	return _bfSetupFilter(jobHandle, flags, _p0, _p1, virtExceptions, virtExceptionPathCount)
}

func _bfSetupFilter(jobHandle windows.Handle, flags uint32, virtRootPath *uint16, virtTargetPath *uint16, virtExceptions **uint16, virtExceptionPathCount uint32) (hr error) {
	hr = procBfSetupFilter.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procBfSetupFilter.Addr(), uintptr(jobHandle), uintptr(flags), uintptr(unsafe.Pointer(virtRootPath)), uintptr(unsafe.Pointer(virtTargetPath)), uintptr(unsafe.Pointer(virtExceptions)), uintptr(virtExceptionPathCount))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}
