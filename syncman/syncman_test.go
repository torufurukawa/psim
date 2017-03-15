package syncman

import "testing"
import "reflect"

func TestParseFlags(t *testing.T) {
	var fixtures = []struct {
		args     []string
		expected *settings
	}{
		{
			args:     []string{},
			expected: &settings{Port: 6000, ConfigPath: "config.xml"},
		},
		{
			args:     []string{"-port", "5000"},
			expected: &settings{Port: 5000, ConfigPath: "config.xml"},
		},
		{
			args:     []string{"-config", "foo.xml"},
			expected: &settings{Port: 6000, ConfigPath: "foo.xml"},
		},
		{
			args:     []string{"-port", "5000", "-config", "foo.xml"},
			expected: &settings{Port: 5000, ConfigPath: "foo.xml"},
		},
	}

	for _, f := range fixtures {
		result := parseFlags(f.args)
		if !eqSettings(result, f.expected) {
			t.Errorf("result for %v is \n%v, want \n%v", f.args, result, f.expected)
		}
	}
}

func eqSettings(a *settings, b *settings) bool {
	return reflect.DeepEqual(a, b)
}
