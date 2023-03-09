// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.WindowSize(504, 284),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel = chromedp.NewContext(
		ctx,
		//context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`http://192.168.21.95:8080/betInactive.html`, 100, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("fullScreenshot5.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
