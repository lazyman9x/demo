package config

import (
	"path/filepath"
	"strings"

	"github.com/demo/server/v2/mode"
	"github.com/gotify/configor"
)

// Configuration is stuff that can be configured externally per env variables or config file (config.yml).
type Configuration struct {
	Server struct {
		ListenAddr string `default:""`
		Port       int    `default:"80"`
		SSL        struct {
			Enabled         *bool  `default:"false"`
			RedirectToHTTPS *bool  `default:"true"`
			ListenAddr      string `default:""`
			Port            int    `default:"443"`
			CertFile        string `default:""`
			CertKey         string `default:""`
			LetsEncrypt     struct {
				Enabled   *bool  `default:"false"`
				AcceptTOS *bool  `default:"false"`
				Cache     string `default:"data/certs"`
				Hosts     []string
			}
		}
		ResponseHeaders map[string]string
		Stream          struct {
			AllowedOrigins []string
		}
		Cors struct {
			AllowOrigins []string
			AllowMethods []string
			AllowHeaders []string
		}
	}
	Database struct {
		Dialect    string `default:"sqlite3"`
		Connection string `default:"data/gotify.db"`
	}
	DefaultUser struct {
		Name string `default:"vietdh"`
		Pass string `default:"vietdh"`
	}
	PassStrength      int    `default:"10"`
	UploadedImagesDir string `default:"data/images"`
	PluginsDir        string `default:"data/plugins"`
}

func configFiles() []string {
	if mode.Get() == mode.TestDev {
		return []string{"config.yml"}
	}
	return []string{"config.yml", "/etc/gotify/config.yml"}
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{EnvironmentPrefix: "DEMO GIN"}).Load(conf, configFiles()...)
	if err != nil {
		panic(err)
	}
	addTrailingSlashToPaths(conf)
	return conf
}

func addTrailingSlashToPaths(conf *Configuration) {
	if !strings.HasSuffix(conf.UploadedImagesDir, "/") && !strings.HasSuffix(conf.UploadedImagesDir, "\\") {
		conf.UploadedImagesDir = conf.UploadedImagesDir + string(filepath.Separator)
	}
}
