# Scavenger
Scavenger is an application used for brute forcing directories very fast with big lists to use. It uses the threads in your computer to divy up your Word list then it brute forces.It is **Highly Recommended** that you use your own Word list since the included word list is an example word list used to test the program.

## Install
Unfortunately, it isn't all that easy right now to download and use since it is still under major development.

1. Download all the files
2. Unzip Scavenger-master.zip

#### Linux
1. Go to the file "Versions"
2. Under Linux is the Linux binary Version
3. Move that file to any path that best


## Run
#### Golang
To run with Go, do the following in your command line.
 ```sh
 cd $PATH/scavenger-master/
 go run scavenger.go -h
 ```
 #### Linux
 To run do:
  ```sh
  cd $PATH
  ./scavenger -h
  ```
 ## Flags

 There are a couple of flags you can use that may help you.

 #### Help(-h)
 Displays help menu
 ```sh
  go run scavenger.go -h
 ````
 #### Wordlist(-w)*
 This is used to tell where the wordlist are
 ```sh
go run scavenger.go -w all.txt -u example.com
 ```
 #### URL(-u)*
 Enter the base URL of the the site that will be brute forced
 ```sh
  go run scavenger.go -w all.txt -u example.com -v true
 ```
#### Output(-o)
Output flag sends all urls found to a file
```sh
go run scavenger.go -w all.txt -u example.com -o output.txt
```


*Required to run
