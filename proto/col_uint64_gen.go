// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColUInt64 represents UInt64 column.
type ColUInt64 []uint64

// Compile-time assertions for ColUInt64.
var (
	_ ColInput  = ColUInt64{}
	_ ColResult = (*ColUInt64)(nil)
	_ Column    = (*ColUInt64)(nil)
)

// Type returns ColumnType of UInt64.
func (ColUInt64) Type() ColumnType {
	return ColumnTypeUInt64
}

// Rows returns count of rows in column.
func (c ColUInt64) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColUInt64) Row(i int) uint64 {
	return c[i]
}

// Append uint64 to column.
func (c *ColUInt64) Append(v uint64) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt64) Reset() {
	*c = (*c)[:0]
}

// NewArrUInt64 returns new Array(UInt64).
func NewArrUInt64() *ColArr {
	return &ColArr{
		Data: new(ColUInt64),
	}
}

// AppendUInt64 appends slice of uint64 to Array(UInt64).
func (c *ColArr) AppendUInt64(data []uint64) {
	d := c.Data.(*ColUInt64)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes UInt64 rows to *Buffer.
func (c ColUInt64) EncodeColumn(b *Buffer) {
	const size = 64 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint64(
			b.Buf[offset:offset+size],
			v,
		)
		offset += size
	}
}

// DecodeColumn decodes UInt64 rows from *Reader.
func (c *ColUInt64) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 64 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	// Move bound check out of loop.
	//
	// See https://github.com/golang/go/issues/30945.
	_ = data[len(data)-size]
	for i := 0; i <= len(data)-size; i += size {
		v = append(v,
			binary.LittleEndian.Uint64(data[i:i+size]),
		)
	}
	*c = v
	return nil
}
