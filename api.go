package nhicard

import (
	"syscall"
)

func dllImport(dllPath string ,functionName string) *syscall.LazyProc {
	handle := syscall.NewLazyDLL(dllPath)
	dllLib := handle.NewProc(functionName)
	return dllLib
}

func hisGetBasicData() (int, []byte) {
	dllFunction := dllImport("hisGetBasicData")
	pBuffer := make([]byte, 72)
	errorCode, _, _ := dllFunction.Call(BytePtr(pBuffer), Int2IntPtr(72))
	return errorCode, pBuffer
}

func csOpenCom(p int) {

}
