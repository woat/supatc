// Package atc is used to add-to-cart and also can perform checkout too.
package atc

import (
	"context"
	"log"
	"time"

	"github.com/woat/supatc/cfg"
	"github.com/woat/supatc/inv"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

// Starts chromedp and takes the available inventory to process tasks with.
func Execute(l []inv.Item) {
	var err error

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	//c, err := chromedp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	err = runTasks(ctxt, c, l)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func runTasks(ctxt context.Context, c *chromedp.CDP, l []inv.Item) error {
	for _, item := range l {
		slug := "/" + item.Category + "/" + item.Slug
		if err := addToCart(ctxt, c, slug); err != nil {
			return err
		}
	}

	if err := checkout(ctxt, c); err != nil {
		return err
	}
	return nil
}

func addToCart(ctxt context.Context, c *chromedp.CDP, slug string) error {
	err := c.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(`https://www.supremenewyork.com/shop` + slug),
		// WaitVisible (?) lag issues
		chromedp.Click(`input[name="commit"]`),
	})
	if err != nil {
		return err
	}
	return nil
}

// Uses SendKeys to fill out form but might be able to get away with lighting fill.
// Not entirely sure what anti-bot constraints are in place.
func checkout(ctxt context.Context, c *chromedp.CDP) error {
	err := c.Run(ctxt, chromedp.Tasks{
		// chromedp.Click(`a[href="/shop/cart"]`, chromedp.NodeVisible), (?)
		chromedp.Click(`a[href="https://www.supremenewyork.com/checkout"]`, chromedp.NodeVisible),
		// BUG(SendKeys) - Types a bit too fast at times causing to typebehind characters.
		// Happens frequently with Telephone and CCNumber. Might just be an anti-bot measure.
		chromedp.SendKeys(`input[placeholder="name"]`, cfg.Name()),
		chromedp.SendKeys(`input[placeholder="email"]`, cfg.Email()),
		chromedp.SendKeys(`input[placeholder="tel"]`, cfg.Telephone()),
		chromedp.SendKeys(`input[placeholder="address"]`, cfg.Address()),
		chromedp.SendKeys(`input[placeholder="apt, unit, etc"]`, cfg.Unit()),
		chromedp.SendKeys(`input[placeholder="zip"]`, cfg.Zipcode()),

		/*
			Automatic fill from Zipcode
			chromedp.SendKeys(`input[placeholder="city"]`, cfg.City()),
			chromedp.Click(`select[name="order[billing_state]"]`, chromedp.NodeVisible),
			chromedp.SendKeys(`select[name="order[billing_state]"]`, cfg.State()+kb.Select),
			chromedp.Click(`select[name="order[billing_country]"]`, chromedp.NodeVisible),
		*/

		chromedp.SendKeys(`input[placeholder="number"]`, cfg.CCNumber()),
		chromedp.SendKeys(`input[placeholder="number"]`, cfg.CCNumber()),
		chromedp.Click(`select[name="credit_card[month]"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`select[name="credit_card[month]"]`, cfg.CCExpMonth()+kb.Select),
		chromedp.Click(`select[name="credit_card[year]"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`select[name="credit_card[year]"]`, cfg.CCExpYear()+kb.Select),
		chromedp.SendKeys(`input[placeholder="CVV"]`, cfg.CVV()),
		// Accept terms
		chromedp.Click(`#order_terms`, chromedp.NodeVisible),
		// Finalize Order
		chromedp.Click(`input[type="submit"]`, chromedp.NodeVisible),
		// Captcha is next
		chromedp.Sleep(150 * time.Second),
	})
	if err != nil {
		return err
	}
	return nil
}
