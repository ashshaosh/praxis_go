package main

import (
	"encoding/json"
	"fmt"
	"os"

	"net/http"

	"github.com/kylelemons/go-gypsy/yaml"
	"gopkg.in/gcfg.v1"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	//parseJSONFile()
	//parseYAMLFile()
	//parsINIFile()
	parseEnvVar()
}

func parseEnvVar() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	//os.Exit(0)
	fmt.Println(req)
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage")

}

func parsINIFile() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		fmt.Printf("Failed to parse %s", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}

func parseYAMLFile() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}

func parseJSONFile() {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
