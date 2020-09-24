package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Defaults struct {
	Catalogs  []string `yaml:"catalogs"`
	Manifests []string `yaml:"manifests"`
}

type Device struct {
	Name      string   `yaml:"name"`
	Serial    string   `yaml:"serial"`
	Catalogs  []string `yaml:"catalogs"`
	Manifests []string `yaml:"manifests"`
}

type Assignments struct {
	*Defaults `yaml:"default"`
	Devices   []*Device `yaml:"devices"`
}

func NewAssignments(path string) (*Assignments, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open %s: %v", path, err)
	}
	defer f.Close()

	config := new(Assignments)
	d := yaml.NewDecoder(f)
	if err = d.Decode(config); err != nil {
		return nil, fmt.Errorf("Unable to decode config: %v", err)
	}

	return config, nil
}

func (a *Assignments) DevicePlist(serial string) ([]byte, error) {
	var catalogs []string
	catalogSet := make(map[string]struct{})
	var manifests []string
	manifestSet := make(map[string]struct{})

	for _, d := range a.Devices {
		if d.Serial == serial {
			for _, c := range d.Catalogs {
				if _, ok := catalogSet[c]; !ok {
					catalogs = append(catalogs, c)
					catalogSet[c] = struct{}{}
				}
			}
			for _, m := range d.Manifests {
				if _, ok := manifestSet[m]; !ok {
					manifests = append(manifests, m)
					manifestSet[m] = struct{}{}
				}
			}
			break
		}
	}

	for _, c := range a.Defaults.Catalogs {
		if _, ok := catalogSet[c]; !ok {
			catalogs = append(catalogs, c)
			catalogSet[c] = struct{}{}
		}
	}
	for _, m := range a.Defaults.Manifests {
		if _, ok := manifestSet[m]; !ok {
			manifests = append(manifests, m)
			manifestSet[m] = struct{}{}
		}
	}

	return EncodePlist(catalogs, manifests)
}
