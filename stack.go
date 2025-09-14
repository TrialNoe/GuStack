package GuStack

type (
	GuStack struct {
		buffer []byte
		offset int
		len    int
	}
)

// New 注意操作是无锁的
func New() *GuStack {
	Len := 512
	g := &GuStack{
		buffer: make([]byte, Len),
		offset: 0,
		len:    Len,
	}
	return g
}

func (g *GuStack) Push(i int) {
	//先拉长缓冲区
	g.buffer[g.offset] = (byte)(i >> 24)
	g.buffer[g.offset+1] = (byte)(i >> 16)
	g.buffer[g.offset+2] = (byte)(i >> 8)
	g.buffer[g.offset+3] = (byte)(i >> 0)
	g.offset += 4
}

func (g *GuStack) Pop() int {
	i := g.offset - 4
	if i < 0 {
		panic("可恶，出栈次数居然比入栈次数多，你看看你写的什么代码")
	}
	bArr := g.buffer[i:g.offset]
	g.offset = i
	return (int(bArr[0])<<24)&-16777216 | ((int(bArr[1]) << 16) & 16711680) | ((int(bArr[2]) << 8) & 65280) | (int(bArr[3]<<0) & 255)

}
