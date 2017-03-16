package syncman

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

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

func TestHaltCommand(t *testing.T) {
	port := uint(6000)

	// given: start sync manager
	sm := newSyncManager(&settings{Port: port, ConfigPath: "config.xml"})
	go func() {
		sm.Run()
	}()

	// when: send HALT command
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Errorf("err is %v, want nil", err)
	}
	n, err := conn.Write([]byte("+HALT\r\n"))
	if n != len("+HALT\r\n") {
		t.Errorf("n is %d, want %d", n, len("+HALT\r\n"))
	}

	// then: -OK is resturend
	// then: sync manager is stopped
	_ = sm
}

func eqSettings(a *settings, b *settings) bool {
	return reflect.DeepEqual(a, b)
}
