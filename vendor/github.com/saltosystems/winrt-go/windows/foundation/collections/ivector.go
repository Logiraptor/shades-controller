// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint
package collections

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const GUIDIVector string = "913337e9-11a1-4345-a3a2-4e7f956e222d"
const SignatureIVector string = "{913337e9-11a1-4345-a3a2-4e7f956e222d}"

type IVector struct {
	ole.IInspectable
}

type IVectorVtbl struct {
	ole.IInspectableVtbl

	GetAt       uintptr
	GetSize     uintptr
	GetView     uintptr
	IndexOf     uintptr
	SetAt       uintptr
	InsertAt    uintptr
	RemoveAt    uintptr
	Append      uintptr
	RemoveAtEnd uintptr
	Clear       uintptr
	GetMany     uintptr
	ReplaceAll  uintptr
}

func (v *IVector) VTable() *IVectorVtbl {
	return (*IVectorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IVector) GetAt(index uint32) (unsafe.Pointer, error) {
	var out unsafe.Pointer
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetAt,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(index),                // in uint32
		uintptr(unsafe.Pointer(&out)), // out unsafe.Pointer
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IVector) GetSize() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetSize,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *IVector) GetView() (*IVectorView, error) {
	var out *IVectorView
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetView,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out IVectorView
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IVector) IndexOf(value unsafe.Pointer) (uint32, bool, error) {
	var index uint32
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().IndexOf,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in unsafe.Pointer
		uintptr(unsafe.Pointer(&index)), // out uint32
		uintptr(unsafe.Pointer(&out)),   // out bool
	)

	if hr != 0 {
		return 0, false, ole.NewError(hr)
	}

	return index, out, nil
}

func (v *IVector) SetAt(index uint32, value unsafe.Pointer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetAt,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(index),                  // in uint32
		uintptr(unsafe.Pointer(&value)), // in unsafe.Pointer
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IVector) InsertAt(index uint32, value unsafe.Pointer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().InsertAt,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(index),                  // in uint32
		uintptr(unsafe.Pointer(&value)), // in unsafe.Pointer
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IVector) RemoveAt(index uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveAt,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(index),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IVector) Append(value unsafe.Pointer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().Append,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in unsafe.Pointer
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IVector) RemoveAtEnd() error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveAtEnd,
		uintptr(unsafe.Pointer(v)), // this
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IVector) Clear() error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().Clear,
		uintptr(unsafe.Pointer(v)), // this
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IVector) GetMany(startIndex uint32, itemsSize uint32) ([]unsafe.Pointer, uint32, error) {
	var items []unsafe.Pointer = make([]unsafe.Pointer, itemsSize)
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMany,
		uintptr(unsafe.Pointer(v)),         // this
		uintptr(startIndex),                // in uint32
		uintptr(itemsSize),                 // in uint32
		uintptr(unsafe.Pointer(&items[0])), // out unsafe.Pointer
		uintptr(unsafe.Pointer(&out)),      // out uint32
	)

	if hr != 0 {
		return nil, 0, ole.NewError(hr)
	}

	return items, out, nil
}

func (v *IVector) ReplaceAll(itemsSize uint32, items []unsafe.Pointer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().ReplaceAll,
		uintptr(unsafe.Pointer(v)),         // this
		uintptr(itemsSize),                 // in uint32
		uintptr(unsafe.Pointer(&items[0])), // in unsafe.Pointer
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
