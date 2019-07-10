package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/nimezhu/box"
	"github.com/nimezhu/data"

	"github.com/nimezhu/nbdata"
)

func mkdir(p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.Mkdir(p, os.ModePerm)
	}
}

func startServer(wb *data.SimpleWorkbook, port int) (*box.Box, error) {
	customCors := ""
	corsOptions := nbdata.GetCors(customCors)
	root := filepath.Join(nbdata.UserHomeDir(), DIR)
	mkdir(root)
	s := box.NewBox("NucleData", VERSION).CorsOptions(&corsOptions).Port(port)
	router := s.GetRouter()
	idxRoot := filepath.Join(root, "index")
	mkdir(idxRoot)
	l := data.NewLoader(idxRoot)
	l.Plugins["tsv"] = nbdata.PluginTsv
	l.LoadWorkbook(wb, router)
	go s.Start("local")
	log.Println("Data Service Ready")
	return s, nil
}
