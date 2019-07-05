package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"time"
	"strings"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Rq_host	string	`envconfig:RQ_HOST`
	Rq_protocol	string	`envconfig:RQ_PROTOCOL`
	Rq_path	string	`envconfig:RQ_PATH`
	Rq_port	string	`envconfig:RQ_PORT`
	Rp_port	string	`envconfig:RP_PORT`
}

func geturi (hostname,protocol,path,port string) string {

        if len(hostname) == 0 {
		fmt.Println("error: environment variable RQ_HOST is empty.")
		os.Exit(1)
        }

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
		fmt.Printf("[ ENDPOINT ] %s\n",uri)
		fmt.Printf("[ RESULT   ] %s \n",string(byteArray))
	}
}

func healthcheck (resport string) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {fmt.Fprintf(w, "{'status':'up'}")})
	http.ListenAndServe(resport, nil)
}

func main () {
        goenv := Env{}
        envconfig.Process("", &goenv)
        hostname := goenv.Rq_host
        protocol := goenv.Rq_protocol
        path := goenv.Rq_path
        reqport := goenv.Rq_port
        resport := goenv.Rp_port

	uri := geturi(hostname,protocol,path,reqport)
	fmt.Printf("requesting %s.\n",uri)

	go request(uri)
	healthcheck(":"+resport)

}
