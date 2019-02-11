// https://github.com/golang/go/wiki/SliceTricks
package main

import (
	"fmt"
	"strings"
)

type Application struct {
	Name             string `yaml:"name"`
	BootstrapRepoURL string `yaml:"bootstrap_repo_url"`
}

type Applications []Application

// Filter removes applications that not found in the list
func (a *Applications) Filter(appList string) {
	for _, filtered := range strings.Split(appList, ",") {
		if len(filtered) < 0 {
			break
		}
		for i, app := range *a {
			if !strings.EqualFold(app.Name, filtered) {
				*a = append((*a)[:i], (*a)[i+1:]...)
				break
			}
		}
	}
	return
}

func main() {
	apps := Applications{
		Application{Name: "app1"},
		Application{Name: "app2"},
	}

	fmt.Printf("before filter: %v\n", apps)
	apps.Filter("app1")
	fmt.Printf("after filter: %v\n", apps)
}
