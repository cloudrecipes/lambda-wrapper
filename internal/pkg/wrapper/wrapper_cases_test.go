package wrapper_test

import (
	"errors"
	"os"
	"path"
)

var buildTemplateFileNameTestCases = []struct {
	cloud    string // cloud provider name
	engine   string // engine name
	expected string // expected result
}{
	{"aws", "node", "aws-node"},
}

var readTemplateFileTestCases = []struct {
	templatedir string
	filename    string
	err         error
	expected    string
}{
	{"", "no_such_template_file.txt", errors.New("open no_such_template_file.txt: no such file or directory"), ""},
	{
		path.Join(os.Getenv("GOPATH"), "src", "github.com", "cloudrecipes",
			"lambda-wrapper", "test", "fixtures"),
		"wrapper_readtemplatefile",
		nil,
		`// AWS SDK dependency
{{aws}}

// library dependency
const handler = require('{{lib}}')

const services = {}

// initiate required AWS services
{{services}}
`,
	},
}

var buildWrapperTestCases = []struct {
	template    string   // wrapper template
	libraryname string   // library name to inject to template
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
