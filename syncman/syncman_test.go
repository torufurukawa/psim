package syncman

import "testing"
import "reflect"

func TestParseFlags(t *testing.T) {
	result, err := parseFlags([]string{"syncman", "-port", "6000", "-config", "config.xml"})
	if err != nil {
		t.Errorf("err is %v, want nil", err)
	}
	expected := &flags{Port: 6000, ConfigPath: "config.xml"}
	if !eqFlags(result, expected) {
		t.Errorf("result is \n%v, want \n%v", result, expected)
	}
}

func eqFlags(a *flags, b *flags) bool {
	return reflect.DeepEqual(a, b)
}
