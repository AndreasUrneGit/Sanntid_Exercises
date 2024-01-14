What happens in the double threaded running of main, is that the two threads acts on the same variable and does calculations on it as soon as they possibly can.
The problem here is that the other thread might write or read from the variable in the meantime that the thread is doing the calculations. 
Paradoxically we therefore do not know what will happen to the end value even though we do run increment and decrement an equal number of million times.
This can be illustrated by the following example
i = 1
thread 1 read i = 1 and starts the calculation to increment
Meenwile thread 2 also reads i = 1 and starts the decrement
thread 1 writes back 2 to i
thread 2 then also writes back 0 to i and i = 0 instead of 1.

GOMAXPROCS seems to alter how many threads are allowed to be used simultaniously.
When altered to 1 the ressult is o as expected since it will first run the first thread, then the seccond.