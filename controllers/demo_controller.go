package controllers

import (
	"net/http"
	"github.com/Radiobox/rest_codecs"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb"
	"log"
)

type DemoController struct {
	rest_codecs.BaseRestController
}

func (controller *DemoController) Path() string {
	return "/api/demo"
}

func (controller *DemoController) ReadMany(ctx context.Context) error {
	log.Print("Responding to request for item list.")
	messages := rest_codecs.NewMessageMap()
	return rest_codecs.Respond(ctx, http.StatusOK, messages, []interface{}{
		map[string]string{
			"description": "This is the first item",
			"link": controller.Path()+"/first",
		},
		map[string]string{
			"description": "This is the second item",
			"link": controller.Path()+"/second",
		},
	})
}

func (controller *DemoController) Read(idString string, ctx context.Context) error {
	log.Print("Responding to request for item details")
	messages := rest_codecs.NewMessageMap()
	messages.AddInfoMessage("Found item for id " + idString)
	return rest_codecs.Respond(ctx, http.StatusOK, messages, map[string]interface{}{
		"description": "This is the "+idString+" item",
		"name": idString,
		"details": "I am some details that are not very detailed because this is a demo",
	})
}

func MapControllers() {
	goweb.MapController(new(DemoController))
}
