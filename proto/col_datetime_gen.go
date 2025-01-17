// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColDateTime represents DateTime column.
type ColDateTime []DateTime

// Compile-time assertions for ColDateTime.
var (
	_ ColInput  = ColDateTime{}
	_ ColResult = (*ColDateTime)(nil)
	_ Column    = (*ColDateTime)(nil)
)

// Type returns ColumnType of DateTime.
func (ColDateTime) Type() ColumnType {
	return ColumnTypeDateTime
}

// Rows returns count of rows in column.
func (c ColDateTime) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDateTime) Row(i int) DateTime {
	return c[i]
}

// Append DateTime to column.
func (c *ColDateTime) Append(v DateTime) {
	*c = append(*c, v)
}

// AppendArr appends slice of DateTime to column.
func (c *ColDateTime) AppendArr(v []DateTime) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDateTime) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for DateTime .
func (c *ColDateTime) LowCardinality() *ColLowCardinalityOf[DateTime] {
	return &ColLowCardinalityOf[DateTime]{
		index: c,
	}
}

// Array is helper that creates Array of DateTime.
func (c *ColDateTime) Array() *ColArrOf[DateTime] {
	return &ColArrOf[DateTime]{
		Data: c,
	}
}

// NewArrDateTime returns new Array(DateTime).
func NewArrDateTime() *ColArr {
	return &ColArr{
		Data: new(ColDateTime),
	}
}

// AppendDateTime appends slice of DateTime to Array(DateTime).
func (c *ColArr) AppendDateTime(data []DateTime) {
	d := c.Data.(*ColDateTime)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
