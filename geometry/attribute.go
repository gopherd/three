package geometry

import "constraints"

type Attribute interface {
	Count() int
	Stride() int
	GetFloat32(offset int) float32
	GetFloat64(offset int) float64
}

type BufferAttribute[T constraints.Integer | constraints.Float] struct {
	data   []T
	count  int
	stride int
}

func NewBufferAttribute[T constraints.Integer | constraints.Float](count, stride int) *BufferAttribute[T] {
	return &BufferAttribute[T]{
		data:   make([]T, count*stride),
		count:  count,
		stride: stride,
	}
}

func (attribute *BufferAttribute[T]) GetFloat32(offset int) float32 {
	return float32(attribute.data[offset])
}

func (attribute *BufferAttribute[T]) GetFloat64(offset int) float64 {
	return float64(attribute.data[offset])
}

func (attribute *BufferAttribute[T]) Count() int {
	return attribute.count
}

func (attribute *BufferAttribute[T]) Stride() int {
	return attribute.stride
}

func (attribute *BufferAttribute[T]) GetX(index int) T {
	return attribute.data[index*attribute.stride]
}

func (attribute *BufferAttribute[T]) GetY(index int) T {
	return attribute.data[index*attribute.stride+1]
}

func (attribute *BufferAttribute[T]) GetZ(index int) T {
	return attribute.data[index*attribute.stride+2]
}

func (attribute *BufferAttribute[T]) GetW(index int) T {
	return attribute.data[index*attribute.stride+3]
}

func (attribute *BufferAttribute[T]) SetX(index int, x T) {
	attribute.data[index*attribute.stride] = x
}

func (attribute *BufferAttribute[T]) SetY(index int, y T) {
	attribute.data[index*attribute.stride+1] = y
}

func (attribute *BufferAttribute[T]) SetZ(index int, z T) {
	attribute.data[index*attribute.stride+2] = z
}

func (attribute *BufferAttribute[T]) SetW(index int, w T) {
	attribute.data[index*attribute.stride+3] = w
}

func (attribute *BufferAttribute[T]) SetXY(index int, x, y T) {
	offset := index * attribute.stride
	attribute.data[offset] = x
	attribute.data[offset+1] = y
}

func (attribute *BufferAttribute[T]) SetXYZ(index int, x, y, z T) {
	offset := index * attribute.stride
	attribute.data[offset] = x
	attribute.data[offset+1] = y
	attribute.data[offset+2] = z
}

func (attribute *BufferAttribute[T]) SetXYZW(index int, x, y, z, w T) {
	offset := index * attribute.stride
	attribute.data[offset] = x
	attribute.data[offset+1] = y
	attribute.data[offset+2] = z
	attribute.data[offset+3] = w
}

type Int8Attribute = BufferAttribute[int8]
type Int16Attribute = BufferAttribute[int16]
type Int32Attribute = BufferAttribute[int32]
type Uint8Attribute = BufferAttribute[uint8]
type Uint16Attribute = BufferAttribute[uint16]
type Uint32Attribute = BufferAttribute[uint32]
type Float32Attribute = BufferAttribute[float32]
type Float64Attribute = BufferAttribute[float64]

func NewInt8Attribute(count, stride int) *Int8Attribute {
	return NewBufferAttribute[int8](count, stride)
}

func NewInt16Attribute(count, stride int) *Int16Attribute {
	return NewBufferAttribute[int16](count, stride)
}

func NewInt32Attribute(count, stride int) *Int32Attribute {
	return NewBufferAttribute[int32](count, stride)
}

func NewUint8Attribute(count, stride int) *Uint8Attribute {
	return NewBufferAttribute[uint8](count, stride)
}

func NewUint16Attribute(count, stride int) *Uint16Attribute {
	return NewBufferAttribute[uint16](count, stride)
}

func NewUint32Attribute(count, stride int) *Uint32Attribute {
	return NewBufferAttribute[uint32](count, stride)
}

func NewFloat32Attribute(count, stride int) *Float32Attribute {
	return NewBufferAttribute[float32](count, stride)
}

func NewFloat64Attribute(count, stride int) *Float64Attribute {
	return NewBufferAttribute[float64](count, stride)
}
