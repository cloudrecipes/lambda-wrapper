package wrapper_test

var buildTemplateFileNameTestCases = []struct {
	cloud    string // cloud provider name
	engine   string // engine name
	expected string // expected result
}{
	{"aws", "node", "aws-node"},
}
