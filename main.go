package main

// Command pdf is a chromedp example demonstrating how to capture a pdf of a
// page.
//package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	log.Println("Starting..")

	chromedpUrl := getDebugURL()


	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), chromedpUrl)
	defer cancelActxt()


	// create context
	ctx, cancel := chromedp.NewContext(actxt)
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(`https://www.google.com/`, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("sample.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Println("Completed")
}

func getDebugURL() string {
	baseUrl := "http://chromedp:9222/json/version"

	ips, err := net.LookupIP("chromedp")
	if err != nil {
		log.Fatal(err)
	}

	url := strings.Replace(baseUrl, "chromedp", ips[0].String(),1 )

	log.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//newStr := buf.String()
	//log.Println(newStr)

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result["webSocketDebuggerUrl"].(string)
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
