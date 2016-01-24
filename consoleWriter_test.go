// Copyright 2015
// Author: huangjunwei@youmi.net

package blog4go

import (
	"fmt"
	"testing"
)

func BenchmarkConsoleWriter(b *testing.B) {
	b.StopTimer()
	writer, err := NewConsoleWriter()
	defer writer.Close()
	if nil != err {
		fmt.Println(err.Error())
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		writer.Debugf("haha %s. en\\en, always %d and %f", "eddie", 18, 3.1415)
	}
}