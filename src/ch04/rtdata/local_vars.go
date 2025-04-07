package rtdata

import "math"

type LocalVars []Slot

/*
*
创建局部变量表存储
*/
func NewLocalVars(maxLocals uint) LocalVars {
	if maxLocals == 0 {
		return nil
	}
	return make([]Slot, maxLocals)
}

/*
*
在局部变量表中存储int变量
*/
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

/*
在局部变量表中存储单浮点
需要将浮点类型转换威bits，然后bits转int处理
*/
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

/*
long 变量需要拆解成两个int变量
*/
func (self LocalVars) SetLong(index uint, val int64) {
	// 存储低位
	self[index].num = int32(val)
	// 存储高位
	self[index+1].num = int32(val >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

/*
double变量可以选转成go的long，然后按照long来处理
*/
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

/*
设置引用类型
*/
func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}
