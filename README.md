# maxon
Golang библиотека для написания ботов [MAX](https://dev.max.ru/docs-api).

```golang
package main

import (
	"context"
	"fmt"

	"github.com/onrik/maxon"
)

func main() {
	ctx := context.Background()

	bot := maxon.New("{access_token}")
    me, err := bot.Me(ctx)
    if err != nil {
        fmt.Println(err)
        return
    }

	fmt.Println(me)
}


```