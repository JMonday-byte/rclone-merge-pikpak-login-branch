// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	_ "github.com/JMonday-byte/rclone-merge-pikpak-login-branch/backend/all" // import all backends
	"github.com/JMonday-byte/rclone-merge-pikpak-login-branch/cmd"
	_ "github.com/JMonday-byte/rclone-merge-pikpak-login-branch/cmd/all"    // import all commands
	_ "github.com/JMonday-byte/rclone-merge-pikpak-login-branch/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
