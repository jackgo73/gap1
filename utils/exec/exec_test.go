package exec

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os/exec"
	"testing"
	"time"
)

var (
	sleepbin, _ = exec.LookPath("sleep")
	echobin, _  = exec.LookPath("echo")
	shell, _    = exec.LookPath("sh")
)

func TestRunTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test due to random failures.")
	}
	if sleepbin == "" {
		t.Skip("'sleep' binary not available on OS, skipping.")
	}
	cmd := exec.Command(sleepbin, "10")
	start := time.Now()
	err := RunTimeout(cmd, time.Millisecond*20)
	elapsed := time.Since(start)

	assert.Equal(t, TimeoutErr, err)
	// Verify that command gets killed in 20ms, with some breathing room
	assert.True(t, elapsed < time.Millisecond*75)
}

func TestRunTimeoutFastExit(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test due to random failures.")
	}
	if echobin == "" {
		t.Skip("'echo' binary not available on OS, skipping.")
	}
	cmd := exec.Command(echobin)
	start := time.Now()
	err := RunTimeout(cmd, time.Millisecond*20)
	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	elapsed := time.Since(start)

	require.NoError(t, err)
	// Verify that command gets killed in 20ms, with some breathing room
	assert.True(t, elapsed < time.Millisecond*75)

	// Verify "process already finished" log doesn't occur.
	time.Sleep(time.Millisecond * 75)
	require.Equal(t, "", buf.String())
}


func TestCombinedOutputTimeout(t *testing.T) {
	// TODO: Fix this test
	t.Skip("Test failing too often, skip for now and revisit later.")
	if sleepbin == "" {
		t.Skip("'sleep' binary not available on OS, skipping.")
	}
	cmd := exec.Command(sleepbin, "10")
	start := time.Now()
	_, err := CombinedOutputTimeout(cmd, time.Millisecond*20)
	elapsed := time.Since(start)

	assert.Equal(t, TimeoutErr, err)
	// Verify that command gets killed in 20ms, with some breathing room
	assert.True(t, elapsed < time.Millisecond*75)
}


func TestCombinedOutput(t *testing.T) {
	if echobin == "" {
		t.Skip("'echo' binary not available on OS, skipping.")
	}
	cmd := exec.Command(echobin, "foo")
	out, err := CombinedOutputTimeout(cmd, time.Second)

	assert.NoError(t, err)
	assert.Equal(t, "foo\n", string(out))
}


// test that CombinedOutputTimeout and exec.Cmd.CombinedOutput return
// the same output from a failed command.
func TestCombinedOutputError(t *testing.T) {
	if shell == "" {
		t.Skip("'sh' binary not available on OS, skipping.")
	}
	cmd := exec.Command(shell, "-c", "false")
	expected, err := cmd.CombinedOutput()

	cmd2 := exec.Command(shell, "-c", "false")
	actual, err := CombinedOutputTimeout(cmd2, time.Second)

	assert.Error(t, err)
	assert.Equal(t, expected, actual)
}


func TestRunError(t *testing.T) {
	if shell == "" {
		t.Skip("'sh' binary not available on OS, skipping.")
	}
	cmd := exec.Command(shell, "-c", "false")
	err := RunTimeout(cmd, time.Second)

	assert.Error(t, err)
}

