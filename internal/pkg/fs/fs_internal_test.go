package fs

import "testing"
import s "strings"

func TestFilepathWalk(t *testing.T) {
	for _, test := range filepathWalkTestCases {
		actual := filepathWalk(test.basedir, test.source, test.archive)(test.path, test.info, test.err)

		if s.Compare(test.expected.Error(), actual.Error()) != 0 {
			t.Fatalf("\n>>> Expected:\n%v\n<<< but got:\n%v", test.expected, actual)
		}
	}
}
