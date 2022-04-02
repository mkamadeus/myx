package cli

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/mkamadeus/myx/pkg/config"
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/spec"
	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "myx",
		Short: "A machine learning code generation tool",
		Long:  "Proof-of-concept machine learning code generation tool.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("please supply the configuration file")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// setup
			config.Config = config.New()
			l, err := logger.New(config.Config.Log.Level, config.Config.Log.Output)
			if err != nil {
				panic(err)
			}
			logger.Logger = l

			specPath := args[0]
			b, err := ioutil.ReadFile(specPath)
			if err != nil {
				panic(err)
			}

			s, err := spec.Parse(b)
			if err != nil {
				panic(err)
			}

			code, err := generator.RenderSpec(s)
			if err != nil {
				panic(err)
			}

			// write spec
			f, err := os.Create(path.Join(config.Config.Output, "main.py"))
			if err != nil {
				panic(err)
			}

			defer f.Close()

			f.WriteString(code)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&config.Output, "output", "o", "./", "generated code output")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
