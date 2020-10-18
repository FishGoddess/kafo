// Copyright 2020 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/09/05 22:06:34

package main

import (
	"flag"
	"log"

	"github.com/FishGoddess/kafo/caches"
	"github.com/FishGoddess/kafo/servers"
)

func main() {

	// Parse all flags
	address := flag.String("address", ":5837", "The address used to listen, such as 127.0.0.1:5837.")

	options := caches.DefaultOptions()
	flag.IntVar(&options.MaxEntrySize, "maxEntrySize", options.MaxEntrySize, "The max memory size that entries can use. The unit is GB.")
	flag.IntVar(&options.MaxGcCount, "maxGcCount", options.MaxGcCount, "The max count of entries that gc will clean.")
	flag.IntVar(&options.GcDuration, "gcDuration", options.GcDuration, "The duration between two gc tasks. The unit is Minute.")
	flag.StringVar(&options.DumpFile, "dumpFile", options.DumpFile, "The file used to dump the cache.")
	flag.IntVar(&options.DumpDuration, "dumpDuration", options.DumpDuration, "The duration between two dump tasks. The unit is Minute.")
	flag.IntVar(&options.MapSizeOfSegment, "mapSizeOfSegment", options.MapSizeOfSegment, "The map size of segment.")
	flag.IntVar(&options.SegmentSize, "segmentSize", options.SegmentSize, "The number of segment in a cache. This value should be the pow of 2 for precision.")
	flag.IntVar(&options.CasSleepTime, "casSleepTime", options.CasSleepTime, "The time of sleep in one cas step. The unit is Microsecond.")

	serverType := flag.String("serverType", "tcp", "The type of server (http, tcp).")
	flag.Parse()

	// Initialize
	cache := caches.NewCacheWith(options)
	cache.AutoGc()
	cache.AutoDump()

	log.Printf("Kafo is running on %s at %s.", *serverType, *address)
	err := servers.NewServer(*serverType, cache).Run(*address)
	if err != nil {
		panic(err)
	}
}
