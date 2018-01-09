package awswrapper

import (
	"errors"

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
	template    string
	opts        *options.Options
	payload     string
	readFileErr string
	err         error  // an error flag
	expected    string // expected result
}{
	{
		template:    "// library dependency\nconst handler = require('{{lib}}')",
		opts:        &options.Options{Output: "tmp"},
		payload:     "",
		readFileErr: "open package.json: no such file or directory",
		err:         errors.New("open package.json: no such file or directory"),
		expected:    "",
	},
	{
		template:    "// library dependency\nconst handler = require('{{lib}}')",
		opts:        &options.Options{Output: "tmp"},
		payload:     "{",
		readFileErr: "",
		err:         errors.New("unexpected end of JSON input"),
		expected:    "",
	},
	{
		template:    "// library dependency\nconst handler = require('{{lib}}')",
		opts:        &options.Options{Output: "tmp"},
		payload:     "{}",
		readFileErr: "",
		err:         errors.New("'Name' field is required in package.json"),
		expected:    "",
	},
	{
		template:    "// library dependency\nconst handler = require('{{lib}}')",
		opts:        &options.Options{Output: "tmp"},
		payload:     "{\"name\": \"test\", \"version\": \"0.0.1\", \"main\": \"go.js\"}",
		readFileErr: "",
		err:         nil,
		expected:    "// library dependency\nconst handler = require('./_git/go.js')",
	},
}
