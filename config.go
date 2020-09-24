package main

type Config struct {
	ManifestRoot    string `required:"true"`
	StripPrefix     string `required:"true"`
	AssignmentsPath string `required:"true"`
	ListenAddr      string `required:"true"`
}
