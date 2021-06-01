// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// // func getPosts() string {
// // 	fmt.Println("kkkkkkkk")
// // 	return "aaaa"
// // }

// func main() {

// 	app := fiber.New()

// 	app.Get("/api/", func(c *fiber.Ctx) error {
// 		return c.SendString("GET request")
// 	})

// 	// app.Post("/api/", func(c *fiber.Ctx) error {
// 	// 	return c.SendString("POST request")
// 	// })

// 	// app.Get("/api/list", func(c *fiber.Ctx) error {
// 	// 	posts := getPosts() // your logic
// 	// 	if len(posts) == 0 {
// 	// 		return c.Status(404).JSON(&fiber.Map{
// 	// 			"success": false,
// 	// 			"error":   "There are no posts!",
// 	// 		})
// 	// 	}
// 	// 	return c.JSON(&fiber.Map{
// 	// 		"success": true,
// 	// 		"data":    posts,
// 	// 	})
// 	// })

// 	// app.Get("/api/:param", func(c *fiber.Ctx) error {
// 	// 	return c.SendString("param: " + c.Params("param"))
// 	// })

// 	// // app.Get("/ws", websocket.New(func(c *websocket.Conn) {
// 	// // 	for {
// 	// // 		mt, msg, err := c.ReadMessage()
// 	// // 		if err != nil {
// 	// // 			log.Println("read:", err)
// 	// // 			break
// 	// // 		}
// 	// // 		log.Printf("recv: %s", msg)
// 	// // 		err = c.WriteMessage(mt, msg)
// 	// // 		if err != nil {
// 	// // 			log.Println("write:", err)
// 	// // 			break
// 	// // 		}
// 	// // 	}
// 	// // }))

// 	app.Static("/", "./public")

// 	app.Listen(":6001")
// }
package main

//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js,*.ico,*.png,*.json
import (
	"log"
	"net/http"

	_ "github.com/jlego/go_test/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":6003", nil)
}
