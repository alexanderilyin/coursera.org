package main

import (
	"fmt"
)

//
// Program to test running functions concurrently
// Function that can be called as a Go Routine
// Passes a message for tracking
//
// Trying to casuse a Race Condition
//
/*                     RESULTS
First Run with no Concurance results are all counters are Sequential No Race Conditions
=======================================================================================

1st Call 100 (Add Red Stars, 101 Executions)
***:0 xxx-2001 ***:1 xxx-2002 ***:2 xxx-2003 ***:3 xxx-2004 ***:4 xxx-2005 ***:5 xxx-2006 ***:6 xxx-2007 ***:7 xxx-2008 ***:8 xxx-2009 ***:9 xxx-2010
2nd Call 100 ( All 1 Colour, 101 Executions)
...:0 xxx-2011 ...:1 xxx-2012 ...:2 xxx-2013 ...:3 xxx-2014 ...:4 xxx-2015 ...:5 xxx-2016 ...:6 xxx-2017 ...:7 xxx-2018 ...:8 xxx-2019 ...:9 xxx-2020

End of Program xxx =  2020 xxxLast1  2020  xxxLast2  2020 - Results as expected
                      ----           ----            ----   ===================

Second Run with 1 Function call (go) Concurance results are all counters are Not Sequential - Race Condition Occurs End of Program Results are not Predicatable
==========================================================================================   ============== ===================================================

1st Call 100

2nd Call 100 ( 101 Executions)
...:0 xxx-1999 ***:0 xxx-2000 ***:1 xxx-2002 ...:1 xxx-2001 ...:2 xxx-2004 ...:3 xxx-2005 ***:2 xxx-2003 ***:3 xxx-2007 ...:4 xxx-2006 ...:5 xxx-2009 ...:6 xxx-2010 ***:4 xxx-2008 ***:5 xxx-2012 ...:7 xxx-2011 ...:8 xxx-2014 ...:9 xxx-2015
 Race Condition ************************************ XXX 2015 xxxLast1  1998 xxxLast2  2015

End of Program xxx =  2015 xxxLast1  2008  xxxLast2  2015
                      ????           ????            ????

Third Run with 2 Function call (go) Concurance results are all counters are Not Sequential - Race Condition Occurs End of Program Results are not Predicatable
==========================================================================================   ============== ===================================================

Race Condition ************************************ XXX 1867 xxxLast1  1715 xxxLast2  1867   ( MULTIPLE RACE Conditions, and you can clearly see the non sequential execution of the two functions

...:7 xxx-1624 ...:8 xxx-1869 ...:9 xxx-1870 ***:3 xxx-1726 ***:4 xxx-1871 ***:5 xxx-1872 ***:6 xxx-1873 ***:7 xxx-1874 ***:8 xxx-1875 ***:9 xxx-1876 ***:0 xxx-1674 ***:1 xxx-1877 ***:2 xxx-1878 ***:3 xxx-1879 ...:5 xxx-1832 ...:6 xxx-1881 ...:7 xxx-1882 ...:8 xxx-1883 ***:8 xxx-1860 ***:9 xxx-1885 ...:3 xxx-1778 ...:4 xxx-1886 ...:0 xxx-1653 ...:1 xxx-1888 ...:2 xxx-1889 ...:3 xxx-1890 ...:4 xxx-1891 ...:5 xxx-1892 ...:6 xxx-1893 ...:7 xxx-1894 ...:8 xxx-1895 ...:9 xxx-1896 ...:5 xxx-1793 ...:6 xxx-1897 ...:7 xxx-1898 ...:8 xxx-1899 ...:9 xxx-1900 ***:8 xxx-1738 ***:9 xxx-1901 ***:4 xxx-1880 ***:5 xxx-1902 ***:0 xxx-1582 ***:1 xxx-1904 ***:2 xxx-1905 ...:3 xxx-1651 ...:4 xxx-1907 ...:5 xxx-1908 ...:6 xxx-1909 ...:7 xxx-1910 ...:8 xxx-1911 ...:9 xxx-1912
1st Call 100

2nd Call 100
***:5 xxx-1845 ***:6 xxx-1913 ***:7 xxx-1914 ***:0 xxx-1916 ***:1 xxx-1918 ***:2 xxx-1919 ***:3 xxx-1920 ***:4 xxx-1921 ***:5 xxx-1922 ***:6 xxx-1923 ***:7 xxx-1924 ***:8 xxx-1925 ...:4 xxx-1852 ***:7 xxx-1804 ...:0 xxx-1676 ***:6 xxx-1827 ***:7 xxx-1927 ***:8 xxx-1928 ***:9 xxx-1929 ...:9 xxx-1867 ...:6 xxx-1590 ...:7 xxx-1930 ...:8 xxx-1931 ...:9 xxx-1932 ***:8 xxx-1917 ...:7 xxx-1749 ...:1 xxx-1935 ...:2 xxx-1936 ***:8 xxx-1937 ***:9 xxx-1939 ...:9 xxx-1884 ***:6 xxx-1903 ***:7 xxx-1940 ***:8 xxx-1941 ***:9 xxx-1942 ...:2 xxx-1699 ...:3 xxx-1943 ...:0 xxx-1868 ...:1 xxx-1945 ...:2 xxx-1946 ...:3 xxx-1947 ...:4 xxx-1948 ...:5 xxx-1949 ...:6 xxx-1950 ...:7 xxx-1951 ...:8 xxx-1952 ...:9 xxx-1953 ***:2 xxx-1751 ***:3 xxx-1954 ***:4 xxx-1955 ***:5 xxx-1956 ***:6 xxx-1957 ***:7 xxx-1958 ...:5 xxx-1959 ...:6 xxx-1961 ...:7 xxx-1962 ...:8 xxx-1963

End of Program xxx =  1912 xxxLast1  1922  xxxLast2  1912

...:8 xxx-1935

Also, If you execute it multiple times you get different results.

*/
var xxx int = 1000
var xxxlast1 int = 0
var xxxlast2 int = 0

//
// Function to recieve a string and
// Display diffent messages per call
//
// This function is used to illustrate Race Conditions
//

func Concurrent(instr string) {
	for i := 0; i < 10; i++ {
		xxx++
		fmt.Print(instr, ":", i, " xxx-", xxx, " ")
	}
}

func main() {

	/*
	  Loop to control multiple chances of a

	*/

	xxx = 0

	for i := 0; i <= 100; i++ {

		// Suppose we have a function call `f(s)`. Here's how
		// we'd call that in the usual way, running it
		// synchronously.

		fmt.Print("\n\033[31;1;4m1st Call\033[0m ", i, "\n")
		go Concurrent("\033[31;1;4m***\033[0m")
		xxxlast1 = xxx

		/*
			Run a Go Routine, so we have an opportunity
			for a Race Condition
		*/
		fmt.Print("\n2nd Call ", i, "\n")
		go Concurrent("...")
		xxxlast2 = xxx

		if xxxlast2 > xxxlast1+10 {
			fmt.Println("\n Race Condition ************************************ XXX", xxx, "xxxLast1 ", xxxlast1, "xxxLast2 ", xxxlast2)
		}

	}

	fmt.Println("\n\nEnd of Program xxx = ", xxx, "xxxLast1 ", xxxlast1+10, " xxxLast2 ", xxxlast2, "\n")
}
