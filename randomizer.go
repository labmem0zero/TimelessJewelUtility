package main

const (
	INITIAL_STATE_CONSTANT_0 = 0x40336050
	INITIAL_STATE_CONSTANT_1 = 0xCFA3723C
	INITIAL_STATE_CONSTANT_2 = 0x3CAC5F6F
	INITIAL_STATE_CONSTANT_3 = 0x3793FDFF
)

type Rngesus struct {
	states [5]uint32
}

func (r *Rngesus) InitRandomize(seeds []uint32) {
	r.states = [5]uint32{
		0,
		INITIAL_STATE_CONSTANT_0,
		INITIAL_STATE_CONSTANT_1,
		INITIAL_STATE_CONSTANT_2,
		INITIAL_STATE_CONSTANT_3,
	}
	index := uint32(1)
	for _, s := range seeds {
		roundState := ManipulateAlpha(
			r.states[(index%4)+1] ^
				r.states[((index+1)%4)+1] ^
				r.states[(((index+4)-1)%4)+1])
		r.states[((index+1)%4)+1] += roundState
		roundState += s + index
		r.states[((index+2)%4)+1] += roundState
		r.states[(index%4)+1] = roundState
		index = (index + 1) % 4
	}

	for i := 0; i < 5; i++ {
		roundState := ManipulateAlpha(r.states[index%4+1] ^ r.states[(index+1)%4+1] ^ r.states[(index+3)%4+1])
		r.states[((index+1)%4)+1] += roundState
		roundState += index
		r.states[((index+2)%4)+1] += roundState
		r.states[(index%4)+1] = roundState
		index = (index + 1) % 4
	}

	for i := 0; i < 4; i++ {
		roundState := ManipulateBravo(r.states[(index%4)+1] +
			r.states[((index+1)%4)+1] +
			r.states[(((index+4)-1)%4)+1])
		r.states[((index+1)%4)+1] ^= roundState
		roundState -= index
		r.states[(((index+1)+1)%4)+1] ^= roundState
		r.states[(index%4)+1] = roundState
		index = (index + 1) % 4
	}
	for i := 0; i < 8; i++ {
		r.GenerateNextState()
	}
}

func (r *Rngesus) GenerateNextState() {
	a := uint32(0)
	b := uint32(0)
	a = r.states[4]
	b = ((r.states[1] & 0x7FFFFFFF) ^ r.states[2]) ^ r.states[3]

	a ^= a << 1
	b ^= (b >> 1) ^ a

	r.states[1] = r.states[2]
	r.states[2] = r.states[3]
	r.states[3] = a ^ (b << 10)
	r.states[4] = b

	r.states[2] ^= (uint32)(-((int)(b & 1)) & 0x8F7011EE)
	r.states[3] ^= (uint32)(-((int)(b & 1)) & 0xFC78FF1F)

	r.states[0]++
}

func (r *Rngesus) Temper() uint32 {
	a := r.states[4]
	b := r.states[1] + (r.states[3] >> 8)
	a ^= b
	if b&1 != 0 {
		a ^= 0x3793FDFF
	}
	return a
}

func ManipulateAlpha(value uint32) uint32 {
	return (value ^ (value >> 27)) * 0x19660D
}

func ManipulateBravo(val uint32) uint32 {
	return (val ^ (val >> 27)) * 0x5D588B65
}

func (r *Rngesus) GenerateUint32() uint32 {
	r.GenerateNextState()
	return r.Temper()
}

func (r *Rngesus) GenerateOne(val uint32) uint32 {
	max := val - 1
	roundState := uint32(0)
	tmpVal := uint32(0)
	for (tmpVal/val >= roundState) && (roundState%val != max) {
		for roundState < max {
			tmpVal = r.GenerateUint32() | (2 * (tmpVal << 31))
			roundState = 0xFFFFFFFF | (2 * (roundState << 31))
		}
	}
	return tmpVal % val
}

func (r *Rngesus) GenerateTwo(min, max uint32) uint32 {
	a := min + 0x80000000
	b := max + 0x80000000

	if min >= 0x80000000 {
		a = min + 0x80000000
	}

	if max >= 0x80000000 {
		a = max + 0x80000000
	}

	roll := r.GenerateOne(b - a + 1)
	return roll + a + 0x80000000
}
