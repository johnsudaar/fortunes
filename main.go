package main

import (
	"fmt"
	"os"
	"path"

	"github.com/go-opencv/go-opencv/opencv"
	"github.com/johnsudaar/fortunes/picker"
	"github.com/johnsudaar/fortunes/reader"
)

func main() {
	picker, err := picker.LoadPicker("fortunes.txt")
	//picker, err := picker.LoadPicker("kamouscope.txt")
	if err != nil {
		panic(err)
	}

	reader := reader.NewReader()

	win := opencv.NewWindow("Go-OpenCV Webcam Face Detection", opencv.CV_WND_PROP_FULLSCREEN)
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("cannot open camera")
	}
	defer cap.Release()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cascade := opencv.LoadHaarClassifierCascade(path.Join(cwd, "haarcascade_frontalface_alt.xml"))

	fmt.Println("Press ESC to quit")
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				faces := cascade.DetectObjects(img)
				for _, value := range faces {
					opencv.Circle(img,
						opencv.Point{
							value.X() + (value.Width() / 2),
							value.Y() + (value.Height() / 2),
						},
						value.Width()/2,
						opencv.ScalarAll(255.0), 1, 1, 0)
				}
				win.ShowImage(img)
				if len(faces) > 0 {
					reader.Read(picker.Pick())
				}

			} else {
				fmt.Println("nil image")
			}
		}
		key := opencv.WaitKey(1)

		if key == 27 {
			os.Exit(0)
		}
	}
}
