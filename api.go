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
func (l *Nhilib) HisGetRegisterPrevent() (uintptr, []byte) {
	l.SetFunction("hisGetRegisterPrevent")
	bufferLen := 126
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetRegisterPregnant */
func (l *Nhilib) HisGetRegisterPregnant() (uintptr, []byte) {
	l.SetFunction("hisGetRegisterPregnant")
	bufferLen := 209
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetTreatmentNoNeedHPC */
func (l *Nhilib) HisGetTreatmentNoNeedHPC() (uintptr, []byte) {
	l.SetFunction("hisGetTreatmentNoNeedHPC")
	bufferLen := 498
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetCumulativeData */
func (l *Nhilib) HisGetCumulativeData() (uintptr, []byte) {
	l.SetFunction("hisGetCumulativeData")
	bufferLen := 134
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetCumulativeFee */
func (l *Nhilib) HisGetCumulativeFee() (uintptr, []byte) {
	l.SetFunction("hisGetCumulativeFee")
	bufferLen := 20
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetTreatmentNeedHPC */
func (l *Nhilib) HisGetTreatmentNeedHPC() (uintptr, []byte) {
	l.SetFunction("hisGetTreatmentNeedHPC")
	bufferLen := 372
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetSeqNumber */
func (l *Nhilib) HisGetSeqNumber(cTreatItem []byte, cBabyTreat []byte) (uintptr, []byte) {
	l.SetFunction("hisGetSeqNumber")
	bufferLen := 167
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(cTreatItem), BytePtr(cBabyTreat), BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisReadPrescription */
func (l *Nhilib) HisReadPrescription() (uintptr, []byte, []byte, []byte, []byte) {
	l.SetFunction("hisReadPrescription")
	outPatientBufferLen := 3660
	outPatientBuffer := make([]byte, outPatientBufferLen)
	longTermBufferLen := 1320
	longTermBuffer := make([]byte, longTermBufferLen)
	importantTreatmentLen := 360
	importantTreatmentBuffer := make([]byte, importantTreatmentLen)
	irritationDrugsLen := 120
	irritationDrugsBuffer := make([]byte, irritationDrugsLen)
	errorCode, _, _ := l.proc.Call(BytePtr(outPatientBuffer), Int2IntPtr(outPatientBufferLen), BytePtr(longTermBuffer), Int2IntPtr(longTermBufferLen), BytePtr(importantTreatmentBuffer), Int2IntPtr(importantTreatmentLen), BytePtr(irritationDrugsBuffer), Int2IntPtr(irritationDrugsLen))
	return errorCode, outPatientBuffer, longTermBuffer, importantTreatmentBuffer, irritationDrugsBuffer
}

/* hisGetInoculateData */
func (l *Nhilib) HisGetInoculateData() (uintptr, []byte) {
	l.SetFunction("hisGetInoculateData")
	bufferLen := 1400
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetOrganDonate */
func (l *Nhilib) HisGetOrganDonate() (uintptr, []byte) {
	l.SetFunction("hisGetOrganDonate")
	bufferLen := 1
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetEmergentTel */
func (l *Nhilib) HisGetEmergentTel() (uintptr, []byte) {
	l.SetFunction("hisGetEmergentTel")
	bufferLen := 14
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetLastSeqNum */
func (l *Nhilib) HisGetLastSeqNum() (uintptr, []byte) {
	l.SetFunction("hisGetLastSeqNum")
	bufferLen := 7
	pBuffer := make([]byte, bufferLen)
	errorCode, _, _ := l.proc.Call(BytePtr(pBuffer), Int2IntPtr(bufferLen))
	return errorCode, pBuffer
}

/* hisGetCardStatus */
func (l *Nhilib) HisGetCardStatus(cardType int) uintptr {
	l.SetFunction("hisGetCardStatus")
	statusCode, _, _ := l.proc.Call(Int2IntPtr(cardType))
	return statusCode
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
