package terminalshaders

import (
	"bytes"
	"fmt"
	"math"
	"syscall"
	"time"
	"unsafe"
)

type winsize struct {
	ws_row uint16
	ws_col uint16
}

type options struct {
	AnsiMode  *bool
	Framerate *int
}

var opts = options{}

func SetAnsiMode(mode bool) {
	opts.AnsiMode = &mode
}

func SetFramerate(duration int) {
	opts.Framerate = &duration
}

func getTerminalSize() (width, height int) {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(retCode) == -1 {
		panic(errno)
	}

	return int(ws.ws_col), int(ws.ws_row)
}

func mapColor(r, g, b float64, useTrueColor bool) string {
	if useTrueColor {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm█", int(r*255), int(g*255), int(b*255))
	}

	return fmt.Sprintf("\x1b[38;5;%dm█", 16+int(r*5)*36+int(g*5)*6+int(b*5))
}

func Run(shaderName string) {
	width, height := getTerminalSize()
	fmt.Print("\033[H\033[2J")

	useTrueColor := true
	if opts.AnsiMode != nil && *opts.AnsiMode {
		useTrueColor = false
	}

	framerate := time.Duration(30 * time.Nanosecond)
	if opts.Framerate != nil {
		framerate = time.Duration(*opts.Framerate) * time.Nanosecond
	}

	activeShader := GetShader(shaderName)
	if activeShader == nil {
		fmt.Printf("Shader %s not found", shaderName)
		return
	}

	lastTime := time.Now()
	for {
		timeNow := float64(time.Now().UnixNano()) / 1e9
		var buffer bytes.Buffer

		for y := 0; y < height; y++ {
			ny := 1 - float64(y)/float64(height)

			for x := 0; x < width-1; x++ {
				nx := float64(x) / float64(width)
				color := activeShader.Compute(Vec2{X: nx, Y: ny}, timeNow)

				color.R = math.Max(0, math.Min(1, color.R))
				color.G = math.Max(0, math.Min(1, color.G))
				color.B = math.Max(0, math.Min(1, color.B))

				buffer.WriteString(mapColor(color.R, color.G, color.B, useTrueColor))
			}

			buffer.WriteString("\x1b[0m\n")
		}

		_, err := fmt.Print("\033[H", buffer.String())
		if err != nil {
			fmt.Printf("Shader rendering failed: %s", err)
			return
		}

		elapsed := time.Since(lastTime)
		if elapsed < time.Second/framerate {
			time.Sleep(time.Second/framerate - elapsed)
		}

		lastTime = time.Now()
	}
}
