package g

var (
	AppName    = "unknown app name"
	AppVersion = "unknown app version"
	GitCommit  = "unknown git commit"
)

// UpdateLDFlags updates the variables set when compile
// example: go build -ldflags "-X 'main.appName=wax-engine' -X 'main.gitCommit=110'"
func UpdateLDFlags(appName, appVersion, gitCommit string) {
	// update app name
	if appName != "" {
		AppName = appName
	}

	// update app version
	if appVersion != "" {
		AppVersion = appVersion
	}

	// update git version
	if gitCommit != "" {
		GitCommit = gitCommit
	}
}
