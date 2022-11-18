package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)

	}
}


func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4394866151444-4405654759841-XGtiABf01lHN1n9Cnj3CPqR6")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04BG1B94RK-4405632659793-5826806912af68bf8f36708532e075c17a2a92cd3e9438ade2dd874efedd3e17")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())

	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, r slacker.Request, w slacker.ResponseWriter){
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				fmt.Println("Error", err)
				return
			}
			age := 2022-yob
			reply := fmt.Sprintf("age is %d", age)
			w.Reply(reply)
		},

	})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
