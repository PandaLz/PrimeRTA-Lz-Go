# Process for the sieving method
## Based on PrimeRT-Lz
https://github.com/PandaLz/PrimeRT-Lz

### Method 2
#### Find sieve of (x * 2) ^ 2
#### Where ((x * 2) ^ 2) + r when r is odd, is prime if p is true
#### And where (((x * 2) + 1) ^ 2) + r when r is pair, is prime if p is true
#### When x > 3

#### Example value x = 5

Step 1. Start by define the values needed for the procedure based on x.
* Min of odd numbers to test, Min = 3
+ Max of odd numbers to test, Max = (x * 2) + 1
* Control number, C = x - 2

Current values

x = 5

Min = 3

Max = 11

C = 3

Step 2. Find total size of the sieve T and generate bool array default to true
* Starting point, S = (C * (C + 1)) * 2
* Range of values for r, R = Max * 2
* Total size of the sieve, T = S + R

S = 24

R = 22

T = 46

Table 1

T |   p   |
--|-------|
0 | true  |
1 | true  |
2 | true  |
3 | true  |
4 | true  |
5 | true  |
6 | true  |
7 | true  |
8 | true  |
9 | true  |
10| true  |
11| true  |
12| true  |
13| true  |
14| true  |
15| true  |
16| true  |
17| true  |
18| true  |
19| true  |
20| true  |
21| true  |
22| true  |
23| true  |
24| true  | 
25| true  |
26| true  |
27| true  |
28| true  |
29| true  |
30| true  |
31| true  |
32| true  |
33| true  |
34| true  | 
35| true  | 
36| true  |
37| true  |
38| true  |
39| true  |
40| true  |
41| true  |
42| true  |
43| true  |
44| true  |
45| true  | 
46| true  | 

Step 3. Start sieving through the table starting from Min up to Max in steps of 2 marking cell as false every nth steps
* In this case Min 3, 5, 7, 9, 11 Max

But each step after 3 will Switch starting point this way:
* Know new start point for each iteration starting from Z = 0
* Next iteration will be Z = Z + (C * 4)
* Updating C per iteration to C = C - 1
* In this case:
* When 3, Z = 0
* When 5, Z = Z + (3 * 4) = 0 + 12 = 12 then C - 1 = 2
* When 7, Z = Z + (2 * 4) = 12 + 8 = 20 then C - 1 = 1
* When 9, Z = Z + (1 * 4) = 20 + 4 = 24 then C - 1 = 0
* When 11, Z = Z + (0 * 4) = 20 + 0 = 24 then end loop

Table 2

T |   p   |Min2Max|
--|-------|-------|
0 | false | Z = 0 |
1 | true  |       |
2 | true  |       |
3 | false |       |
4 | true  |       |
5 | true  |       |
6 | false |       |
7 | true  |       |
8 | true  |       |
9 | false |       |
10| true  |       |
11| true  |       |
12| false | Z = 12|
13| true  |       |
14| true  |       |
15| false |       |
16| true  |       |
17| false |       |
18| false |       |
19| true  |       |
20| false | Z = 20|
21| false |       |
22| false |       |
23| true  |       |
24| false | Z = 24|
25| true  |       |
26| true  |       |
27| false |       |
28| true  |       |
29| true  |       |
30| false |       |
31| true  |       |
32| false |       |
33| false |       |
34| false |       |
35| false |       |
36| false |       |
37| false |       |
38| true  |       |
39| false |       |
40| true  |       |
41| false |       |
42| false |       |
43| true  |       |
44| true  |       |
45| false |       |
46| false |       |

Step 4. Find odd and pair values of r
* Odd part start next to the current Z value O = Z + 1
* In this case O = 25, which will be set as 1 and increase in steps of 2
* Until reach a value of (x * 4) - 1 next cell after reach that value
* will be the new Z and set to 0
* Pair part is found next to the new Z value P = Z + 1
* In this case Z = 35, P = 36 which will be set as 2 and increase in steps of 2
* until reach a value of (x * 4) + 2 which should be the end of the table

