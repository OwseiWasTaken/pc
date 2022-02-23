package main

include "gutil"
include "defs.go"

include "syscall.go"

func main(){
	InitGu()

	var (
		fl fd
	)
	fl = MakeFd("log.txt", false)

	S_write(fl, "hello!\n")

	S_exit(4)

	exit(0)
}


// fl = fopen("log.txt", FD_WRITE)
// S_write(fl, "hello!\n")
// S_exit(0)
//
//
//

