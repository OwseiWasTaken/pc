// pack about stack and mem

const Heap_Size = 100

type _indv_ struct{ x int }
var _i_fram = &(_indv_{0}) // free ram
var _i_rram = &(_indv_{1}) // ref  ram
var used uint
var Green string = RGB(0x60, 0xc0, 0x60)
var Red string = RGB(0xc0, 0x60, 0x60)
var _PrintHeapNlDiv = 10

var RAM = MakeArray(Heap_Size, _i_fram)

func InitRam() {
	for i:=0;i<Heap_Size;i++ {
		RAM[i] = _i_fram
	}
}

func GetStart( size uint ) int {
	var can bool
	if ( Heap_Size - used < size ) {
		fprintf(stderr,
"not enough memory!\ntried to allocate %d byte%s, only have %d (%d used)",
		size, bog(size>1, "s", ""), Heap_Size - used, used)
		exit(1)
	}
	for i:=uint(0);i+size<Heap_Size;i++ {
		can = true
		if ( RAM[i] != _i_fram) { continue } // used
		for j:=i;j<size;j++ {
			if ( RAM[j] != _i_fram) {
				can = false
			}
		}
		if (!can) { continue }
		return int(i)
	}
	return -1
}

func assing( size int ) []interface{} {
	var start int
	start = GetStart(uint(size))
	used += uint(size)
	if (start == -1) {
		PrintHeap()
		fprintf(stderr, "can't find where to allocate linear memory even with space (broken free mem blocks)")
		exit(2)
	}
	size+=start
	for i:=start;i<size;i++ {
		if (RAM[i] != _i_fram) {
			fprintf(stderr, "can't find where to allocate linear memory even with space (broken free mem blocks)")
			exit(2)
		}
		RAM[i] = NULL
	}
	return RAM[start:size]
}

func free ( allcd []interface{} ) {
	var start int
	var size int
	var findings int

	allcd[0] = _i_rram

	//malloc's &auth struct should not be allocated in the RAM! (unless in this func)
	for i:=0;i<Heap_Size;i++ {
		if RAM[i] == _i_rram{
			findings = i
			break
		}
	}
	start = findings
	size = len(allcd)+start
	for i:=start;i<size;i++ {
		//assinged[i] = false
		RAM[i] = _i_fram
	}
	used -= uint(size)
}

func PrintHeap() {
	for i:=0;i<Heap_Size;i++ {
		for j:=0;j<_PrintHeapNlDiv-1;j++ {
			if (i == Heap_Size) {break}
			printf("%s%s:", bog(RAM[i] != _i_fram, Green+"a", Red+"f"), RGB(0xff, 0xff, 0xff))
			printf("%v, ", bog(RAM[i] != _i_fram, RAM[i], "{}"))
			i++
		}
		if (i < Heap_Size) {
			printf("%s%s:", bog(RAM[i] != _i_fram, Green+"a", Red+"f"), RGB(0xff, 0xff, 0xff))
			printf("%v\n", bog(RAM[i] != _i_fram, RAM[i], "{}"))
		}
		if (i == Heap_Size) {break}
	}
	print("\n")
}
