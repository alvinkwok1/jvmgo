package rtdata

type OperandStack struct {
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}
