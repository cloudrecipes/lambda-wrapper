package wrapper

var injectLibraryIntoTemplateTestCases = []struct {
	template    string // template payload
	libraryName string // library name to inject into template
	expected    string // expected result
}{
	{
		"// library dependency\nconst handler = require('{{lib}}')",
		"@foo/bar",
		"// library dependency\nconst handler = require('@foo/bar')",
	},
}

var injectServicesIntoTemplateTestCases = []struct {
	template string   // template payload
	services []string // list of services to inject into template
	expected string   // expected result
}{
	{
		"// AWS SDK dependency\n{{aws}}\n// initiate required AWS services\n{{services}}",
		[]string{},
		"// AWS SDK dependency\n\n// initiate required AWS services\n",
	},
	{
		"// AWS SDK dependency\n{{aws}}\n// initiate required AWS services\n{{services}}",
		[]string{"s3", "sqs"},
		"// AWS SDK dependency\nconst aws = require('aws-sdk')\n// initiate required AWS services\nconst s3 = new aws.S3({apiVersion: 'latest'})",
	},
}

var initiateAwsHandlerTestCases = []struct {
	services []string // list of services to inject into template
	expected string   // expected result
}{
	{[]string{}, ""},
	{[]string{"s3", "sqs"}, "const aws = require('aws-sdk')"},
}

var initialeServiceHandlersTestCases = []struct {
	services []string // list of services to inject into template
	expected string   // expected result
}{
	{[]string{}, ""},
	{[]string{"s3", "sqs"}, "const s3 = new aws.S3({apiVersion: 'latest'})"},
}