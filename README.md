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
    -s, --service <items>     a list of cloud services, the wrapper automatically
                              connects to these services and passes handlers to
                              the library
    -S, --source <source>     the source where to find library's code
    -N, --name <name>         the name of the library in the source   
    -o, --output <path>       path to save deployable lambda archive 
    -h, --help                output usage information
```

All options might be set via `.lwrc.yaml` file:
```yaml
cloud: AWS
engine: node
service:
  - s3
  - sqs
lib:
  source: npm
  name: @foo/bar
```

## Supported cloud providers
* AWS

## Supported engines
| Cloud | Engine |
| --- | --- |
| AWS | node |

## Supported services
| Cloud | Service |
| --- | --- |
| AWS | S3 |

## Supported library sources
* npm
* git

## Requirements for libraries
Library which should be wrapped into lambda should expose public `main` method.
This method will be used by wrapper as an entry point. 

### API for NodeJS lambda functions on AWS
__main(data, services)__
* `data` Object - event object passed by wrapper to lambda 
* `services` Object - object contains required service handlers such as `s3`, `sqs`, etc
* Returns: Promise - callback of lambda handler will be invoked when promise resolves. If promise rejected `err` object will be passed to the callback.

## Built With

## Authors
* [Anton Klimenko](https://github.com/antklim)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/cloudrecipes/lambda-wrapper/blob/master/LICENSE) file for details
