package main

import (
    "log"
    "os"

    tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
    var (
        port      = os.Getenv("PORT")
        publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
        token     = os.Getenv("TOKEN")      // you must add it to your config vars
    )

    webhook := &tb.Webhook{
        Listen:   ":" + port,
        Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
    }

    pref := tb.Settings{
        Token:  token,
        Poller: webhook,
    }

    b, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }

    inlineBtn := tb.InlineButton{
        Unique: "iasmin",
        Text:   "CONFIRA AQUI A RESPOSTA",
    }

    b.Handle(&inlineBtn, func(c *tb.Callback) {
        // necessário para o funcionamento
        b.Respond(c, &tb.CallbackResponse{
            ShowAlert: false,
        })

        b.Send(c.Sender, "iasmin totosa.")
    })

    inlineKeys := [][]tb.InlineButton{
        []tb.InlineButton{inlineBtn},
    }

    b.Handle("/iasmin", func(m *tb.Message) {
        b.Send(
            m.Sender,
            "Iasmin é oq?",
            &tb.ReplyMarkup{InlineKeyboard: inlineKeys})
    })

    b.Start()
}
