package main

/**
 * @file main.go
 * @author Original author: Free Code Camp (Akhil Sharma)
 *		   Changes made by 0xChristopher for learning purposes
 *
 * This simple Slack Age Bot implementation initializes a slack bot to a workspace, takes specific input,
 * and responds to a user after processing their request. In this case, the bot takes a user's year of birth
 * as input and responds with their approximate age based on that and the current year.
 */

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

// The PrintCommandEvents() function formats bot logging to the console
func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Set environment variables
	godotenv.Load(".env")

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")

	os.Setenv("SLACK_BOT_TOKEN", slackBotToken)
	os.Setenv("SLACK_APP_TOKEN", slackAppToken)

	// Define bot instance and setup go routine
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())

	// Define bot command to invoke bot response
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}

			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	// Context for incoming requests
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set bot listener in current context
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
