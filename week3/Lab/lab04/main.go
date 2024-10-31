package main

//#cgo LDFLAGS: -L. -lcstring_concat
//char* string_concat(const char* s1, const char* s2);
import "C"
import (
	"fmt"
)

func main() {
	s1 := C.CString("Hello ")
	s2 := C.CString("World")
	result := C.string_concat(s1, s2)

	res := C.GoString(result)
	fmt.Println("Concatenated string:", res)

	//C.free(result)
}
