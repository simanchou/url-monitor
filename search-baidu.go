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
	err := chromedp.Run(ctx, submit2(`https://www.baidu.com/`, `//input[@id="kw"]`, `jenkins`, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("got: `%s`", res)
}

func submit2(urlstr, sel, q string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, q),
		//chromedp.Click(`//input[@id="su"]`),
		chromedp.Submit(sel),
		chromedp.OuterHTML(`//*[@id="2"]`, res),
	}
}
