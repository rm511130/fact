package main
import ( "fmt"
         "github.com/go-martini/martini"
         "strconv" )

func fact(n uint64) uint64 { if n == 0 { return 1 }; return n * fact(n-1) }

func main() { m := martini.Classic() 
              m.Get("/**", func(params martini.Params) string { s, err := strconv.Atoi(params["_1"])
                                                                fmt.Sprintf("%s %d\n",err,s);
                                                                return "Calculating Factorials: " + params["_1"] + "! = " + fmt.Sprintf("%d\n", fact(uint64(s)))
                                                              })
              m.Get("/", func() string { return "Try passing an integer! Try localhost:3000/5" })
              m.Run()
            }
