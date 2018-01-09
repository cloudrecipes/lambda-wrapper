package fs

import "testing"

func TestFilepathWalk(t *testing.T) {
	for _, test := range filepathWalkTestCases {
		actual := filepathWalk(test.basedir, test.source, test.archive)(test.path, test.info, test.err)

		if test.expected.Error() != actual.Error() {
			t.Fatalf("\n>>> Expected:\n%v\n<<< but got:\n%v", test.expected, actual)
		}
	}
}
