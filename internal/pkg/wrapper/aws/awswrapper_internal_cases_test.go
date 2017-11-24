package awswrapper

var injectLibraryIntoTemplateTestCases = []struct {
	template    string // template payload
	libraryname string // library name to inject into template
	expected    string // expected result
}{
	{
		template:    "// library dependency\nconst handler = require('{{lib}}')",
		libraryname: "@foo/bar",
		expected:    "// library dependency\nconst handler = require('@foo/bar')",
	},
}

var injectServicesIntoTemplateTestCases = []struct {
	template string   // template payload
	services []string // list of services to inject into template
	expected string   // expected result
}{
	{
		template: "// AWS SDK dependency\n{{aws}}\n// initiate required AWS services\n{{services}}",
		services: []string{},
		expected: "// AWS SDK dependency\n\n// initiate required AWS services\n",
	},
	{
		template: "// AWS SDK dependency\n{{aws}}\n// initiate required AWS services\n{{services}}",
		services: []string{"s3", "sqs"},
		expected: "// AWS SDK dependency\nconst aws = require('aws-sdk')\n// initiate required AWS services\nservices.s3 = new aws.S3({apiVersion: 'latest'})\n",
	},
}

var initiateAwsHandlerTestCases = []struct {
	services []string // list of services to inject into template
	expected string   // expected result
}{
	{services: []string{}, expected: ""},
	{services: []string{"s3", "sqs"}, expected: "const aws = require('aws-sdk')"},
}

var initiateServiceHandlersTestCases = []struct {
	services []string // list of services to inject into template
	expected string   // expected result
}{
	{services: []string{}, expected: ""},
	{services: []string{"s3", "sqs", "sns"}, expected: "services.s3 = new aws.S3({apiVersion: 'latest'})\n\nservices.sns = new aws.SNS()"},
}
