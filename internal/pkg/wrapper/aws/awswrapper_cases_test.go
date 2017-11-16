package awswrapper_test

import o "github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"

var wrapperTestCases = []struct {
	template string // wrapper template
	options  *o.Options
	expected string // expected result
}{
	{
		`// AWS SDK dependency
{{aws}}
// library dependency
const handler = require('{{lib}}')
const services = {}
// initiate required AWS services
{{services}}`,
		&o.Options{LibName: "@foo/bar", Services: []string{"s3"}},
		`// AWS SDK dependency
const aws = require('aws-sdk')
// library dependency
const handler = require('@foo/bar')
const services = {}
// initiate required AWS services
services.s3 = new aws.S3({apiVersion: 'latest'})`,
	},
}
