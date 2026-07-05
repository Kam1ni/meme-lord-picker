package fetcher

import (
	"fmt"
	"meme-lord-picker/memelord"
	"sync"
	"time"
)

type OnDataFunc func(memelord.MemesResponse)

type Fetcher struct {
	isQueued         bool
	mutex            sync.Mutex
	queuedQuery      string
	lastFetchedQuery string
	client           memelord.Client
	onData           OnDataFunc
	didFirstFetch    bool
}

func CreateFetcher(client memelord.Client, onData OnDataFunc) Fetcher {
	return Fetcher{
		client: client,
		onData: onData,
	}
}

func (f *Fetcher) fetch() {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.isQueued = false
	query := f.queuedQuery
	if f.lastFetchedQuery == query && f.didFirstFetch {
		return
	}
	result, err := f.client.FetchMemes(stringQueryToMemelordQuery(query))
	if err != nil {
		fmt.Println("Failed to fetch memes\n", err.Error())
		return
	}
	f.didFirstFetch = true
	f.lastFetchedQuery = f.queuedQuery
	f.onData(result)
}

func (f *Fetcher) QueueFetch(query string) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.queuedQuery = query
	if f.isQueued {
		return
	}
	f.isQueued = true
	go func() {
		time.Sleep(500 * time.Millisecond)
		f.fetch()
	}()
}
