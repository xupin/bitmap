package bitmap

type bitMap struct {
	bits []byte
	max  uint64
}

func New(max uint64) *bitMap {
	bm := &bitMap{}
	bm.max = max
	bm.bits = make([]byte, (max>>3)+1)
	return bm
}

// 将1左移{num % 8}位，按位或运算把对应bit设为1
func (r *bitMap) Add(num uint64) {
	r.bits[num>>3] |= 1 << (num % 8)
}

// 将1左移{num % 8}位后，按位与、异或运算把相同bit设为0
func (r *bitMap) Del(num uint64) {
	r.bits[num>>3] &^= 1 << (num % 8)
}

// 将1左移{num % 8}位后，按位与运算判断是否bit都为1
func (r *bitMap) Lookup(num uint64) bool {
	return r.bits[num>>3]&1<<(num%8) != 0
}

func (r *bitMap) GetBits() []byte {
	return r.bits
}
