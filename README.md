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

	for {
		updates, err := bot.Updates(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, u := range updates {
			if u.Type == maxon.UpdateTypeMessageCreated {
				bot.MessageSend(ctx, u.Message.ChatID(), u.Message.Text(), nil)
			}
		}
	}
}

```