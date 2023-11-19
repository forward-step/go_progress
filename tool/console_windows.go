package tool

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32DLL                     = syscall.MustLoadDLL("kernel32.dll")
	setConsoleCursorPositionProc    = kernel32DLL.MustFindProc("SetConsoleCursorPosition")
	getConsoleScreenBufferInfoProc  = kernel32DLL.MustFindProc("GetConsoleScreenBufferInfo")
	fillConsoleOutputCharacterWProc = kernel32DLL.MustFindProc("FillConsoleOutputCharacterW")
)

type coord struct {
	x, y int16
}
type smallRect struct {
	left, top, right, bottom int16
}
type consoleScreenBufferInfo struct {
	dwSize              coord
	dwCursorPosition    coord
	wAttributes         uint16
	srWindow            smallRect
	dwMaximumWindowSize coord
}

func ClearLine(file *os.File, line int) {
	if line == 0 || file == nil {
		return
	}

	var fd = file.Fd()
	var info consoleScreenBufferInfo
	_, _, _ = getConsoleScreenBufferInfoProc.Call(fd, uintptr(unsafe.Pointer(&info)))
	for i := 1; i <= line; i++ {
		// 设置光标的位置
		_, _, _ = setConsoleCursorPositionProc.Call(fd, uintptr(*(*int32)(unsafe.Pointer(&info.dwCursorPosition))))
		// 清空当前这行
		cursor := coord{
			x: info.srWindow.left,
			y: info.srWindow.top + info.dwCursorPosition.y,
		}
		var w int
		count := info.dwSize.x
		_, _, _ = fillConsoleOutputCharacterWProc.Call(fd, uintptr(' '), uintptr(count), *(*uintptr)(unsafe.Pointer(&cursor)), uintptr(unsafe.Pointer(&w)))
		// 移动到上一行
		info.dwCursorPosition.y--
		info.dwCursorPosition.x = 0
	}
}
