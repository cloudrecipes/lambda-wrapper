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
