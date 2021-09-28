package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"runtime"
)

type App struct {
}

type UserApp interface {
	Initialize()
}

func (u App) GetConfig() (map[string]interface{}, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("error getting config")
	}
	folder := path.Dir(filename)
	f, err := ioutil.ReadFile(fmt.Sprintf("%s/apps.yaml", folder))
	if err != nil {
		return nil, err
	}

	conf := map[string]interface{}{}
	if err := yaml.Unmarshal(f, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}
