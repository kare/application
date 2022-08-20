package application

// Meta contains application version information.
type Meta struct {
	BuildTime string `json:"build_time"`
	Version   string `json:"version"`
	Revision  string `json:"revision"`
}
