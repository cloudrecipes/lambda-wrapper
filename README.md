# Lambda Wrapper
This tool wraps libraries into lambda functions to be run on cloud. It creates
deployable `*.zip` file.

## Installing

## Usage

```
Usage: lambda-wrapper [options]


  Options:

    -V, --version             output the version number
    -c, --cloud <cloud>       cloud provider [AWS]
    -e, --engine <engine>     lambda function engine [node]
    -s, --source <source>     the source where to find library's code
    -n, --name <name>         the name of the library in the source   
    -o, --output <path>       path to output extended API file
    -h, --help                output usage information
```

All options might be set via `.lwrc.yaml` file.

Example of `.lwrc.yaml` file:
```yaml
cloud: AWS
engine: node
lib:
  source: npm
  name: @foo/bar
```

## Built With

## Authors
* [Anton Klimenko](https://github.com/antklim)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/cloudrecipes/lambda-wrapper/blob/master/LICENSE) file for details
