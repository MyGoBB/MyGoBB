package constants

import "os"

// Version is the current wings version.
const Version = "0.0.1-alpha"

/* ---------- PATHS ---------- */

// DefaultFilePerms are the file perms used for created files.
const DefaultFilePerms os.FileMode = 0644

// DefaultFolderPerms are the file perms used for created folders.
const DefaultFolderPerms os.FileMode = 0744

// UploadsPath is the path of data for user uploads
const UploadsPath = "uploads"