// @Author: Ciusyan 2023/7/29

package other

import (
	"net/http"
	"sync"
	"time"
)

type HttpClientPool struct {
	pool sync.Pool
}

func NewHttpClientPool() *HttpClientPool {
	return &HttpClientPool{

		//
		pool: sync.Pool{
			New: func() interface{} {
				client := http.Client{
					Timeout: 10 * time.Second,
				}

				return client
			},
		},
	}
}

func (h *HttpClientPool) Get() *http.Client {
	return h.pool.Get().(*http.Client)
}

func (h *HttpClientPool) Put(client *http.Client) {
	h.pool.Put(client)
}
