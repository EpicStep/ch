{{- /*gotype: github.com/go-faster/ch/proto/cmd/ch-gen-col.Variant*/ -}}
//go:build amd64 && !nounsafe
// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes {{ .Name }} rows from *Reader.
func (c *{{ .Type }}) DecodeColumn(r *Reader, rows int) error {
  if rows == 0 {
	return nil
  }
  *c = append(*c, make([]{{ .ElemType }}, rows)...)
  s := *(*slice)(unsafe.Pointer(c))
  {{- if not .SingleByte }}
  const size = {{ .Bits }} / 8
  s.Len *= size
  s.Cap *= size
  {{- end }}
  dst := *(*[]byte)(unsafe.Pointer(&s))
  if err := r.ReadFull(dst); err != nil {
  	return errors.Wrap(err, "read full")
  }
  return nil
}

// EncodeColumn encodes {{ .Name }} rows to *Buffer.
func (c {{ .Type }}) EncodeColumn(b *Buffer) {
	if len(c) == 0 {
		return
	}
	offset := len(b.Buf)
{{- if .SingleByte }}
	b.Buf = append(b.Buf, make([]byte, len(c))...)
{{- else }}
	const size = {{ .Bits }} / 8
	b.Buf = append(b.Buf, make([]byte, size * len(c))...)
{{- end }}
	s := *(*slice)(unsafe.Pointer(&c))
{{- if not .SingleByte }}
	s.Len *= size
	s.Cap *= size
{{- end }}
	src := *(*[]byte)(unsafe.Pointer(&s))
    dst := b.Buf[offset:]
	copy(dst, src)
}
