package syncman

import "testing"
import "reflect"

func TestParseFlags(t *testing.T) {
	var fixtures = []struct {
		args     []string
		expected *flags
	}{
		{
			args:     []string{"syncman", "-port", "6000", "-config", "config.xml"},
			expected: &flags{Port: 6000, ConfigPath: "config.xml"},
		},
	}
	_ = fixtures
	for _, f := range fixtures {
		result := parseFlags(f.args)
		if !eqFlags(result, f.expected) {
			t.Errorf("result for %v is \n%v, want \n%v", f.args, result, f.expected)
		}
	}
}

func eqFlags(a *flags, b *flags) bool {
	return reflect.DeepEqual(a, b)
}
