package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
)

//global varibels
var Output bool = false
var Ht string = "http://"
var Verbose bool = false
var Found []string

func main() {
	//Start wait WaitGroup
	var wg sync.WaitGroup

	//load all flags
	list := flag.String("w", "", "Path to the word list")
	url := flag.String("u", "", "Base of the url to be bruted(EX: google.com)")
	ofile := flag.String("o", "", "Path to output file")
	ver := flag.String("v", "", "Verbose activity")

	//Parse all of the flags
	flag.Parse()

	//Test all of the flags
	if *list == "" {
		if *url == "" {
			panic("No list(-w) or url(-u) given. Use -h for help.")
		} else {
			panic("No list(-w) for brute force given. Use -h for help.")
		}
	}
	if *url == "" {
		panic("No url(-u) given. Use -h for help.")
	}
	if *ofile != "" {
		Output = true
	}
	if *ver != "" {
		Verbose = true
	}
	//See how many threads a computer has and add them to wait WaitGroup
	wg.Add(runtime.NumCPU())

	//Test for valid url
	_, err := http.Get("http://" + *url)
	if err != nil {
		panic(`
      Not a valid url
          [Hint] Make sure you DON'T have "http://www." in your url
          [Hint] If you don't have internet, this also can be triggered`)
	}

	//Test if https works
	_, err = http.Get("https://" + *url)
	if err == nil {
		Ht = "https://"
	}

	//Test for valid wordlist
	test, err := os.Open(*list)
	if err != nil {
		panic(`Not a valid wordlist
              [Hint] Make sure your give the WHOLE file name end extension
              [Hint] Make sure that that is a direct path from the file your in`)
	}
	test.Close()
	//Test or make Outputfile
	if Output == true {
		test, err = os.Open(*ofile)
		if err != nil {
			test, _ := os.Create(*ofile)
			test.Close()
		}
		test.Close()
	}
	r, _ := os.Open(*list)

	defer r.Close()

	//Prep Scanner
	s := bufio.NewScanner(r)

	//Prep Bruter
	var bru []string
	var n int = 0
	ur := (Ht + *url + "/")
	for s.Scan() {
		bru = append(bru, s.Text())
		n++
	}
	//Divid up the wordlist to each thread
	spacer := n / runtime.NumCPU()

	// Get invalid method
	rt, _ := http.Get(ur + "/ubgkvjkgfvjhvjvhjchkcjhbkjgvkhio7yi86tr65dcj")
	if rt.StatusCode == 200 {
		fmt.Println("Brute force maybe ineffective against target.")
		fmt.Println(ur + "/ubgkvjkgfvjhvjvhjchkcjhbkjgvkhio7yi86tr65dcj returned status code: " + rt.Status)
	}
	invalid := rt.StatusCode
	//Brute force url and send them to Output (if selected)
	n = runtime.NumCPU()
	start := 0
	end := spacer
	fmt.Println("Starting bruteforce...")
	if Verbose == true {
		for n != 0 {
			go vBrute(bru, ur, invalid, start, end, &wg)
			n--
			start = start + spacer
			end = end + spacer
		}
	} else {
		for n != 0 {
			go brute(bru, ur, invalid, start, end, &wg)
			n--
			start = start + spacer
			end = end + spacer
		}
	}
	wg.Wait()
	//Print Findings
	fmt.Println("These URLS were found: ")
	for _, text := range Found {
		fmt.Println(text)

		//Send Output to a file(if selected)
		if Output == true {
			f, _ := os.OpenFile(*ofile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer f.Close()
			f.WriteString(text + "\n")
		}
	}
	if Output == true{
		fmt.Println("All URLS found were saved to " + *ofile)
	}
}

//Brute with Verbose
func vBrute(bru []string, ur string, invalid int, start int, stop int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n, text := range bru {
		if n >= start {
			if n <= stop {
				re, _ := http.Get(ur + text)
				fmt.Println(ur + text + " => " + re.Status)
				if re.StatusCode != invalid {
					Found = append(Found, ur+text)
				}
			}
		}
	}
}

//Normal Brute
func brute(bru []string, ur string, invalid int, start int, stop int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n, text := range bru {
		if n >= start {
			if n <= stop {
				re, _ := http.Get(ur + text)
				if re.StatusCode != invalid {
					Found = append(Found, ur+text)
				}
			}
		}
	}
}
