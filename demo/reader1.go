package main
import(
	"fmt"
	//"strings" 相关搬迁至本文件，易于保存
	"io"
)

func main() {
	var r *reader= newReader("Hello, Reader!")
	var b []byte = make([]byte, 8)    // 8 这里控制每次读取的字节数
	for{
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF{
			break
		}
	}
}

/*
A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
io.ByteScanner, and io.RuneScanner interfaces by reading from a string.
*/
type reader struct {
	str        string
	index        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *reader) Read(b []byte) (n int, err error) {
	if r.index >= int64(len(r.str)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.str[r.index:])
	r.index += int64(n)
	return
}

func newReader(s string) *reader { return &reader{s, 0, -1} }