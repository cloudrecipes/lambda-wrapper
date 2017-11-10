package wrapper_test

var buildTemplateFileNameTestCases = []struct {
	cloud    string // cloud provider name
	engine   string // engine name
	expected string // expected result
}{
	{"aws", "node", "aws-node"},
}

var buildWrapperTestCases = []struct {
	template    string   // wrapper template
	libraryName string   // library name to inject to template
	services    []string // services to initiate and inject to template
	expected    string   // expected result
}{
	{
		`// AWS SDK dependency
{{aws}}
// library dependency
const handler = require('{{lib}}')
const services = {}
// initiate required AWS services
{{services}}`,
		"@foo/bar",
		[]string{"s3"},
		`// AWS SDK dependency
const aws = require('aws-sdk')
// library dependency
const handler = require('@foo/bar')
const services = {}
// initiate required AWS services
services.s3 = new aws.S3({apiVersion: 'latest'})`,
	},
}
