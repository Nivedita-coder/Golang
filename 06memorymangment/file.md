memory managment is pretty simple in golang as compare to other languages. Memory allocation and deallocation happens automatically in golang.

In memory managment, you will usually gpoing to use these two methods.
1. new()
(i) allocate memory but no INIT
(ii) you will get a memeory address
(iii) zeroed storage - in which you cannot put data initially.

2. make()
(i) allocate memeory and INIT
(ii) you will get a memory address
(iii) non-zeroed storage - you can put any of the data.

GC happens automatically
