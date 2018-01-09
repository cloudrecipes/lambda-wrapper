package awswrapper

import (
	"github.com/cloudrecipes/lambda-wrapper/internal/pkg/options"
)

var injectLibraryIntoTemplateTestCases = []struct {
	template string // template payload
	opts     *options.Options
	expected string // expected result
}{
	{
		template: "// library dependency\nconst handler = require('{{lib}}')",
		opts:     &options.Options{LibName: "@foo/bar", LibSource: "npm"},
		expected: "// library dependency\nconst handler = require('@foo/bar')",
	},
	{
		template: "// library dependency\nconst handler = require('{{lib}}')",
		opts:     &options.Options{LibName: "https://github.com/cloudrecipes/aws-lambda-greeter.git", LibSource: "git"},
		expected: "// library dependency\nconst handler = require('./_git')",
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

var injectGitLibraryIntoTemplateTestCases = []struct {
	// filename string // filename to read payload
	// err      error  // an error flag
	// expected string // expected file payload
}{
// {
// 	filename: "no_such_file.txt",
// 	err:      errors.New("open no_such_file.txt: no such file or directory"),
// 	expected: "",
// },
// {
// 	filename: path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
// 		"lambda-wrapper", "test", "fixtures", "fs_readfile.txt"),
// 	err:      nil,
// 	expected: "Hello Test!",
// },
}
