package main

import "fmt"

func main() {
    bf := newBloomFilter(10, 2)
    bf.add("abc")
    bf.add("efg")
    bf.add("ijk")
    fmt.Println(bf.contains("abc"))
    fmt.Println(bf.contains("bc"))
    fmt.Println(bf.contains("c"))
    fmt.Println(bf.contains("ijk"))
}
