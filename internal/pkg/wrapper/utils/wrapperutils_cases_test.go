package wrapperutils_test

import (
	"errors"
	"os"
	"path"
)

var templateFileNameTestCases = []struct {
	cloud    string // cloud provider name
	engine   string // engine name
	expected string // expected result
}{
	{"aws", "node", "aws-node"},
}

var readTemplateFileTestCases = []struct {
	templatedir string // directory where templates live
	filename    string // file name to read
	err         error  // expected error
	expected    string // expected result
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
