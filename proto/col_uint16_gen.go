// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColUInt16 represents UInt16 column.
type ColUInt16 []uint16

// Compile-time assertions for ColUInt16.
var (
	_ ColInput  = ColUInt16{}
	_ ColResult = (*ColUInt16)(nil)
	_ Column    = (*ColUInt16)(nil)
)

// Type returns ColumnType of UInt16.
func (ColUInt16) Type() ColumnType {
	return ColumnTypeUInt16
}

// Rows returns count of rows in column.
func (c ColUInt16) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColUInt16) Row(i int) uint16 {
	return c[i]
}

// Append uint16 to column.
func (c *ColUInt16) Append(v uint16) {
	*c = append(*c, v)
}

// AppendArr appends slice of uint16 to column.
func (c *ColUInt16) AppendArr(v []uint16) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt16) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for UInt16 .
func (c *ColUInt16) LowCardinality() *ColLowCardinalityOf[uint16] {
	return &ColLowCardinalityOf[uint16]{
		index: c,
	}
}

// Array is helper that creates Array of uint16.
func (c *ColUInt16) Array() *ColArrOf[uint16] {
	return &ColArrOf[uint16]{
		Data: c,
	}
}

// NewArrUInt16 returns new Array(UInt16).
func NewArrUInt16() *ColArr {
	return &ColArr{
		Data: new(ColUInt16),
	}
}

// AppendUInt16 appends slice of uint16 to Array(UInt16).
func (c *ColArr) AppendUInt16(data []uint16) {
	d := c.Data.(*ColUInt16)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
