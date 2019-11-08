# Process for the sieving method
## Based on PrimeRT-Lz
https://github.com/PandaLz/PrimeRT-Lz

### Method 1
#### Find sieve of x * 2 squared

#### Example value x = 5

Start by finding the width of the values to test which will be same as x, so, 5 columns starting from 3 in steps of 2.

H | 3 | 5 | 7 | 9 | 11
--|---|---|---|---|---

Step two is find starting points for each column, starting from the column from the right which is 3, this value is obtained x - 2, then getting the triangular of it, in this case triangular of 3 the last step is multiply by 4 so get 24, the initial position of 3 column start with 0

Next column is triangular * 4 of current triangular - 1, and so on until triangular is 1, in this case column 5 would be triangular of 2 * 4 and 7 triangular of 1 * 4, both cases start with 0, must remain 2 columns unfilled, both start with 0 in this case 9 and 11 columns.

H | 3 | 5 | 7 | 9 | 11
--|---|---|---|---|---
0 | 0 |   |   |   |   
1 |   |   |   |   |   
2 |   |   |   |   |   
3 |   |   |   |   |   
4 |   |   |   |   |   
5 |   |   |   |   |   
6 |   |   |   |   |   
7 |   |   |   |   |   
8 |   |   |   |   |   
9 |   |   |   |   |   
10|   |   |   |   |   
11|   |   |   |   |   
12|   | 0 |   |   |   
13|   |   |   |   |   
14|   |   |   |   |   
15|   |   |   |   |   
16|   |   |   |   |   
17|   |   |   |   |   
18|   |   |   |   |   
19|   |   |   |   |   
20|   |   | 0 |   |   
21|   |   |   |   |   
22|   |   |   |   |   
23|   |   |   |   |   
24|   |   |   | 0 | 0 

Step 3 is add the height of range that will be used to get primes if any, the height is determined by ((x * 2) + 1) * 2 so add 22 to the current height, next, can fill up values starting from their starting zero position to it's corresponding nth which will be 0 and so on.

H |   3   |   5   |   7   |   9   |  11
--|-------|-------|-------|-------|-------
0 | **0** |       |       |       |   
1 |   1   |       |       |       |   
2 |   2   |       |       |       |   
3 | **0** |       |       |       |   
4 |   1   |       |       |       |   
5 |   2   |       |       |       |   
6 | **0** |       |       |       |   
7 |   1   |       |       |       |   
8 |   2   |       |       |       |   
9 | **0** |       |       |       |   
10|   1   |       |       |       |   
11|   2   |       |       |       |   
12| **0** | **0** |       |       |   
13|   1   |   1   |       |       |   
14|   2   |   2   |       |       |   
15| **0** |   3   |       |       |   
16|   1   |   4   |       |       |   
17|   2   | **0** |       |       |   
18| **0** |   1   |       |       |   
19|   1   |   2   |       |       |   
20|   2   |   3   | **0** |       |   
21| **0** |   4   |   1   |       |   
22|   1   | **0** |   2   |       |   
23|   2   |   1   |   3   |       |   
24| **0** |   2   |   4   | **0** | **0** 
25|   1   |   3   |   5   |   1   |   1 
26|   2   |   4   |   6   |   2   |   2 
27| **0** | **0** | **0** |   3   |   3 
28|   1   |   1   |   1   |   4   |   4 
29|   2   |   2   |   2   |   5   |   5 
30| **0** |   3   |   3   |   6   |   6 
31|   1   |   4   |   4   |   7   |   7 
32|   2   | **0** |   5   |   8   |   8 
33| **0** |   1   |   6   | **0** |   9 
34|   1   |   2   | **0** |   1   |   10 
35|   2   |   3   |   1   |   2   | **0** 
36| **0** |   4   |   2   |   3   |   1 
37|   1   | **0** |   3   |   4   |   2 
38|   2   |   1   |   4   |   5   |   3 
39| **0** |   2   |   5   |   6   |   4 
40|   1   |   3   |   6   |   7   |   5 
41|   2   |   4   | **0** |   8   |   6 
42| **0** | **0** |   1   | **0** |   7 
43|   1   |   1   |   2   |   1   |   8 
44|   2   |   2   |   3   |   2   |   9 
45| **0** |   3   |   4   |   3   |   10 
46|   1   |   4   |   5   |   4   | **0** 

<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Attribution-ShareAlike 4.0 International License</a>.
