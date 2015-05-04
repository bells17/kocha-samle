package config

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/naoina/kocha"
	"github.com/naoina/kocha/log"
)

var (
	AppName   = "kocha-sample"
	AppConfig = &kocha.Config{
		Addr:          kocha.Getenv("KOCHA_ADDR", "127.0.0.1:9100"),
		AppPath:       rootPath,
		AppName:       AppName,
		DefaultLayout: "app",
		Template: &kocha.Template{
			PathInfo: kocha.TemplatePathInfo{
				Name: AppName,
				Paths: []string{
					filepath.Join(rootPath, "app", "view"),
				},
			},
			FuncMap: kocha.TemplateFuncMap{},
		},

		// Logger settings.
		Logger: &kocha.LoggerConfig{
			Writer:    os.Stdout,
			Formatter: &log.LTSVFormatter{},
			Level:     log.INFO,
		},

		// Middlewares.
		Middlewares: []kocha.Middleware{
			&kocha.RequestLoggingMiddleware{},
			&kocha.PanicRecoverMiddleware{},
			&kocha.FormMiddleware{},
			&kocha.SessionMiddleware{
				Name: "kocha-sample_session",
				Store: &kocha.SessionCookieStore{
					// AUTO-GENERATED Random keys. DO NOT EDIT.
					SecretKey:  "\b\xef(\xf3֯e\x12iV{\x80fޓ\x8e\x925^\xa0\xb3\xbd\x82\xee\x1c\xfdu\xb6ɯw\xc1",
					SigningKey: "hԴ\x80\ba'\u007fVY\xb6\x1fq\xf6'*",
				},

				// Expiration of session cookie, in seconds, from now.
				// Persistent if -1, For not specify, set 0.
				CookieExpires: time.Duration(90) * time.Hour * 24,

				// Expiration of session data, in seconds, from now.
				// Perssitent if -1, For not specify, set 0.
				SessionExpires: time.Duration(90) * time.Hour * 24,
				HttpOnly:       false,
			},
			&kocha.FlashMiddleware{},
			&kocha.DispatchMiddleware{},
		},

		MaxClientBodySize: 1024 * 1024 * 10, // 10MB
	}

	_, configFileName, _, _ = runtime.Caller(0)
	rootPath                = filepath.Dir(filepath.Join(configFileName, ".."))
)
