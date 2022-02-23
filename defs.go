import (
	"io"
)

type fd struct {
	ws bufio.Writer
	rs bufio.Reader
	readable bool
	file string
	fl *FILE
}

func MakeFd ( file string, read bool ) ( fd ) {
	var f *FILE
	var err error
	if ( read ) {
		f, err = fopen(file)
		panic(err)
		return fd{
			rs: *bufio.NewReader(io.Reader(f)),
			file: file,
			readable: read, // always true
			fl: f,
		}
	} else {
		f, err = os.Create(file)
		panic(err)
		return fd{
			ws: *bufio.NewWriter(io.Writer(f)),
			file: file,
			readable: read, // always false
			fl: f,
		}
	}
}
