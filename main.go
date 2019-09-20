package debugger

import (

	"fmt"
	"time"
	
	hotconfig "github.com/jarzamendia/hotconfig"

)

//SendDebugMessage Verify a Env Variable, if true, send  a debug message.
func SendDebugMessage(msg string){

	isDebugEnableed := hotconfig.GetEnvVarOrDefault("DEBUG","false")

	currentTime := time.Now();

	if isDebugEnableed == "true" {

		debugMessage := currentTime.String() + " :: " + msg

		fmt.Println(debugMessage)

	}

}