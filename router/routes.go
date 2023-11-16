package router

import (
	"net/http"
	"url_shortner/constant"
	"url_shortner/controller"
)

var urlShortner = Routes{
	Route{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheURL},
	Route{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectURL},
}
