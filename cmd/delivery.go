package cmd

import (
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	werfcommon "github.com/werf/werf/cmd/werf/common"
	werfroot "github.com/werf/werf/cmd/werf/root"
	"github.com/werf/werf/pkg/storage"
)

func init() {
	storage.DefaultHttpSynchronizationServer = "https://delivery-sync.deckhouse.ru"

	ctx := werfcommon.GetContextWithLogger()

	werfRootCmd, err := werfroot.ConstructRootCmd(ctx)
	if err != nil {
		werfcommon.TerminateWithError(err.Error(), 1)
	}

	werfRootCmd.Use = "d"
	werfRootCmd.Aliases = []string{"delivery"}
	werfRootCmd = ReplaceCommandName("werf", "d8 d", werfRootCmd)
	werfRootCmd.Short = strings.Replace(werfRootCmd.Short, "werf", "d8 d", 1)
	werfRootCmd.Long = strings.Replace(werfRootCmd.Long, "werf", "d8 d", 1)
	werfRootCmd.Long = werfRootCmd.Long + `

LICENSE NOTE: The d8 delivery functionality is exclusively available to users holding a valid license for any commercial version of the Deckhouse Kubernetes Platform.

© Flant JSC 2024`

	removeKubectlCmd(werfRootCmd)

	rootCmd.AddCommand(werfRootCmd)
	rootCmd.SetContext(ctx)
}

func removeKubectlCmd(werfRootCmd *cobra.Command) {
	kubectlCmd, _ := lo.Must2(werfRootCmd.Find([]string{"kubectl"}))
	kubectlCmd.Hidden = true

	for _, cmd := range kubectlCmd.Commands() {
		kubectlCmd.RemoveCommand(cmd)
	}

	werfRootCmd.RemoveCommand(kubectlCmd)
}
