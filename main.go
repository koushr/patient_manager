package main

import (
	"fmt"
	"github.com/corazawaf/coraza/v2"
	"github.com/corazawaf/coraza/v2/seclang"
)

func main() {
	// First we initialize our waf and our seclang parser
	waf := coraza.NewWaf()
	fmt.Println(1)
	parser, _ := seclang.NewParser(waf)
	fmt.Println(2)

	// Now we parse our rules
	if err := parser.FromString(`SecRule REMOTE_ADDR "@rx .*" "id:1,phase:1,deny,status:403"`); err != nil {
		fmt.Println(err)
	}

	// Then we create a transaction and assign some variables
	tx := waf.NewTransaction()
	fmt.Println(3)
	defer func() {
		tx.ProcessLogging()
		tx.Clean()
	}()
	tx.ProcessConnection("127.0.0.1", 8080, "127.0.0.1", 12345)
	fmt.Println(4)
	// Finally we process the request headers phase, which may return an interruption
	if it := tx.ProcessRequestHeaders(); it != nil {
		fmt.Printf("Transaction was interrupted with status %d\n", it.Status)
	}
}
