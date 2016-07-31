A very simple logging package. 
It helps track down error to specific line in file. 
Can add valuable trace data to locate error to exact point.

package main

import (
  "errors"
	"github.com/vishnuone/clearlog"
)

func main() {
	clearlog.Info("Starting...")
	
	dataId := "10"
  err := foo(dataId)
	if err != nil {
	    clearlog.Error(err, "dataId:" + dataId)
  }	
	
	clearlog.Info("Stopping...")
}

func foo() error {
	return errors.New("error from foo func.")
}
