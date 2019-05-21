// Command submit is a chromedp example demonstrating how to fill out and
// submit a form.
package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {

	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		//chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"),
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx, submit1(`http://127.0.0.1:9001/`, `/html/body/div/div[2]/div[2]/div[2]/div/div/div/div/form/input[1]`, `t6.com`, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got: `%s`", res)
}

func submit1(urlstr, sel, q string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		chromedp.Submit(sel),
		chromedp.OuterHTML("/html/body/div", res),
	}
}
