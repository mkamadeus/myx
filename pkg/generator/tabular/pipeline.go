package tabular

import (
	"fmt"
	"sort"

	"github.com/mkamadeus/myx/pkg/generator"
	"github.com/mkamadeus/myx/pkg/generator/tabular/pipeline"
	"github.com/mkamadeus/myx/pkg/logger"
	"github.com/mkamadeus/myx/pkg/template/pipeline/tabular"
)

func (g *TabularGenerator) RenderPipelineSpec() (*generator.PipelineCode, error) {
	logger.Logger.Instance.Debug("running in tabular input mode")

	// map input in temporary buffer, save targets
	inputMapper := make([]pipeline.PipelineModule, 0)
	targetsMapper := make([]int, 0)
	imports := make([]string, 0)

	logger.Logger.Instance.Info("mapping input in temporary buffer")

	columns := g.Spec.Input.Metadata["columns"].([]interface{})
	for _, input := range columns {
		// if input is not preprocessed
		casted := input.(map[interface{}]interface{})
		if casted["preprocessed"] == nil || casted["preprocessed"] == false {
			logger.Logger.Instance.Debugf("direct input %v", casted)

			// make module for direct input
			module := &pipeline.DirectModule{
				Target: casted["target"].(int),
				Name:   casted["name"].(string),
			}
			inputMapper = append(inputMapper, module)
			targetsMapper = append(targetsMapper, module.Target)

		}
	}

	// for each module
	for _, p := range g.Spec.Pipeline {
		logger.Logger.Instance.Debug(p)

		// find the preprocessing module
		if p.Module == "preprocessing/scale" {
			// add module to mapper
			names := make([]string, 0)
			for _, n := range p.Metadata["for"].([]interface{}) {
				names = append(names, n.(string))
			}
			targets := make([]int, 0)
			for _, t := range p.Metadata["target"].([]interface{}) {
				targets = append(targets, t.(int))
			}

			module := &pipeline.ScaleModule{
				Names:   names,
				Targets: targets,
				Path:    p.Metadata["path"].(string),
			}

			imports = append(imports, "import joblib")
			inputMapper = append(inputMapper, module)
			targetsMapper = append(targetsMapper, module.Targets...)
		} else if p.Module == "preprocessing/onehot" {
			// add module to mapper
			values := make([]string, 0)
			for _, n := range p.Metadata["values"].([]interface{}) {
				values = append(values, n.(string))
			}
			targets := make([]int, 0)
			for _, t := range p.Metadata["target"].([]interface{}) {
				targets = append(targets, t.(int))
			}

			module := &pipeline.OneHotModule{
				Name:    p.Metadata["for"].(string),
				Targets: targets,
				Values:  values,
			}
			inputMapper = append(inputMapper, module)
			targetsMapper = append(targetsMapper, module.Targets...)
		} else if p.Module == "preprocessing/label" {
			// add module to mapper
			names := make([]string, 0)
			for _, n := range p.Metadata["for"].([]interface{}) {
				names = append(names, n.(string))
			}
			targets := make([]int, 0)
			for _, t := range p.Metadata["target"].([]interface{}) {
				targets = append(targets, t.(int))
			}

			module := &pipeline.LabelModule{
				Names:   names,
				Targets: targets,
				Path:    p.Metadata["path"].(string),
			}

			imports = append(imports, "import joblib")
			inputMapper = append(inputMapper, module)
			targetsMapper = append(targetsMapper, module.Targets...)
		} else {
			logger.Logger.Instance.Debug("module not found")
			return nil, fmt.Errorf("invalid module found")
		}
	}

	// map buffer to actual code
	logger.Logger.Instance.Info("mapping buffer to code")

	pipelineCodes := make([]string, 0)
	for _, mapper := range inputMapper {
		c, err := mapper.Run()
		if err != nil {
			return nil, err
		}

		pipelineCodes = append(pipelineCodes, c...)
		logger.Logger.Instance.Debug(pipelineCodes)
	}

	// aggregation code
	variables := make([]string, len(targetsMapper))
	sort.Ints(targetsMapper)
	for it, t := range targetsMapper {
		variables[it] = fmt.Sprintf("%d", t)
	}

	aggregationCode, err := tabular.GenerateTabularAggregationCode(&tabular.TabularAggregationValues{
		Variables: variables,
	})
	if err != nil {
		return nil, err
	}

	return &generator.PipelineCode{
		Imports:     imports,
		Pipelines:   pipelineCodes,
		Aggregation: aggregationCode,
	}, nil
}
