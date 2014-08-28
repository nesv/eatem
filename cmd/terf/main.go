package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/hashicorp/hcl"
)

var (
	Templates *template.Template
)

type (
	Config map[string]ConfigProfile

	ConfigProfile struct {
		Terf TerfConfigBlock `json:"terf" hcl:"terf"`
	}

	TerfConfigBlock struct {
		DatabaseURL string `json:"database_url" hcl:"database_url"`
		ListenAddr  string `json:"listen_addr" hcl:"listen_addr"`
		TemplateDir string `json:"template_dir" hcl:"template_dir"`
		StaticDir   string `json:"static_dir" hcl:"static_dir"`
	}
)

func main() {
	configFile := flag.String("config", "terf.conf", "path to the terf configuration file")
	configProfile := flag.String("profile", "development", "configuration profile to use")
	flag.Parse()

	// Load the configuration file.
	var config Config
	p, err := ioutil.ReadFile(*configFile)
	if err != nil {
		glog.Fatalln("failed to read config file:", err)
	}
	if err := hcl.Decode(&config, string(p)); err != nil {
		glog.Fatalln("failed to decode config file:", err)
	}
	profile, ok := config[*configProfile]
	if !ok {
		glog.Fatalf("no such config profile %q", *configProfile)
	}

	glog.Infoln("connecting to database")

	tplGlob := filepath.Join(profile.Terf.TemplateDir, "*.html")
	glog.Infof("loading templates %q", tplGlob)
	Templates, err = template.ParseGlob(tplGlob)
	if err != nil {
		glog.Fatalln(err)
	}

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(profile.Terf.StaticDir))))
	http.HandleFunc("/", handleIndex)

	glog.Infoln("listening on", profile.Terf.ListenAddr)
	if err := http.ListenAndServe(profile.Terf.ListenAddr, nil); err != nil {
		glog.Fatalln(err)
	}
	glog.Infoln("terminated gracefully")
}
