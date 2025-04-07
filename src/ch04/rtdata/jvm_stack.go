package rtdata

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	frame.lower = self._top
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	frame := self._top
	self._top = frame.lower
	self.size--
	return frame
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
