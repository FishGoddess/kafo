// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/10/01 16:31:41

package _examples

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	// concurrency is the concurrency of test.
	concurrency = 1000
)

// testTask is a wrapper wraps task to testTask.
func testTask(task func(no int)) string {

	beginTime := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			task(no)
		}(i)
	}
	wg.Wait()
	return time.Now().Sub(beginTime).String()
}

// go test -v performance_test.go -run=^TestHttpServer$
func TestHttpServer(t *testing.T) {

	writeTime := testTask(func(no int) {
		data := strconv.Itoa(no)
		request, err := http.NewRequest("PUT", "http://localhost:5837/v1/cache/"+data, strings.NewReader(data))
		if err != nil {
			t.Fatal(err)
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			t.Fatal(err)
		}
		response.Body.Close()
	})

	t.Logf("写入消耗时间为 %s！", writeTime)

	time.Sleep(3 * time.Second)

	readTime := testTask(func(no int) {
		data := strconv.Itoa(no)
		request, err := http.NewRequest("GET", "http://localhost:5837/v1/cache/"+data, nil)
		if err != nil {
			t.Fatal(err)
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			t.Fatal(err)
		}
		response.Body.Close()
	})

	t.Logf("读取消耗时间为 %s！", readTime)
}
