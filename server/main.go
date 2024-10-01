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

// Run starts the app
// @title Skiza API
// @version 1.0
// @description Skiza is a Callback Listener and Pusher.
// @termsOfService http://swagger.io/terms/
// @contact.name Cozy
// @contact.url https://github.com/ItsCosmas
// @contact.email devcosmas@gmail.com
// @license.name MIT
// @license.url https://github.com/ItsCosmas/skiza/blob/master/LICENSE
// @host localhost:8000
// @BasePath /api/v1
func main() {
	// Print Signature to terminal
	fmt.Printf("\n%s\n%s\n%s\n\n", signatureBorders, signature, signatureBorders)
	app.Run()
}
