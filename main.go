package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KhashayarKhm/stamp/cmd"
	"github.com/spf13/cobra"
)

func main() {
  log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	root := &cobra.Command{Short:  "Watermark command", Version: "1.1.0"}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Watermark{}.Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute command:\n%v", err)
	}
}
