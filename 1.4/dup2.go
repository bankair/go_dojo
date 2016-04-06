package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// Duplicated comment
// Duplicated comment

func main() {
  counts := make(map[string]map[string]int)
  files := os.Args[1:]
  for _, arg := range files {
    f, err := os.Open(arg)
    if err != nil {
      fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
      continue
    }
    countLines(arg, f, counts)
    f.Close()
  }
  for line, file_map := range counts {
    target_files := make([]string, 0)
    total := 0
    for path, count := range file_map {
      total += count
      target_files = append(target_files, path)
    }
    if total > 1 {
      fmt.Printf("%d\t%v\t(%v)\n", total, line, strings.Join(target_files, ";"))
    }
  }
}

func countLines(path string, f *os.File, counts map[string]map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    line := input.Text()
    if len(counts[line]) == 0 {
      counts[line] = make(map[string]int)
    }
    counts[line][path]++
  }
}
