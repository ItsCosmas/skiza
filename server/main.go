package main

import (
	"fmt"
	app "skiza/api"
	"strings"
)

var signatureBorders = strings.Repeat("-", 56)

// http://patorjk.com/software/taag/#p=display&c=c&f=Graceful&t=Cozy
var signature = `
	   ____  __ _  __  ____   __  
	  / ___)(  / )(  )(__  ) / _\ 
	  \___ \ )  (  )(  / _/ /    \
	  (____/(__\_)(__)(____)\_/\_/

	🚀 Status: Skiza is now listening ...
`

func main() {
	// Print Signature to terminal
	fmt.Printf("\n%s\n%s\n%s\n\n", signatureBorders, signature, signatureBorders)
	app.Run()
}
