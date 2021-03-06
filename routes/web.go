package routes

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/controller"
	"net/http"
)

func initWebRoute()  {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/categories", controller.Categories)
	http.HandleFunc("/works", controller.Works)
	http.HandleFunc("/about", controller.About)
	//二级页面
	http.HandleFunc("/article", controller.Article)
	http.HandleFunc("/category", controller.CategoryArticle)
	http.HandleFunc( config.Cfg.DashboardEntrance, controller.Dashboard)
	//静态文件服务器
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("resources/public"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(config.Cfg.DocumentPath + "/assets"))))

}