package cmd

import (
	"github.com/68696c6c/goat"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(buildCommand)
}

var buildCommand = &cobra.Command{
	Use:   "gen:states",
	Short: "Generate US Map State components.",
	Run: func(cmd *cobra.Command, args []string) {

		for _, state := range mapStatesData {
			fileName := nameToFileName(state.Name)
			template := templateStateComponentLarge
			if state.LabelBackground != nil {
				template = templateStateComponentSmall
			}
			err := generateFile("states", fileName, template, state)
			if err != nil {
				goat.ExitWithError(errors.Wrap(err, "failed to create file"))
			}
		}

		err := generateFile("states", "index.js", indexTemplate, "")
		if err != nil {
			goat.ExitWithError(errors.Wrap(err, "failed to create file"))
		}

		println("Done!")
		goat.ExitSuccess()
	},
}
