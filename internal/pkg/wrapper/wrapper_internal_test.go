package wrapper

import "testing"

func TestInjectLibraryIntoTemplate(t *testing.T) {
	for _, test := range injectLibraryIntoTemplateTestCases {
		actual := injectLibraryIntoTemplate(test.template, test.libraryName)
		if test.expected != actual {
			t.Fatalf("Expected %s but got %s", test.expected, actual)
		}
	}
}

func TestInjectServicesIntoTemplate(t *testing.T) {
	for _, test := range injectServicesIntoTemplateTestCases {
		actual := injectServicesIntoTemplate(test.template, test.services)
		if test.expected != actual {
			t.Fatalf("Expected %s but got %s", test.expected, actual)
		}
	}
}

func TestInitiateAwsHandler(t *testing.T) {
	for _, test := range initiateAwsHandlerTestCases {
		actual := initiateAwsHandler(test.services)
		if test.expected != actual {
			t.Fatalf("Expected %s but got %s", test.expected, actual)
		}
	}
}

func TestInitialeServiceHandlers(t *testing.T) {
	for _, test := range initialeServiceHandlersTestCases {
		actual := initialeServiceHandlers(test.services)
		if test.expected != actual {
			t.Fatalf("Expected %s but got %s", test.expected, actual)
		}
	}
}