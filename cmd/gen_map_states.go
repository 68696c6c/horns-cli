package cmd

import (
	"fmt"

	"github.com/68696c6c/goat"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(buildCommand)
}

var buildCommand = &cobra.Command{
	Use:   "gen:states path",
	Short: "Generate US Map State components.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		basePath := fmt.Sprintf("%s/states", args[0])
		println("Generating components to ", basePath)

		for _, state := range mapStatesData {
			fileName := nameToFileName(state.Name)
			template := templateStateComponentLarge
			if state.LabelBackground != nil {
				template = templateStateComponentSmall
			}
			err := generateFile(basePath, fileName, template, state)
			if err != nil {
				goat.ExitWithError(errors.Wrap(err, "failed to create file"))
			}
		}

		err := generateFile(basePath, "index.js", indexTemplate, "")
		if err != nil {
			goat.ExitWithError(errors.Wrap(err, "failed to create file"))
		}

		println("Done!")
		goat.ExitSuccess()
	},
}
