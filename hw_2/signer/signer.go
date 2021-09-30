package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func startWorker(jobFunc job, in, out chan interface{}, waiter *sync.WaitGroup) {
	defer waiter.Done()
	defer close(out)

	jobFunc(in, out)
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, j := range jobs {
		wg.Add(1)
		out := make(chan interface{})
		go startWorker(j, in, out, wg)
		in = out
	}

	wg.Wait()
}

type DataSigner func(string2 string) string

func DataSignerWorker(dataSigner DataSigner, data string, out chan string) {
	out <- dataSigner(data)
}

const kMultiHash int = 6

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for ch := range in {
		data := strconv.Itoa(ch.(int))

		md5 := DataSignerMd5(data)

		firstTermCh := make(chan string)
		go DataSignerWorker(DataSignerCrc32, data, firstTermCh)

		secondTermCh := make(chan string)
		go DataSignerWorker(DataSignerCrc32, md5, secondTermCh)

		wg.Add(1)
		go func() {
			defer wg.Done()

			out <- (<-firstTermCh) + "~" + (<-secondTermCh)
		}()
	}

	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for ch := range in {
		data := ch.(string)

		wg.Add(1)

		go func() {
			defer wg.Done()

			crc32Arr := make([]string, kMultiHash)

			wgRes := &sync.WaitGroup{}
			for th := 0; th < kMultiHash; th++ {
				wgRes.Add(1)
				go func(th int, data string, crc32 *string, wg *sync.WaitGroup) {
					defer wg.Done()
					*crc32 = DataSignerCrc32(strconv.Itoa(th)+data)
				}(th, data, &crc32Arr[th], wgRes)
			}
			wgRes.Wait()

			out <- strings.Join(crc32Arr, "")
		}()
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var results []string
	for ch := range in {
		results = append(results, ch.(string))
	}

	sort.Strings(results)

	out <- strings.Join(results, "_")
}
