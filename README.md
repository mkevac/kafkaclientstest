# kafkaclientstest
Simple test to compare speed and amount of garbage for sarama, optiopay and confluent kafka clients

# results

Optiopay is fastest but creates a lot more garbage.

```
	Command being timed: "./saramatest"
	User time (seconds): 17.32
	System time (seconds): 8.60
	Percent of CPU this job got: 73%
	Elapsed (wall clock) time (h:mm:ss or m:ss): 0:35.24
	Average shared text size (kbytes): 0
	Average unshared data size (kbytes): 0
	Average stack size (kbytes): 0
	Average total size (kbytes): 0
	Maximum resident set size (kbytes): 14088
	Average resident set size (kbytes): 0
	Major (requiring I/O) page faults: 0
	Minor (reclaiming a frame) page faults: 2447
	Voluntary context switches: 1062218
	Involuntary context switches: 1880
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
	Command being timed: "./optiopaytest"
	User time (seconds): 14.74
	System time (seconds): 1.53
	Percent of CPU this job got: 56%
	Elapsed (wall clock) time (h:mm:ss or m:ss): 0:29.02
	Average shared text size (kbytes): 0
	Average unshared data size (kbytes): 0
	Average stack size (kbytes): 0
	Average total size (kbytes): 0
	Maximum resident set size (kbytes): 63812
	Average resident set size (kbytes): 0
	Major (requiring I/O) page faults: 0
	Minor (reclaiming a frame) page faults: 15501
	Voluntary context switches: 308722
	Involuntary context switches: 1666
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
	Command being timed: "./confluenttest"
	User time (seconds): 13.22
	System time (seconds): 2.99
	Percent of CPU this job got: 48%
	Elapsed (wall clock) time (h:mm:ss or m:ss): 0:33.34
	Average shared text size (kbytes): 0
	Average unshared data size (kbytes): 0
	Average stack size (kbytes): 0
	Average total size (kbytes): 0
	Maximum resident set size (kbytes): 19820
	Average resident set size (kbytes): 0
	Major (requiring I/O) page faults: 0
	Minor (reclaiming a frame) page faults: 725979
	Voluntary context switches: 506082
	Involuntary context switches: 204
	Swaps: 0
	File system inputs: 0
	File system outputs: 0
	Socket messages sent: 0
	Socket messages received: 0
	Signals delivered: 0
	Page size (bytes): 4096
	Exit status: 0
```
