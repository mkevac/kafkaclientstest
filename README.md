# kafkaclientstest
Simple test to compare speed and amount of garbage for sarama and optiopay kafka clients

# results

Optiopay is faster (42s VS 1m), but creates a lot more garbage (73MiB VS 19MiB)

```
~> /usr/bin/time -v ./optiopaytest 
	Command being timed: "./optiopaytest"
	User time (seconds): 44.60
	System time (seconds): 3.88
	Percent of CPU this job got: 112%
	Elapsed (wall clock) time (h:mm:ss or m:ss): 0:42.96
	Average shared text size (kbytes): 0
	Average unshared data size (kbytes): 0
	Average stack size (kbytes): 0
	Average total size (kbytes): 0
	Maximum resident set size (kbytes): 73268
	Average resident set size (kbytes): 0
	Major (requiring I/O) page faults: 0
	Minor (reclaiming a frame) page faults: 17884
	Voluntary context switches: 468749
	Involuntary context switches: 2971
	Swaps: 0
	File system inputs: 0
	File system outputs: 8
	Socket messages sent: 0
	Socket messages received: 0
	Signals delivered: 0
	Page size (bytes): 4096
	Exit status: 0
```

```
~> /usr/bin/time -v ./saramatest 
May 18 10:11:59.560932 [INFO]  <45905> max offset is 25073416
May 18 10:13:01.499921 [INFO]  <45905> Consumed 10966892 messages
	Command being timed: "./saramatest"
	User time (seconds): 36.54
	System time (seconds): 14.85
	Percent of CPU this job got: 82%
	Elapsed (wall clock) time (h:mm:ss or m:ss): 1:01.95
	Average shared text size (kbytes): 0
	Average unshared data size (kbytes): 0
	Average stack size (kbytes): 0
	Average total size (kbytes): 0
	Maximum resident set size (kbytes): 19048
	Average resident set size (kbytes): 0
	Major (requiring I/O) page faults: 0
	Minor (reclaiming a frame) page faults: 3722
	Voluntary context switches: 2402727
	Involuntary context switches: 3599
	Swaps: 0
	File system inputs: 0
	File system outputs: 8
	Socket messages sent: 0
	Socket messages received: 0
	Signals delivered: 0
	Page size (bytes): 4096
	Exit status: 0
```
