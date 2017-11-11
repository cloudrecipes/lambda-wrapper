package wrapper

import "testing"

func TestInjectLibraryIntoTemplate(t *testing.T) {
	for _, test := range injectLibraryIntoTemplateTestCases {
		actual := injectLibraryIntoTemplate(test.template, test.libraryname)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestInjectServicesIntoTemplate(t *testing.T) {
	for _, test := range injectServicesIntoTemplateTestCases {
		actual := injectServicesIntoTemplate(test.template, test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestInitiateAwsHandler(t *testing.T) {
	for _, test := range initiateAwsHandlerTestCases {
		actual := initiateAwsHandler(test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}

func TestInitiateServiceHandlers(t *testing.T) {
	for _, test := range initiateServiceHandlersTestCases {
		actual := initiateServiceHandlers(test.services)
		if test.expected != actual {
			t.Fatalf("\n>>> Expected:\n%s\n<<< but got:\n%s", test.expected, actual)
		}
	}
}
