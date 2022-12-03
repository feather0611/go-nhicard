package nhicard

import (
	"syscall"
)

type Nhilib struct {
	path    string
	handler *syscall.LazyDLL
	proc    *syscall.LazyProc
}

func SetDll(path string) *Nhilib {
	lib := new(Nhilib)
	lib.path = path
	lib.handler = syscall.NewLazyDLL(path)
	return lib
}

func (l *Nhilib) SetFunction(functionName string) {
	l.proc = l.handler.NewProc(functionName)
}

func (l *Nhilib) GetBasicData() (uintptr, []byte) {
	l.SetFunction("hisGetBasicData")
	pBuffer := make([]byte, 72)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(72))
	return errorCode, pBuffer
}

func (l *Nhilib) OpenCom(p int) uintptr {
	l.SetFunction("csOpenCom")
	errorCode, _, _ := l.proc.Call(Int2IntPtr(p))
	return errorCode
}

func (l *Nhilib) CloseCom() uintptr {
	l.SetFunction("csCloseCom")
	errorCode, _, _ := l.proc.Call()
	return errorCode
}

func (l *Nhilib) GetDllVersion() uintptr {
	l.SetFunction("csGetVersionEx")
	addr := []byte(l.path)
	returnCode, _, _ := l.proc.Call(BytePtr(addr))
	return returnCode
}
