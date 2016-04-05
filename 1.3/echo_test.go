package echo_test

import(
  "testing"
  "strings"
  "os"
)

func BenchmarkIterative(b * testing.B) {
  for j := 0; j < b.N; j++ {
    var s, sep string
    for i := 0; i< len(os.Args); i++ {
      s += sep + os.Args[i]
      sep = " "
    }
  }
}

func BenchmarkJoin(b * testing.B) {
  for j := 0; j < b.N; j++ {
    strings.Join(os.Args, " ")
  }
}
