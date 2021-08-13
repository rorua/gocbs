package main

import (
	"encoding/json"
	"gocbs/app/database"
	"gocbs/app/email"
	"gocbs/app/jsonconfig"
	"gocbs/app/recaptcha"
	"gocbs/app/server"
	"gocbs/app/session"
	"gocbs/app/view"
	"gocbs/app/view/plugin"
	"gocbs/route"
	"html/template"
	"log"
	"os"
	"runtime"
	"time"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Connect to database
	database.Connect(config.Database)

	// Configure the Google reCAPTCHA prior to loading view plugins
	recaptcha.Configure(config.Recaptcha)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(
		plugin.TagHelper(config.View),
		plugin.NoEscape(),
		plugin.PrettyTime(),
		PrettyDate(),
		recaptcha.Plugin())

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Info   `json:"Database"`
	Email     email.SMTPInfo  `json:"Email"`
	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
	Session   session.Session `json:"Session"`
	Template  view.Template   `json:"Template"`
	View      view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func PrettyDate() template.FuncMap {
	f := make(template.FuncMap)

	f["PRETTYDATE"] = func(t time.Time) string {
		return t.Format("01/02/2006")
	}

	return f
}
