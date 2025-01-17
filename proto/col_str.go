package proto

import (
	"encoding/binary"

	"github.com/go-faster/errors"
)

type Position struct {
	Start int
	End   int
}

// ColStr represents String column.
//
// Use ColBytes for []bytes ColumnOf implementation.
type ColStr struct {
	Buf []byte
	Pos []Position
}

// Append string to column.
func (c *ColStr) Append(v string) {
	start := len(c.Buf)
	c.Buf = append(c.Buf, v...)
	end := len(c.Buf)
	c.Pos = append(c.Pos, Position{Start: start, End: end})
}

// AppendBytes append byte slice as string to column.
func (c *ColStr) AppendBytes(v []byte) {
	start := len(c.Buf)
	c.Buf = append(c.Buf, v...)
	end := len(c.Buf)
	c.Pos = append(c.Pos, Position{Start: start, End: end})
}

func (c *ColStr) AppendArr(v []string) {
	for _, e := range v {
		c.Append(e)
	}
}

// ArrAppend appends data to array of ColStr.
func (ColStr) ArrAppend(arr *ColArr, data []string) {
	c := arr.Data.(*ColStr)
	for _, v := range data {
		c.Append(v)
	}
	c.Rows()
	arr.Offsets = append(arr.Offsets, uint64(len(c.Pos)))
}

// Compile-time assertions for ColStr.
var (
	_ ColInput  = ColStr{}
	_ ColResult = (*ColStr)(nil)
	_ Column    = (*ColStr)(nil)
)

// Type returns ColumnType of String.
func (ColStr) Type() ColumnType {
	return ColumnTypeString
}

// Rows returns count of rows in column.
func (c ColStr) Rows() int {
	return len(c.Pos)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColStr) Reset() {
	c.Buf = c.Buf[:0]
	c.Pos = c.Pos[:0]
}

// EncodeColumn encodes String rows to *Buffer.
func (c ColStr) EncodeColumn(b *Buffer) {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, p := range c.Pos {
		n := binary.PutUvarint(buf, uint64(p.End-p.Start))
		b.Buf = append(b.Buf, buf[:n]...)
		b.Buf = append(b.Buf, c.Buf[p.Start:p.End]...)
	}
}

// ForEach calls f on each string from column.
func (c ColStr) ForEach(f func(i int, s string) error) error {
	return c.ForEachBytes(func(i int, b []byte) error {
		return f(i, string(b))
	})
}

// First returns first row of column.
func (c ColStr) First() string {
	return c.Row(0)
}

// Row returns row with number i.
func (c ColStr) Row(i int) string {
	p := c.Pos[i]
	return string(c.Buf[p.Start:p.End])
}

// RowBytes returns row with number i as byte slice.
func (c ColStr) RowBytes(i int) []byte {
	p := c.Pos[i]
	return c.Buf[p.Start:p.End]
}

// ForEachBytes calls f on each string from column as byte slice.
func (c ColStr) ForEachBytes(f func(i int, b []byte) error) error {
	for i, p := range c.Pos {
		if err := f(i, c.Buf[p.Start:p.End]); err != nil {
			return err
		}
	}
	return nil
}

// DecodeColumn decodes String rows from *Reader.
func (c *ColStr) DecodeColumn(r *Reader, rows int) error {
	var p Position
	for i := 0; i < rows; i++ {
		n, err := r.StrLen()
		if err != nil {
			return errors.Wrapf(err, "row %d: read length", i)
		}

		p.Start = p.End
		p.End += n

		c.Buf = append(c.Buf, make([]byte, n)...)
		if err := r.ReadFull(c.Buf[p.Start:p.End]); err != nil {
			return errors.Wrapf(err, "row %d: read full", i)
		}
		c.Pos = append(c.Pos, p)
	}
	return nil
}

// LowCardinality returns LowCardinality(String).
func (c *ColStr) LowCardinality() *ColLowCardinalityOf[string] {
	return &ColLowCardinalityOf[string]{
		index: c,
	}
}

// Array is helper that creates Array(String).
func (c *ColStr) Array() *ColArrOf[string] {
	return &ColArrOf[string]{
		Data: c,
	}
}

// ColBytes is ColStr wrapper to be ColumnOf for []byte.
type ColBytes struct {
	ColStr
}

// Row returns row with number i.
func (c ColBytes) Row(i int) []byte {
	return c.RowBytes(i)
}

// Append byte slice to column.
func (c *ColBytes) Append(v []byte) {
	c.AppendBytes(v)
}

// AppendArr append slice of byte slices to column.
func (c *ColBytes) AppendArr(v [][]byte) {
	for _, s := range v {
		c.Append(s)
	}
}

// Array is helper that creates Array(String).
func (c *ColBytes) Array() *ColArrOf[[]byte] {
	return &ColArrOf[[]byte]{
		Data: c,
	}
}
