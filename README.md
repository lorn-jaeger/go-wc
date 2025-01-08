### go-wc
A simple wc implementation in Go. 


### usage

```
  -c 
        Count the bytes in a file
  -l 
        Count the number of lines in a file
  -w 
        Count the number of words in a file
  -m 
        Count the number or characters in a file

  call the binary with no flags to get the lines, words and bytes
```

### speed
Tested by running the command on *The Art of War* a thousand times.
#### go-wc
``` shell
$ time for i in {1..1000}; do ./go-wc test.txt > /dev/null; done
```
```
  real	0m5.469s
  user	0m4.008s
  sys	0m1.651s
```
#### wc
``` shell
$ time for i in {1..1000}; do wc test.txt > /dev/null; done
```
```
  real	0m3.430s
  user	0m2.801s
  sys	0m0.645s
```
