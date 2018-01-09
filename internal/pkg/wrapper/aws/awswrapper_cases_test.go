package awswrapper_test

import (
	"errors"

	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

var template = `// AWS SDK dependency
{{aws}}
// library dependency
const handler = require('{{lib}}')
const services = {}
// initiate required AWS services
{{services}}`

var wrapperTestCases = []struct {
	template string // wrapper template
	opts     *options.Options
	expected string // expected result
	err      error
}{
	{
		template: template,
		opts:     &options.Options{LibName: "@foo/bar", Services: []string{"s3"}},
		expected: `// AWS SDK dependency
const aws = require('aws-sdk')
// library dependency
const handler = require('@foo/bar')
const services = {}
// initiate required AWS services
services.s3 = new aws.S3({apiVersion: 'latest'})`,
		err: nil,
	},
	{
		template: template,
		opts: &options.Options{
			LibName:   "https://github.com/cloudrecipes/aws-lambda-greeter.git",
			LibSource: "git",
		},
		expected: "",
		err:      errors.New("open .lwtmp/lib/_git/package.json: no such file or directory"),
	},
}
