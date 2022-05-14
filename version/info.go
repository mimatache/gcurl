package version

import (
	"fmt"
	"runtime"
	"time"
)

var (
	version    = "dev"
	commitHash = "dev"
	buildDate  = time.Now().Format(time.RFC3339)
)

// Version contains information about the running application such as name and version
type Version struct {
	Version   string `json:"version"`
	Hash      string `json:"hash"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
}

// AppInfo transforms application information into a struct
func AppInfo() Version {
	app := Version{
		Version:   version,
		Hash:      commitHash,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
	}
	return app
}

func (v Version) String() string {
	return fmt.Sprintf(
		`
Version:    %s
Hash:       %s
Build Date: %s
Go Version: %s
		`, v.Version, v.Hash, v.BuildDate, v.GoVersion,
	)
}
