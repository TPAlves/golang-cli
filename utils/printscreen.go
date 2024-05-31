package utils

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GetChromeScreenShot(url string, quality int) {
	messageErro := "Não foi possível salvar a captura\n"
	screenShotURL := url

	if !strings.HasPrefix(screenShotURL, "http") {
		screenShotURL = fmt.Sprintf("http://%s", url)
	}

	nameSite := extractWebsiteName(screenShotURL)

	var buf []byte
	var extension string = "png"

	if quality < 100 {
		extension = "jpeg"
	}

	log.Printf("Capturando a tela do site %s", screenShotURL)

	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1400, 2000))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	tasks := chromedp.Tasks{
		chromedp.Navigate(screenShotURL),
		chromedp.Sleep(3 * time.Second),
		chromedp.CaptureScreenshot(&buf),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(messageErro, err)
	}

	fileName := fmt.Sprintf("%s-%d.%s", nameSite, time.Now().UTC().Unix(), extension)
	if err := os.WriteFile(fileName, buf, 0644); err != nil {
		log.Fatal(messageErro, err)
	}

	log.Printf("Captura salva em %s", fileName)
}

func extractWebsiteName(site string) string {
	u, err := url.Parse(site)
	if err != nil {
		log.Fatal("URL inválida", err)
	}
	name := strings.Split(u.Hostname(), ".")
	if name[0] == "www" {
		return name[1]
	}

	return name[0]
}
