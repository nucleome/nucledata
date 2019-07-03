package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/nimezhu/box"
	"github.com/nimezhu/data"

	"github.com/nimezhu/nbdata"
)

func startServer(wb *data.SimpleWorkbook, port int) (*box.Box, error) {
	customCors := ""
	corsOptions := nbdata.GetCors(customCors)
	//root, DIR
	root := nbdata.UserHomeDir() //TODO
	router := mux.NewRouter()
	/*
		if nbdata.GuessURIType(uri) == "gsheet" {
			dir := filepath.Join(root, DIR)
			ctx := context.Background()
			config := nbdata.GsheetConfig()
			gA := asheets.NewGAgent(dir)
			if !gA.HasCacheFile() {
				gA.GetClient(ctx, config)
			}
		}
	*/
	s := box.NewBox("Nucleome Data", root, DIR, VERSION)
	s.InitRouter(router)
	s.InitHome(root)
	idxRoot := s.InitIdxRoot(root) //???
	l := data.NewLoader(idxRoot)
	l.Plugins["tsv"] = nbdata.PluginTsv
	l.LoadWorkbook(wb, router)
	go s.StartLocalServer(port, router, &corsOptions)
	log.Println("Data Service Ready")

	return s, nil
}
