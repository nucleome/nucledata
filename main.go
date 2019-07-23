package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"

	"github.com/nimezhu/box"
	"github.com/nimezhu/data"
	"github.com/zserge/lorca"
)

const (
	VERSION = "0.0.3"
	DIR     = ".nucle"
)

type nbRunner struct {
	sync.Mutex
	data *data.SimpleWorkbook
	box  *box.Box
}

func (c *nbRunner) GetJson(a string) error {
	c.Lock()
	defer c.Unlock()
	log.Println("Get Data")
	err := json.Unmarshal([]byte(a), &c.data)
	if err != nil {
		return err
	}
	return nil
}
func (c *nbRunner) Run() {
	c.Lock()
	defer c.Unlock()
	c.box, _ = startServer(c.data, 8611)
	log.Println("Data Service is Ready")
}
func (c *nbRunner) Stop() {
	c.Lock()
	defer c.Unlock()
	c.box.Stop()
	log.Println("Reset")
}

type iWriter struct {
	out io.Writer
	f   func([]byte) (int, error)
}

func (c *iWriter) Write(p []byte) (int, error) {
	n, e := c.f(p)
	return n, e
}
func logFuncFactory(ui lorca.UI) func([]byte) (int, error) {
	return func(s []byte) (int, error) {
		k := fmt.Sprintf(`nbLog("%s")`, strings.Trim(string(s), "\n"))
		v := ui.Eval(k)
		n := len(s)
		return n, v.Err()
	}
}
func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 800, 600, args...)
	iLog := logFuncFactory(ui)
	writer := iWriter{os.Stdout, iLog}
	log.SetOutput(&writer)

	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	ui.Bind("start", func() {
		log.Println("UI is ready")
	})
	// Create and bind Go object to the UI
	c := &nbRunner{}
	ui.Bind("sendJson", c.GetJson)
	ui.Bind("nbRun", c.Run)
	ui.Bind("nbStop", c.Stop)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, BindataServer("www"))

	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("Exiting...")
}
