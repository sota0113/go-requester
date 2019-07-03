package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"os"
	"time"
	"strings"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Rq_host	string	`envconfig:RQ_HOST`
	Rq_protocol	string	`envconfig:RQ_PROTOCOL`
	Rq_path	string	`envconfig:RQ_PATH`
	Rq_port	string	`envconfig:RQ_PORT`
}

func geturi () string {
	goenv := Env{}
	envconfig.Process("", &goenv)

	hostname := goenv.Rq_host
	protocol := goenv.Rq_protocol
	path := goenv.Rq_path
	port := goenv.Rq_port

	if len(protocol) != 0 {
		if !strings.Contains(protocol,"://") {
			protocol = protocol+"://"
		}
	} else {
		protocol = "http://"
	}

	if len(path) != 0 {
		if !strings.Contains(path,"/") {
			path = "/"+path
		}
	}

	if len(port) != 0 {
		if !strings.Contains(port,":") {
			port = ":"+port
		}
	} else {
		port = ":80"
	}
	uri := protocol+hostname+port+path
	return uri
}

func request (uri string) string {
	for {
		time.Sleep(1 * time.Second)
		resp,err := http.Get(uri)
	        if err != nil {
	                fmt.Printf("Get Error in requesting %s.\n",uri)
			continue
	        }
	        defer resp.Body.Close() //?
	        byteArray,err := ioutil.ReadAll(resp.Body)
	        if err != nil {
	                fmt.Println("Get Error in parsing received data.")
			continue
	        }
		fmt.Println(string(byteArray))
	}
}


func handler (w http.ResponseWriter ,r *http.Request ) {
	fmt.Fprintf(w, "up and running.")
}

func main () {
	uri := geturi()
	fmt.Printf("requesting %s.\n",uri)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3030",nil)
	request(uri)
}
