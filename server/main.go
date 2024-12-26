package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Server Start")

	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("Warning: .env file not found")
		}
	}

	MongoDB_URL := os.Getenv("MONOGODB_URL")
	clientOptions := options.Client().ApplyURI(MongoDB_URL)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongo db atlas")

	collection = client.Database("react-go").Collection("todos")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                                                       // Allows all origins (no restrictions on the client making requests)
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",                       // Allow necessary HTTP methods
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Content-Length", // Allow common headers
		AllowCredentials: false,                                                     // No credentials are required (you can remove this line if you don't need it)
		ExposeHeaders:    "Content-Length,Content-Range",                            // Expose headers that the client can access
		MaxAge:           86400,                                                     // Cache pre-flight requests for 24 hours (optional)
	}))
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "4000"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	defer cursor.Close(context.Background())

	if err != nil {
		return err
	}
	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.Status(200).JSON(todos)

}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(401).JSON(fiber.Map{"error": "TODO body is empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid todo Id"})
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return c.Status(201).JSON(objectId)

}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Id"})
	}
	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(objectId)

}