Finally can apply:

((x * 2) ^ 2) + r for odd values of r or

This case be ((5 * 2) ^ 2) + r = 100 + r

(((x * 2) + 1) ^ 2) + r for pair values of r where the row is true

This case be (((5 * 2) + 1) ^ 2) + r = 121 + r

Table 3

T |   p   |Min2Max|   r   | prime
--|-------|-------|-------|------
0 | false | Z = 0 |       |
1 | true  |       |       |
2 | true  |       |       |
3 | false |       |       |
4 | true  |       |       |
5 | true  |       |       |
6 | false |       |       |
7 | true  |       |       |
8 | true  |       |       |
9 | false |       |       |
10| true  |       |       |
11| true  |       |       |
12| false | Z = 12|       |
13| true  |       |       |
14| true  |       |       |
15| false |       |       |
16| true  |       |       |
17| false |       |       |
18| false |       |       |
19| true  |       |       |
20| false | Z = 20|       |
21| false |       |       |
22| false |       |       |
23| true  |       |       |
24| false | Z = 24|   0   |
25| true  |       |   1   |100 + 1 = 101
26| true  |       |   3   |100 + 3 = 103
27| false |       |   5   |
28| true  |       |   7   |100 + 7 = 107
29| true  |       |   9   |100 + 9 = 109
30| false |       |  11   |
31| true  |       |  13   |100 + 13 = 113
32| false |       |  15   |
33| false |       |  17   |
34| false |       |  19   |
35| false |       |   0   |
36| false |       |   2   |
37| false |       |   4   |
38| true  |       |   6   |121 + 6 = 127
39| false |       |   8   |
40| true  |       |  10   |121 + 10 = 131
41| false |       |  12   |
42| false |       |  14   |
43| true  |       |  16   |121 + 16 = 137
44| true  |       |  18   |121 + 18 = 139
45| false |       |  20   |
46| false |       |  22   |


--------------------------------------------------------------------------------------

### Method 1
#### Find sieve of (x * 2) squared

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

Now can discard the rows 25 and above that contains 0 values as not primes only left is to determine the prime number from value x, observe last column of the left only contains two 0 values, the middle 0 will define the process to use

From values below middle 0:
First get (x * 2) squared, then take the value of last column * 2 - 1 and add both values.

From values above middle 0:
Get ((x * 2) + 1) squared, then take the value of last column * 2 and add both values.

H |   3   |   5   |   7   |   9   |   11  | prime |
--|-------|-------|-------|-------|-------|-------|
25|   1   |   3   |   5   |   1   |   1   |  101  |
26|   2   |   4   |   6   |   2   |   2   |  103  |
27| **0** | **0** | **0** |   3   |   3   |       |
28|   1   |   1   |   1   |   4   |   4   |  107  |
29|   2   |   2   |   2   |   5   |   5   |  109  |
30| **0** |   3   |   3   |   6   |   6   |       |
31|   1   |   4   |   4   |   7   |   7   |  113  |
32|   2   | **0** |   5   |   8   |   8   |       |
33| **0** |   1   |   6   | **0** |   9   |       |
34|   1   |   2   | **0** |   1   |   10  |       |
35|   2   |   3   |   1   |   2   | **0** |       |
36| **0** |   4   |   2   |   3   |   1   |       |
37|   1   | **0** |   3   |   4   |   2   |       |
38|   2   |   1   |   4   |   5   |   3   |  127  |
39| **0** |   2   |   5   |   6   |   4   |       |
40|   1   |   3   |   6   |   7   |   5   |  131  |
41|   2   |   4   | **0** |   8   |   6   |       |
42| **0** | **0** |   1   | **0** |   7   |       |
43|   1   |   1   |   2   |   1   |   8   |  137  |
44|   2   |   2   |   3   |   2   |   9   |  139  |
45| **0** |   3   |   4   |   3   |   10  |       |
46|   1   |   4   |   5   |   4   | **0** |       |


<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Attribution-ShareAlike 4.0 International License</a>.
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
