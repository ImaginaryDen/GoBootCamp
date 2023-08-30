package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	mean   = 1
	median = 2
	mode   = 4
	sd     = 8
	sorted = 16
)

type Day0Struct struct {
	mean   float64
	median float64
	mode   int
	sd     float64

	nums []int
	key  int
}

func (res *Day0Struct) GetMean() float64 {
	if mean&res.key == 0 {
		sum := 0
		for _, d := range res.nums {
			sum += d
		}
		res.mean = float64(sum) / float64(len(res.nums))
	}
	res.key |= mean
	return res.mean
}

func (res *Day0Struct) GetMedian() float64 {
	if median&res.key == 0 {

		if sorted&res.key == 0 {
			sort.Ints(res.nums)
		}
		res.key |= sorted
		l := len(res.nums)
		if l%2 == 0 {
			res.median = float64((res.nums[l/2-1] + res.nums[l/2]) / 2)
		} else {
			res.median = float64(res.nums[l/2])
		}
	}
	res.key |= median
	return res.median
}

func (res *Day0Struct) GetMode() int {
	countMap := make(map[int]int)
	if mode&res.key == 0 {
		for _, d := range res.nums {
			countMap[d]++
		}

		maxD := 0
		for _, key := range res.nums {
			freq := countMap[key]
			if freq > maxD {
				res.mode = key
				maxD = freq
			}
		}
	}
	res.key |= mode
	return res.mode
}

func (res *Day0Struct) GetSd() float64 {
	if sd&res.key == 0 {
		for _, key := range res.nums {
			res.sd += math.Pow(float64(key)-res.GetMean(), 2)
		}
		res.sd = math.Sqrt(res.sd / 10)
	}
	res.key |= sd
	return res.sd
}

func GetArrayStdin() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, errParse := strconv.Atoi(scanner.Text())
		isRightBounds := num >= -100_000 && num <= 100_000
		if errParse != nil || !isRightBounds {
			fmt.Println("Error: input must be an integer number")
			continue
		}
		if errParse != io.EOF {
			nums = append(nums, num)
		}
	}
	return nums
}

func initArgs() [4]*bool {
	return [4]*bool{
		flag.Bool("mean", false, "Print mean value"),
		flag.Bool("median", false, "Print median value"),
		flag.Bool("mode", false, "Print mode value"),
		flag.Bool("sd", false, "Print SD value"),
	}
}

func main() {
	var metrics = initArgs()
	flag.Parse()
	var res Day0Struct = Day0Struct{nums: GetArrayStdin()}
	if len(res.nums) == 0 {
		fmt.Println("Error: Empty data entry")
		return
	}
	// fmt.Println(res.nums)
	var defaultArg = true
	for i, value := range metrics {
		if *value == false {
			continue
		}
		defaultArg = false
		switch i {
		case 0:
			fmt.Println("mean", res.GetMean())
		case 1:
			fmt.Println("median", res.GetMedian())
		case 2:
			fmt.Println("mode", res.GetMode())
		case 3:
			fmt.Println("sd", res.GetSd())
		}
	}
	if defaultArg {
		fmt.Println("mean", res.GetMean())
		fmt.Println("median", res.GetMedian())
		fmt.Println("mode", res.GetMode())
		fmt.Println("sd", res.GetSd())
	}
}
