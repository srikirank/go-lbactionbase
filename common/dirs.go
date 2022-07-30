package common

import (
	"log"
	"os"
	"path/filepath"
)

// LBAppSupportDir Launchbar app for Mac if installed creates a directory under UserHome > Library > Application Support
// with the name "Launchbar"
func LBAppSupportDir() string {
	d, e := os.UserHomeDir()
	if e != nil {
		log.Println("There was an error getting the User home directory. Returning empty string")
		return ""
	}
	return filepath.Join(d, "Library", "Application Support", "Launchbar")
}

// LBCustomActionsDir "Actions" dir under LBAppSupportDir
func LBCustomActionsDir() string {
	return filepath.Join(LBAppSupportDir(), "Actions")
}

// LBCustomActionsSharedScriptsDir "SharedScripts" dir under LBCustomActionsDir
func LBCustomActionsSharedScriptsDir() string {
	return filepath.Join(LBCustomActionsDir(), "SharedScripts")
}

func LBCustomActionsCacheDir() string {
	return filepath.Join(LBCustomActionsDir(), "Cache")
}
