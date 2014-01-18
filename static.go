package main

import (
	"github.com/stretchr/goweb"
	"github.com/nelsam/goweb_angular_demo/settings"
)

func MapStaticFiles() {
	goweb.MapStatic("/js", settings.ProjectPath+"/public/js")
	goweb.MapStatic("/less", settings.ProjectPath+"/public/less")
	goweb.MapStatic("/partials", settings.ProjectPath+"/public/partials")
	goweb.MapStaticFile("***", settings.ProjectPath+"/public/app.html")
}
