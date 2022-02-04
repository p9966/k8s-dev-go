package main

import (
	"net/http"
	"os"
	"text/template"
	"time"
)

type EnvInfo struct {
	HostName string
	Env      string
	Version  string
	Time     string
}

func main() {
	http.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})

	http.HandleFunc("/info", func(rw http.ResponseWriter, r *http.Request) {
		envInfo := getInfo()

		envTmp := `k8s-dev-myapp(go)
---------------------------------------
HostName: {{.HostName}}
Env:      {{.Env}}
Version:  {{.Version}}
Time:     {{.Time}}
---------------------------------------`
		tmp, err := template.New("report").Parse(envTmp)
		if err != nil {
			rw.Write([]byte("build tmp err: " + err.Error()))
			return
		}
		tmp.Execute(rw, envInfo)
	})

	http.ListenAndServe(":80", nil)
}

func getInfo() EnvInfo {
	envInfo := EnvInfo{
		Env:     os.Getenv("ENV"),
		Version: "v1.0",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	hostName, _ := os.Hostname()
	envInfo.HostName = hostName
	return envInfo
}
