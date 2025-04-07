package rtdata

type Thread struct {
	pc    int    // 程序计数器
	stack *Stack // 线程内的虚拟机栈
}

func NewThread() *Thread {
	// TODO 可以通过jvm指令-Xss来设置栈大小
	return &Thread{
		stack: NewStack(1024),
	}
}

// 压栈
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 弹栈
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 返回当前栈顶元素
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
