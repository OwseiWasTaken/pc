// pack about stack and mem

const Heap_Size = 100

type _indv_ struct{ x int }
var _i_fram = &(_indv_{0}) // free ram
var _i_rram = &(_indv_{1}) // ref  ram

var HEAP = MakeArray(Heap_Size, _MLC_nothing)

func InitRam() {
	for i:=0;i<Heap_Size;i++ {
		HEAP[i] = _i_fram
	}
}
