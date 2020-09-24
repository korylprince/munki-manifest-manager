package main

import (
	"howett.net/plist"
)

func EncodePlist(catalogs, manifests []string) ([]byte, error) {
	type data struct {
		Catalogs          []string `plist:"catalogs"`
		IncludedManifests []string `plist:"included_manifests"`
	}

	d := &data{Catalogs: catalogs, IncludedManifests: manifests}
	return plist.MarshalIndent(d, plist.XMLFormat, "\t")
}
