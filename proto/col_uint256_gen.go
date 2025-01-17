// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColUInt256 represents UInt256 column.
type ColUInt256 []UInt256

// Compile-time assertions for ColUInt256.
var (
	_ ColInput  = ColUInt256{}
	_ ColResult = (*ColUInt256)(nil)
	_ Column    = (*ColUInt256)(nil)
)

// Type returns ColumnType of UInt256.
func (ColUInt256) Type() ColumnType {
	return ColumnTypeUInt256
}

// Rows returns count of rows in column.
func (c ColUInt256) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColUInt256) Row(i int) UInt256 {
	return c[i]
}

// Append UInt256 to column.
func (c *ColUInt256) Append(v UInt256) {
	*c = append(*c, v)
}

// AppendArr appends slice of UInt256 to column.
func (c *ColUInt256) AppendArr(v []UInt256) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt256) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for UInt256 .
func (c *ColUInt256) LowCardinality() *ColLowCardinalityOf[UInt256] {
	return &ColLowCardinalityOf[UInt256]{
		index: c,
	}
}

// Array is helper that creates Array of UInt256.
func (c *ColUInt256) Array() *ColArrOf[UInt256] {
	return &ColArrOf[UInt256]{
		Data: c,
	}
}

// NewArrUInt256 returns new Array(UInt256).
func NewArrUInt256() *ColArr {
	return &ColArr{
		Data: new(ColUInt256),
	}
}

// AppendUInt256 appends slice of UInt256 to Array(UInt256).
func (c *ColArr) AppendUInt256(data []UInt256) {
	d := c.Data.(*ColUInt256)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
