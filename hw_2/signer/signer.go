package main

import (
	"fmt"
	"sort"
	"strconv"
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

	for ch := range in {
		fmt.Println(ch)
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

		DataSignerOut1 := make(chan string)
		go DataSignerWorker(DataSignerCrc32, data, DataSignerOut1)

		DataSignerOut3 := make(chan string)
		go DataSignerWorker(DataSignerCrc32, md5, DataSignerOut3)

		wg.Add(1)
		go func() {
			defer wg.Done()

			s1 := <-DataSignerOut1
			s3 := <-DataSignerOut3

			out <- s1 + "~" + s3
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

			var chs [kMultiHash]chan string
			for i := range chs {
				chs[i] = make(chan string)
			}

			for th := 0; th < kMultiHash; th++ {
				go DataSignerWorker(DataSignerCrc32, strconv.Itoa(th)+data, chs[th])
			}

			var result string

			for i := range chs {
				result += <-chs[i]
			}

			out <- result
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

	var resultStr string

	for _, str := range results {
		resultStr += str + "_"
	}

	out <- resultStr[:len(resultStr)-1]
}
