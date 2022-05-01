package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/mkamadeus/myx/pkg/config"
	"github.com/mkamadeus/myx/pkg/executor"
	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/generator/image"
	"github.com/mkamadeus/myx/pkg/generator/tabular"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/models"
	"github.com/mkamadeus/myx/pkg/template/code"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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

			// parse yaml spec
			s := &models.MyxSpec{}
			err = yaml.Unmarshal(b, &s)
			if err != nil {
				panic(err)
			}

			var apiGenerator generator.Generator

			if s.Input.Format == "tabular" {
				apiGenerator = &tabular.TabularGenerator{
					Spec: s,
				}
			} else if s.Input.Format == "image" {
				apiGenerator = &image.ImageGenerator{
					Spec: s,
				}
			}

			logger.Logger.Instance.Info("generating API code")
			apiCode, err := apiGenerator.RenderCode()
			if err != nil {
				panic(err)
			}

			// get imports for executor
			logger.Logger.Instance.Info("getting required imports")
			imports := make([]string, 0)

			if s.Input.Format == "image" {
				imports = append(imports, "python-multipart")
			}

			codeLines := strings.Split(apiCode, "\n")
			r1 := regexp.MustCompile(`from (\w+) import .*`)
			r2 := regexp.MustCompile(`import (\w+)`)
			for _, code := range codeLines {
				if r1.Match([]byte(code)) || r2.Match([]byte(code)) {
					pkgName := strings.Split(code, " ")[1]
					if pkgName == "PIL" {
						imports = append(imports, "Pillow")
					} else if strings.Contains(pkgName, "keras") {
						imports = append(imports, "tensorflow")
					} else if pkgName != "io" {

						imports = append(imports, pkgName)
					}
				}
			}

			logger.Logger.Instance.Info("installing dependencies")
			pwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			e := &executor.Executor{
				Imports: imports,
				Path:    path.Join(pwd, config.Output),
			}
			if err := e.Execute(); err != nil {
				panic(err)
			}

			// write spec
			logger.Logger.Instance.Info("writing files")
			f, err := os.Create(path.Join(config.Config.Output, "main.py"))
			if err != nil {
				panic(err)
			}

			defer f.Close()

			f.WriteString(apiCode)

			// copy related files
			code.RenderFiles(config.Config.Output)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&config.Output, "output", "o", "./", "generated code output")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
