package immortal

import (
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestWatchPidGetpid(t *testing.T) {
	sup := new(Sup)
	ch := make(chan error)

	cmd := exec.Command("go", "version")
	cmd.Start()
	pid := cmd.Process.Pid
	go func() {
		sup.WatchPid(pid, ch)
		ch <- cmd.Wait()
	}()
	select {
	case <-time.After(time.Millisecond):
		syscall.Kill(pid, syscall.SIGTERM)
	case err := <-ch:
		if err != nil {
			if err.Error() != "EXIT" {
				t.Error(err)
			}
		}
	}
}

func TestWatchPidGetpidKill(t *testing.T) {
	sup := new(Sup)
	ch := make(chan error)

	cmd := exec.Command("sleep", "100")
	cmd.Start()
	pid := cmd.Process.Pid
	go func() {
		sup.WatchPid(pid, ch)
		ch <- cmd.Wait()
	}()

	select {
	case err := <-ch:
		if err != nil {
			if err.Error() != "EXIT" {
				t.Error(err)
			}
		}
	case <-time.After(1 * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			t.Errorf("failed to kill: %s", err)
		}
	}
}
