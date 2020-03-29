package main

import (
	"bufio"
	"fmt"
	"github.com/flosch/pongo2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Target struct {
	Addr string
	Weight string
}

type Upstream struct {
	Name string
	Targets []Target
}

type Location struct {
	ProjectName string
	ProxyHost string
	Upstream string
}

type Server struct {
	Domain string
	Locations []Location
}

func main()  {
	var locations []Location
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(),",")
		if len(tmp) == 2 {
			local := Location {
				ProjectName:tmp[0],
				ProxyHost:tmp[1],
				Upstream:"cloud.gateway",
			}
			locations = append(locations,local)
		} else {
			fmt.Println("error data line :", tmp)
		}
	}

	ups := []Upstream {
		{ Name: "cloud.gateway",  Targets: []Target{ {Addr: "10.27.3.127:80", Weight: "10"}, }, },
	}

	server := []Server{
		{ Domain:"api.bdp.autohome.com.cn", Locations: locations },
	}
	for _,v := range server {
		var ser = pongo2.Must(pongo2.FromFile("server.tpl"))
		cont, err := ser.Execute(pongo2.Context{"upstreams": ups, "server": v})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			aa := strings.ReplaceAll(cont,"\n\n","\n")
			err := ioutil.WriteFile(v.Domain+".conf", []byte(aa), 0644)
			if err != nil {
				fmt.Println("parse error:" ,v.Domain)
			}
		}
	}
}
