A very simple logging package. 

It helps track down the error to a specific line in the file. 
Can add valuable trace data to locate error to exact point too.
Because the log is tab separated we can use excel or other spreadsheet programs to view it. We can also parse it easily for use with analysis tools.


```
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

func foo(dataId string) error {
	return errors.New("error from foo func.")
}
```


Log : 

```
2016/07/31 11:03:23 	info	Starting...
2016/07/31 11:03:23 	error	/Users/jaymac/golang/src/github.com/vishnuone/demo.go:14	error from foo func.	dataId:10
2016/07/31 11:03:23 	info	Stopping...
```
