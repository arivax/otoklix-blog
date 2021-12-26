package http

import (
	log "github.com/sirupsen/logrus"

	fiber "github.com/gofiber/fiber/v2"

	"go-otoklix-blog/infra/logger"
	"go-otoklix-blog/usecase/blogpost"
)

// Routes ...
func BlogPostRoutes(app *fiber.App, group string) {
	api := app.Group(group)

	api.Get("/BlogPost/", FindAll)
	api.Get("/BlogPost/:id", FindById)
	api.Post("/BlogPost/", CreateOne)
	api.Put("/BlogPost/:id", UpdateOne)
	api.Delete("/BlogPost/:id", DeleteOne)
}

// FindAll BlogPost
func FindAll(c *fiber.Ctx) error {
	res, err := blogpost.Svc.FindAll()
	if err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "FindAll " + c.Method(), "user": "1"})
		c.Status(500).JSON(err)
		return err
	}
	c.Status(200).JSON(res)
	return nil
}

// FindById BlogPost
func FindById(c *fiber.Ctx) error {
	id := c.Params("id", "0")

	res, err := blogpost.Svc.FindById(id)
	if err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "FindById " + c.Method(), "user": "1"})
		c.Status(500).JSON(err)
		return err
	}
	c.Status(200).JSON(res)
	return nil
}

// CreateOne BlogPost
func CreateOne(c *fiber.Ctx) error {
	p := new(blogpost.BlogPostReq) //  Parse body into oBlogPost struct
	if err := c.BodyParser(p); err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "CreateOne " + c.Method(), "user": "1"})
		c.Status(400).JSON(err)
		return err
	}
	res, err := blogpost.Svc.CreateOne(p)
	if err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "CreateOne " + c.Method(), "user": "2"})
		c.Status(500).JSON(err)
		return err
	}
	c.Status(200).JSON(res)
	return nil
}

// UpdateOne BlogPost
func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")

	p := new(blogpost.BlogPostReq) //  Parse body into data struct
	if err := c.BodyParser(p); err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "UpdateOne " + c.Method(), "user": "1"})
		c.Status(400).JSON(err)
		return err
	}
	res, err := blogpost.Svc.UpdateOne(p, id)
	if err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "UpdateOne " + c.Method(), "user": "2"})
		c.Status(500).JSON(err)
		return err
	}
	c.Status(200).JSON(res)
	return nil
}

// DeleteOne BlogPost
func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")

	p := new(blogpost.BlogPostReq) //  Parse body into data struct
	if err := c.BodyParser(p); err != nil {
		logger.Error(err.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "DeleteOne " + c.Method(), "user": "1"})
		c.Status(400).JSON(err)
		return err
	}
	res, err1 := blogpost.Svc.DeleteOne(p, id)
	if err1 != nil {
		logger.Error(err1.Error(), log.Fields{"module": "BlogPost " + c.BaseURL() + c.Path(), "method": "DeleteOne " + c.Method(), "user": "2"})
		c.Status(500).JSON(err1)
		return err1
	}
	c.Status(200).JSON(res)
	return nil
}
