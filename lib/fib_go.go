package fib
import "C"
import "unsafe"

type GoFib struct {
	fib C.foo
}
func New()(GoFib){
	var ret GoFib;
	ret.foo = C.FibInit();
	return ret;
}
func (f GoFib)Free(){
	C.FibFree(unsafe.Pointer(f.foo));
}
func (f GoFib)Bar(){
	C.FibBar(unsafe.Pointer(f.foo));
}
