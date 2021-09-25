package main

import (
	"fmt"
	"html/template"
	models "news-poolerr/Model"
	router "news-poolerr/Router"
	"runtime/debug"
	"strings"

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
	viper.SetEnvPrefix("news-poolerr")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/news-poolerr/")
	viper.AddConfigPath("$HOME/.news-poolerr")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

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

	router.RegisterFrontendRoute(r)
	router.RegisterApisRoute(r)

	r.Static("/assets/css", "Views/assets/css")
	r.Static("/assets/js", "Views/assets/js")

	fmt.Println(viper.GetString("SITE_PORT"))

	r.Run(":" + viper.GetString("SITE_PORT"))
}
