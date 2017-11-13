package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addrs, err := net.InterfaceAddrs()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var ips []string

		for _, address := range addrs {

			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips = append(ips, ipnet.IP.String())
				}

			}
		}
		ipStr := strings.Join(ips, " ")
		fmt.Fprintf(os.Stdout, "my ips is:%s\n", ipStr)
		fmt.Fprintf(w, "my ips is:%s\n", ipStr)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
