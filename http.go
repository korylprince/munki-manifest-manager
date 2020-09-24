package main

import (
	"log"
	"net/http"
	"os"
)

type Manager struct {
	fileSystem      http.FileSystem
	assignmentsPath string
}

func NewManager(config *Config) *Manager {
	return &Manager{fileSystem: http.Dir(config.ManifestRoot), assignmentsPath: config.AssignmentsPath}
}

func (m *Manager) GetManifest(w http.ResponseWriter, r *http.Request) {
	if _, err := m.fileSystem.Open(r.URL.Path); err == nil {
		http.FileServer(m.fileSystem).ServeHTTP(w, r)
		return
	} else if err != nil && !os.IsNotExist(err) {
		log.Println("Error: Unable to open file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	assign, err := NewAssignments(m.assignmentsPath)
	if err != nil {
		log.Println("Error: Unable to open assignments:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf, err := assign.DevicePlist(r.URL.Path[1:len(r.URL.Path)])
	if err != nil {
		log.Println("Error: Unable to generate plist:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(buf); err != nil {
		log.Println("Error: Unable to write response:", err)
	}
}
