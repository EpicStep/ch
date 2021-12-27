// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColDate represents Date column.
type ColDate []Date

// Compile-time assertions for ColDate.
var (
	_ ColInput  = ColDate{}
	_ ColResult = (*ColDate)(nil)
	_ Column    = (*ColDate)(nil)
)

// Type returns ColumnType of Date.
func (ColDate) Type() ColumnType {
	return ColumnTypeDate
}

// Rows returns count of rows in column.
func (c ColDate) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDate) Row(i int) Date {
	return c[i]
}

// Append Date to column.
func (c *ColDate) Append(v Date) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDate) Reset() {
	*c = (*c)[:0]
}

// NewArrDate returns new Array(Date).
func NewArrDate() *ColArr {
	return &ColArr{
		Data: new(ColDate),
	}
}

// AppendDate appends slice of Date to Array(Date).
func (c *ColArr) AppendDate(data []Date) {
	d := c.Data.(*ColDate)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Date rows to *Buffer.
func (c ColDate) EncodeColumn(b *Buffer) {
	const size = 16 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint16(
			b.Buf[offset:offset+size],
			uint16(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Date rows from *Reader.
func (c *ColDate) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 16 / 8
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
			Date(binary.LittleEndian.Uint16(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
