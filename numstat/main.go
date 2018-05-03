package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stat struct {
	min, max, count int64
	sum             int64
	hasData         bool
}

func NewStat() *Stat {
	return &Stat{}
}

func (s *Stat) Update(v int64) {
	// s.sum.Add(s.sum, big.NewInt(v))
	s.sum += v
	s.count++
	if s.min > v || !s.hasData {
		s.min = v
	}
	if s.max < v || !s.hasData {
		s.max = v
	}
	s.hasData = true
}

func (s *Stat) String() string {
	if s.hasData {
		return fmt.Sprintf("min: %v max: %v mean: %v",
			s.min,
			s.max,
			float64(s.sum)/float64(s.count))
	}
	return "No stats yet"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s filePath\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	stat := NewStat()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "-" {
			fmt.Printf("%s\n", stat)
			continue
		}

		if newVal, err := strconv.ParseInt(str, 10, 64); err == nil {
			stat.Update(newVal)
		} else {
			fmt.Println(err)
		}
	}
}
