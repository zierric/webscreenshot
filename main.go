package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	webUrl := "https://habr.com/"
	windowWidth := 1440
	windowHeight := 900

	waitDuration := 2
	userAgent := "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36"

	options := make([]chromedp.ExecAllocatorOption, 0)
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	options = append(options, chromedp.DisableGPU)
	options = append(options, chromedp.UserAgent(userAgent))
	options = append(options, chromedp.Flag("ignore-certificate-errors", true))
	options = append(options, chromedp.WindowSize(windowWidth, windowHeight))

	allocatorContext, allocatorCancel := chromedp.NewExecAllocator(context.Background(), options...)
	ctx, cancel := chromedp.NewContext(allocatorContext)

	defer allocatorCancel()
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(webUrl),
		chromedp.Sleep(time.Duration(waitDuration) * time.Second),
		chromedp.CaptureScreenshot(&buf),
	}); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("example.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
