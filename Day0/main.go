package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"sort"
)

const (
	mean   = 1
	median = 2
	mode   = 4
	sd     = 8
	sorted = 16
)

type result struct {
	mean   float64
	median float64
	mode   int
	sd     float64

	nums []int
	key  int
}

func (res *result) GetMean() float64 {
	if mean&(*res).key == 0 {
		sum := 0
		for _, d := range (*res).nums {
			sum += d
		}
		(*res).mean = float64(sum) / float64(len((*res).nums))
	}
	(*res).key |= mean
	return (*res).mean
}

func (res *result) GetMedian() float64 {
	if median&(*res).key == 0 {

		if sorted&(*res).key == 0 {
			sort.Ints((*res).nums)
		}
		(*res).key |= sorted
		l := len((*res).nums)
		if l%2 == 0 {
			res.median = float64(((*res).nums[l/2-1] + (*res).nums[l/2]) / 2)
		} else {
			res.median = float64((*res).nums[l/2])
		}
	}
	(*res).key |= median
	return (*res).median
}

func (res *result) GetMode() float64 {
	countMap := make(map[int]int)
	if mode&(*res).key == 0 {
		for _, d := range (*res).nums {
			countMap[d]++
		}

		maxD := 0
		for _, key := range (*res).nums {
			freq := countMap[key]
			if freq > maxD {
				res.mode = key
				maxD = freq
			}
		}
	}
	(*res).key |= mode
	return (*res).mean
}

func (res *result) GetSd() float64 {
	if sd&(*res).key == 0 {
		for _, key := range (*res).nums {
			res.sd += math.Pow(float64(key)-res.GetMean(), 2)
		}
		res.sd = math.Sqrt(res.sd / 10)
	}
	(*res).key |= sd
	return (*res).sd
}

func GetArrayStdin() []int {
	var nums []int
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	fmt.Println(nums)
	return nums
}

func main() {
	var res result = result{nums: GetArrayStdin()}

	var metrics = [4]*bool{
		flag.Bool("mean", false, "Print mean value"),
		flag.Bool("median", false, "Print median value"),
		flag.Bool("mode", false, "Print mode value"),
		flag.Bool("sd", false, "Print SD value"),
	}
	flag.Parse()

	for i, value := range metrics {
		if *value == false {
			continue
		}
		switch i {
		case 0:
			fmt.Println(res.GetMean())
		case 1:
			fmt.Println(res.GetMedian())
		case 2:
			fmt.Println(res.GetMode())
		case 3:
			fmt.Println(res.GetSd())
		}
	}
}
