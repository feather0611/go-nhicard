package nhicard

import "unsafe"

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func Int2IntPtr(n int) uintptr {
	return uintptr(unsafe.Pointer(&n))
}

func IntPtr2Ptr(n *int) uintptr {
	return uintptr(unsafe.Pointer(n))
}

func BytePtr(s []byte) uintptr {
	return uintptr(unsafe.Pointer(&s[0]))
}
