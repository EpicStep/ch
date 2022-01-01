//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes DateTime rows from *Reader.
func (c *ColDateTime) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]DateTime, rows)...)
	s := *(*slice)(unsafe.Pointer(c))
	const size = 32 / 8
	s.Len *= size
	s.Cap *= size
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}

// EncodeColumn encodes DateTime rows to *Buffer.
func (c ColDateTime) EncodeColumn(b *Buffer) {
	if len(c) == 0 {
		return
	}
	offset := len(b.Buf)
	const size = 32 / 8
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	s := *(*slice)(unsafe.Pointer(&c))
	s.Len *= size
	s.Cap *= size
	src := *(*[]byte)(unsafe.Pointer(&s))
	dst := b.Buf[offset:]
	copy(dst, src)
}
