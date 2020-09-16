// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/09/15 22:59:40

package caches

// Status is the status of cache.
type Status struct {

	// Count is how many entries stored in cache.
	Count int `json:"count"`

	// KeySize is the size of key.
	KeySize int `json:"keySize"`

	// ValueSize is the size of value.
	ValueSize int `json:"valueSize"`
}

// newStatus returns a new status holder.
func newStatus() *Status {
	return &Status{
		Count:     0,
		KeySize:   0,
		ValueSize: 0,
	}
}

// addEntry adds all information to status with key and value.
func (s *Status) addEntry(key string, value []byte) {
	s.Count++
	s.KeySize += len(key)
	s.ValueSize += len(value)
}

// subEntry subs all information to status with key and value
func (s *Status) subEntry(key string, value []byte) {
	s.Count--
	s.KeySize -= len(key)
	s.ValueSize -= len(value)
}
