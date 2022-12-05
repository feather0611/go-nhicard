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

/* hisGetBasicData */
func (l *Nhilib) HisGetBasicData() (uintptr, []byte) {
	l.SetFunction("hisGetBasicData")
	bufferLen := 72
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetRegisterBasic */
func (l *Nhilib) HisGetRegisterBasic() (uintptr, []byte) {
	l.SetFunction("hisGetRegisterBasic")
	bufferLen := 78
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetRegisterPrevent */
func (l *Nhilib) HisGetRegisterPrevent() (uintptr []byte) {
	l.SetFunction("hisGetRegisterPrevent")
	bufferLen := 126
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetRegisterPregnant */
func (l *Nhilib) HisGetRegisterPregnant() (uintptr []byte) {
	l.SetFunction("hisGetRegisterPregnant")
	bufferLen := 209
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetTreatmentNoNeedHPC */
func (l *Nhilib) HisGetTreatmentNoNeedHPC() (uintptr []byte) {
	l.SetFunction("hisGetTreatmentNoNeedHPC")
	bufferLen := 498
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetCumulativeData */
func (l *Nhilib) HisGetCumulativeData() (uintptr []byte) {
	l.SetFunction("hisGetCumulativeData")
	bufferLen := 134
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetCumulativeFee */
func (l *Nhilib) HisGetCumulativeFee() (uintptr []byte) {
	l.SetFunction("hisGetCumulativeFee")
	bufferLen := 20
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetTreatmentNeedHPC */
func (l *Nhilib) HisGetTreatmentNeedHPC() (uintptr []byte) {
	l.SetFunction("hisGetTreatmentNeedHPC")
	bufferLen := 372
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

func (l *Nhilib) CsOpenCom(p int) uintptr {
	l.SetFunction("csOpenCom")
	errorCode, _, _ := l.proc.Call(Int2IntPtr(p))
	return errorCode
}

func (l *Nhilib) CsCloseCom() uintptr {
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
