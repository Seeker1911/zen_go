package main

import (
	"fmt"
	"image/png"
	"math"
	"os"
	"os/exec"
	"strconv"

	"github.com/vova616/screenshot"
)

func rgb(i int) (int, int, int) {
	// lolcats implementation
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func lolPrint(output []byte) {
	// take screenshot
	defer screen()

	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}

// SetFromFile uses AppleScript to tell Finder to set the desktop wallpaper to specified file.
func SetFromFile(file string) error {
	return exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(file)).Run()
}

func screen() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		println(err.Error())
		return
	}
	f, err := os.Create("./zen_.png")

	if err != nil {
		println(err.Error())
		return
	}

	err = png.Encode(f, img)

	if err != nil {
		println(err.Error())
		return
	}
	f.Close()

	dir, err := os.Getwd()
	if err != nil {
		println(err.Error())
	}
	file := string(dir + "/zen_.png")

	//set wallpaper
	defer SetFromFile(file)

}

func main() {
	// run python import this
	cmd := exec.Command("./script.py", " | ", "lolcat")

	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	// pipe to lolcat
	lolPrint(out)

}
