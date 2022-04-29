# `myx` - ML Experiment to Service utility tool

![itb](https://shields.io/badge/made%20for-ITB-blue?style=flat)
![goversion](https://shields.io/badge/version-1.18-blue?logo=go&style=flat)
![author](https://shields.io/badge/-mkamadeus-black?logo=github&style=flat)

### Introduction

Made to fulfill the requirements of graduation in Bandung Institute of Technology as an Informatics Engineer (2018).
This repository holds a proof-of-concept tool that converts complete ML experiments to a working code.
_Don't expect this project to be used in production_; this was really made to see if a certain tool can be useful and be adopted into a MLOps workflow.

### How It Works

- It accepts a specification file which contains information regarding what service should be made
- In the specification file, users can configure how to process their input before going to the model itself, since there usually exist an extra step in data preprocessing.
- It uses existing scalers/encoders and pretrained models that is used in the training process to make the process more streamlined.
- This tool generates Python code which are widely used in ML services. Currently based only on FastAPI.

### Running the Project

#### Prerequisites

- Go v1.18
- Python, pip for installing the requirements

#### Building and Running

```
$ make build
$ ./myx

Usage:
  myx [flags]

Flags:
  -h, --help            help for myx
  -o, --output string   generated code output (default "./")
  -v, --verbose         verbose output
```

### Specification Format

Some examples are generated in the `examples` directory.
In general, there are five main sections in the directory.

```yaml
input:
  format: <input-format>
  meta: <input-metadata>
output:
  - name: <output-name>
    type: <output-type>
pipeline:
  - module: <module-name>
    meta: <module-parameters>
model:
  format: <model-format>
  path: <model-path>
interface:
  - type: <interface-type>
    port: <interface-port>
```

Specification format is open to changes and additions.
Author designed this to be especially extensible.

### Author

[@mkamadeus](https://github.com/mkamadeus)
[13518035@std.stei.itb.ac.id](mailto:13518035@std.stei.itb.ac.id)
