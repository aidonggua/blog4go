// Copyright (c) 2015, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package blog4go

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

type MyHook struct {
	cnt     int
	level   Level
	message string
}

func (hook *MyHook) add() {
	hook.cnt++
}

func (hook *MyHook) Fire(level Level, args ...interface{}) {
	hook.add()
	hook.level = level
	hook.message = fmt.Sprint(args...)
}

func TestHook(t *testing.T) {
	hook := new(MyHook)
	hook.cnt = 0

	err := NewFileWriter("/tmp", false)
	defer Close()
	if nil != err {
		t.Errorf("initialize file writer faied. err: %s", err.Error())
	}

	blog.SetHook(hook)
	blog.SetHookLevel(INFO)

	blog.Debug("something")
	// wait for hook called
	time.Sleep(1 * time.Millisecond)
	if 0 != hook.cnt {
		t.Error("hook called not valid")
	}

	if DEBUG == hook.level || "something" == hook.message {
		t.Errorf("hook parameters wrong. level: %s, message: %s", hook.level.String(), hook.message)
	}

	blog.Info("yes")
	// wait for hook called
	time.Sleep(1 * time.Millisecond)
	if 1 != hook.cnt {
		t.Error("hook not called")
	}

	if INFO != hook.level || "yes" != hook.message {
		t.Errorf("hook parameters wrong. level: %d, message: %s", hook.level, hook.message)
	}

	// clean logs
	_, err = exec.Command("/bin/sh", "-c", "/bin/rm /tmp/*.log*").Output()
	if nil != err {
		t.Errorf("clean files failed. err: %s", err.Error())
	}
}
