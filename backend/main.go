package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
  "log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	svix "github.com/svix/svix-webhooks/go"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	webhookSecret := os.Getenv("CLERK_WEBHOOK_SIGNING_SECRET")
	if webhookSecret == "" {
		log.Fatal("CLERK_WEBHOOK_SIGNING_SECRET is not set")
	}

	router := gin.Default()

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Failed to set websocket upgrade:", err)
			return
		}
		defer conn.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}

			name := string(message)
			response := fmt.Sprintf("Hello %s", name)

			err = conn.WriteMessage(messageType, []byte(response))
			if err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}
	})

	// Clerk webhook endpoint for `user.created` event
	// TODO: write user metadata to MongoDB
	// TODO: setup another webhook for `user.updated` event
	router.POST("/api/webhooks", func(c *gin.Context) {
		body, _ := c.GetRawData()

		headers := http.Header{}
		headers.Set("svix-id", c.GetHeader("svix-id"))
		headers.Set("svix-timestamp", c.GetHeader("svix-timestamp"))
		headers.Set("svix-signature", c.GetHeader("svix-signature"))

		webhook, _ := svix.NewWebhook(webhookSecret)
		if err := webhook.Verify(body, headers); err != nil {
			log.Printf("webhook: verification failed: %v", err)
			c.Status(http.StatusBadRequest)
			return
		}

		log.Printf("webhook: received payload: %s", string(body))
		c.Status(http.StatusOK)
	})

	router.Run(":8080")
}
