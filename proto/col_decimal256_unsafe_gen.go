//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"reflect"
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes Decimal256 rows from *Reader.
func (c *ColDecimal256) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]Decimal256, rows)...)
	s := *(*reflect.SliceHeader)(unsafe.Pointer(c))
	s.Len *= 32
	s.Cap *= 32
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}
