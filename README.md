# Lambda Wrapper
[![Build Status](https://travis-ci.org/cloudrecipes/lambda-wrapper.svg?branch=master)](https://travis-ci.org/cloudrecipes/lambda-wrapper)
[![Coverage Status](https://coveralls.io/repos/github/cloudrecipes/lambda-wrapper/badge.svg?branch=master)](https://coveralls.io/github/cloudrecipes/lambda-wrapper?branch=master)

This tool wraps libraries into lambda functions to be run on cloud. It creates
deployable `*.zip` file.

## Installing

## Usage

```
Usage: lambda-wrapper [options]


  Options:

    --cloud value, -c value      cloud provider name
    --engine value, -e value     lambda function engine
    --service value, -s value    a list of cloud services, the wrapper will automatically
                                 initiate handlers to these services and pass then to
                                 the library
    --libsource value, -S value  the source where to find library's code
    --libname value, -N value    the name of the library in the source
    --output value, -o value     path to save deployable lambda archive
    --test, -t                   flag to run library's unit tests before wrapping
                                 into lambda package
    --help, -h                   show help
    --version, -v                print the version
```

All options might be set via `.lwrc.yaml` file:
```yaml
cloud: AWS
engine: node
service:
  - s3
  - sqs
libsource: npm
libname: '@foo/bar'
output: lambda.zip
test: true
```
Configuration file should be located in the process working directory.

### Precedence
The precedence for option value sources is as follows (highest to lowest):
1. Command line option value from user
2. Configuration file (if found)

## Supported cloud providers
* AWS

## Supported engines
| Cloud | Engine |
| --- | --- |
| AWS | node |

## Supported services
| Cloud | Engine | Service |
| --- | --- | --- |
| AWS | node | S3 |
| AWS | node | SNS |

## Supported library sources
* npm
* git

_Note_: In case when git is a library source and engine is NodeJS then `package.json` __must__ have `main` property. The value of the field will be used as an entry point to your library.

Also `test` field of a `package.json` will be used as a flag to run unit tests prior wrapping library for lambda.

## Requirements for libraries
Library which should be wrapped into lambda should expose public `main` method.
This method will be used by wrapper as an entry point.

### API for NodeJS lambda functions on AWS
__main(data, services)__
* `data` Object - event object passed by wrapper to lambda
* `services` Object - object contains required service handlers such as `s3`, `sqs`, etc
* Returns: Promise - callback of lambda handler will be invoked when promise resolves. If promise rejected `err` object will be passed to the callback.

## Built With
* [project-layout](https://github.com/golang-standards/project-layout) - Standard Go Project Layout
* [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go

## Authors
* [Anton Klimenko](https://github.com/antklim)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/cloudrecipes/lambda-wrapper/blob/master/LICENSE) file for details
