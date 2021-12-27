// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColInt64 represents Int64 column.
type ColInt64 []int64

// Compile-time assertions for ColInt64.
var (
	_ ColInput  = ColInt64{}
	_ ColResult = (*ColInt64)(nil)
	_ Column    = (*ColInt64)(nil)
)

// Type returns ColumnType of Int64.
func (ColInt64) Type() ColumnType {
	return ColumnTypeInt64
}

// Rows returns count of rows in column.
func (c ColInt64) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt64) Row(i int) int64 {
	return c[i]
}

// Append int64 to column.
func (c *ColInt64) Append(v int64) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt64) Reset() {
	*c = (*c)[:0]
}

// NewArrInt64 returns new Array(Int64).
func NewArrInt64() *ColArr {
	return &ColArr{
		Data: new(ColInt64),
	}
}

// AppendInt64 appends slice of int64 to Array(Int64).
func (c *ColArr) AppendInt64(data []int64) {
	d := c.Data.(*ColInt64)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Int64 rows to *Buffer.
func (c ColInt64) EncodeColumn(b *Buffer) {
	const size = 64 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint64(
			b.Buf[offset:offset+size],
			uint64(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Int64 rows from *Reader.
func (c *ColInt64) DecodeColumn(r *Reader, rows int) error {
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
			int64(binary.LittleEndian.Uint64(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
