package main

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"runtime"
)

func getLocalIPs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	ips := make([]string, 0)
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip != nil && ip.IsGlobalUnicast() {
			ips = append(ips, ip.String())
		}
	}
	return ips
}

func main() {
	_, file, _, _ := runtime.Caller(0)
	webDataDir := filepath.Dir(file) + "/webdata"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, webDataDir+"/index.html")
		} else {
			http.FileServer(http.Dir(webDataDir)).ServeHTTP(w, r)
		}
	})


	ips := getLocalIPs()
for _, ip := range ips {
	fmt.Println("Server started on " + ip + ":8080")
}
http.ListenAndServe(ips[0]+":8080", nil)
}