package main

import (
	"fmt"
	"html/template"
	models "news-poolerr/Model"
	router "news-poolerr/Router"
	services "news-poolerr/Services"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TryCatch(f func()) func() error {
	return func() (err error) {
		defer func() {
			if panicInfo := recover(); panicInfo != nil {
				err = fmt.Errorf("%v, %s", panicInfo, string(debug.Stack()))
				return
			}
		}()
		f() // calling the decorated function
		return err
	}
}

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	AppEnv := viper.GetString("APP_ENV")
	fmt.Println(AppEnv)
	// initialize database
	models.Init()
}

func main() {

	r := gin.Default()
	//new template engine
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "index",
		Partials:     []string{"partials/ad"},
		DisableCache: true,
		Funcs: template.FuncMap{
			"siteUrl": func() string {
				return "http://" + viper.GetString("SITE_HOST") + ":" + viper.GetString("SITE_PORT")
			},
		},
	})

	// shedule every 1 hour
	go func() {
		scheduleTime , _ := strconv.Atoi(viper.GetString("SHEDULE_TIME"))

		defer func() {
			fmt.Println("finished")
		}()

		for true {
			links := services.GetLinks()
			for _, link := range links {

				services.FetchNews(link)
				fmt.Println("fetched finished for", link.Title)
			}
			
			time.Sleep(time.Duration(scheduleTime) * time.Second)
		}
	}()

	router.RegisterFrontendRoute(r)
	router.RegisterApisRoute(r)

	r.Static("/assets/css", "Views/assets/css")
	r.Static("/assets/js", "Views/assets/js")

	fmt.Println(viper.GetString("SITE_PORT"))

	r.Run(":" + viper.GetString("SITE_PORT"))
}
