package atc

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func Execute() {
	var err error

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	err = runTasks(ctxt, c)
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

func runTasks(ctxt context.Context, c *chromedp.CDP) error {
	if err := openShop(ctxt, c); err != nil {
		return err
	}

	if err := addToCart(ctxt, c); err != nil {
		return err
	}

	if err := checkout(ctxt, c); err != nil {
		return err
	}
	return nil
}

func openShop(ctxt context.Context, c *chromedp.CDP) error {
	err := c.Run(ctxt, chromedp.Navigate(`https://www.supremenewyork.com/shop/accessories/pzm43dci0`))
	//err := c.Run(ctxt, chromedp.Navigate(`https://www.supremenewyork.com/shop`))
	if err != nil {
		return err
	}
	return nil
}

func addToCart(ctxt context.Context, c *chromedp.CDP) error {
	err := c.Run(ctxt, chromedp.Tasks{
		chromedp.WaitVisible(`input[type="submit"]`), // (?)
		chromedp.Click(`input[type="submit"]`, chromedp.NodeVisible),
	})
	if err != nil {
		return err
	}
	return nil
}

func checkout(ctxt context.Context, c *chromedp.CDP) error {
	err := c.Run(ctxt, chromedp.Tasks{
		chromedp.Click(`a[href="/shop/cart"]`, chromedp.NodeVisible),
		chromedp.Click(`a[href="https://www.supremenewyork.com/checkout"]`, chromedp.NodeVisible),
		// Name
		chromedp.SendKeys(`input[placeholder="name"]`, "Yung Boolean"),
		// Email
		chromedp.SendKeys(`input[placeholder="email"]`, "plz@end.me"),
		// Telephone
		chromedp.SendKeys(`input[placeholder="tel"]`, "0123456789"),
		// Address
		chromedp.SendKeys(`input[placeholder="address"]`, "101 main st."),
		// Apt/Unit
		chromedp.SendKeys(`input[placeholder="apt, unit, etc"]`, "1"),
		// ZipCode
		chromedp.SendKeys(`input[placeholder="zip"]`, "10293"),
		// City
		chromedp.SendKeys(`input[placeholder="city"]`, "gary"),
		// State (Abbrev.)
		chromedp.Click(`select[name="order[billing_state]"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`select[name="order[billing_state]"]`, "in"+kb.Select),
		/*
			Country
			chromedp.Click(`select[name="order[billing_country]"]`, chromedp.NodeVisible),
			chromedp.Click(`option[value="USA"]`, chromedp.NodeVisible),
		*/
		// CC Number
		chromedp.SendKeys(`input[placeholder="number"]`, "4242424242424242"),
		// CC Exp Month
		chromedp.Click(`select[name="credit_card[month]"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`select[name="credit_card[month]"]`, "12"+kb.Select),
		// CC Exp Year
		chromedp.Click(`select[name="credit_card[year]"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`select[name="credit_card[year]"]`, "2021"+kb.Select),
		// CVV
		chromedp.SendKeys(`input[placeholder="CVV"]`, "123"),
		// Accept terms
		chromedp.Click(`#order_terms`, chromedp.NodeVisible),
		// Finalize Order
		chromedp.Click(`input[type="submit"]`, chromedp.NodeVisible),
		chromedp.Sleep(150 * time.Second),
	})
	if err != nil {
		return err
	}
	return nil
}
