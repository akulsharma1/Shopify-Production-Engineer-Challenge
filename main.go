package main

import (
	"shopifyproductionengineertest/dbfunctions"
	"shopifyproductionengineertest/inventory"

	// "shopifyproductionengineertest/managerfunctions"
	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	client, err := dbfunctions.Init()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// add product

	app.Post("/api/v1/addproduct", func(c *fiber.Ctx) error {
		p := new(inventory.Product)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(405)
		}
		err := dbfunctions.AddProduct(client, *p)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})

	// add variant
	app.Post("/api/v1/addvariant", func(c *fiber.Ctx) error {
		p := new(inventory.FiberAddVariant)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(405)
		}
		variant := inventory.Variant {
			VariantID: p.VariantID,
			VariantName: p.VariantName,
			Size: p.Size,
			Quantity: p.Quantity,
			Price: p.Price,
		}
		err := dbfunctions.AddVariant(client, p.ProductID, variant)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})


	// edit product
	app.Post("/api/v1/editproduct", func(c *fiber.Ctx) error {
		p := new(inventory.Product)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(405)
		}
		err := dbfunctions.EditProduct(client, *p)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})

	// edit variant
	app.Post("/api/v1/editvariant", func(c *fiber.Ctx) error {
		p := new(inventory.FiberAddVariant)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(405)
		}
		variant := inventory.Variant {
			VariantID: p.VariantID,
			VariantName: p.VariantName,
			Size: p.Size,
			Quantity: p.Quantity,
			Price: p.Price,
		}
		err := dbfunctions.EditVariant(client, p.ProductID, variant)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})

	// delete product

	app.Post("/api/v1/deleteproduct", func(c *fiber.Ctx) error {
		p := new(inventory.FiberGetProductID)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(405)
		}
		err := dbfunctions.DeleteProduct(client, p.ProductID)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})

	// delete variant
	app.Post("/api/v1/deletevariant", func(c *fiber.Ctx) error {
		p := new(inventory.FiberGetVariantInfo)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(405)
		}
		err := dbfunctions.DeleteVariant(client, p.ProductID, p.VariantID)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	})

	// get product
	app.Get("/api/v1/getproducts", func(c *fiber.Ctx) error {
		data, err := dbfunctions.ListAllProducts(client)
		if err != nil {
			return c.SendStatus(405)
		}
		return c.JSON(data)
	})

	// test endpoint
	app.Get("/api/v1/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	app.Listen(":" + port)
	
}