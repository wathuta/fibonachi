package main

import "fmt"

func main()  {
		fmt.Println(MemFib(5))
	
}

func fibWrapper(n int)[]int{
	retArr:=[]int{}
	for i:=0;i<n;i++{
		// fmt.Println(fibSpace(i))
		retArr=append(retArr,fibonnacci(i))
	}
	return retArr
}
func fibonnacci(n int)int{
	if n<2{
		return n
	}
	return fibonnacci(n-1)+fibonnacci(n-2)
}
//dynamic programming
func fibo(n int)[]int{
	retArr:= []int{}
	retArr = append(retArr, 0)
	if n==1{
		return retArr
	}
	retArr = append(retArr, 1)
	if n==2 {
		return retArr
	}
	

	for i := 0; i < n-2; i++ {
		retArr=append(retArr,  retArr[i]+retArr[i+1])
	}
	return retArr
}
//with a better space complexity
func fibSpace(n int)[]int{
	retArr:=[]int{}
	var a,b,result int
	b=1
	for i:=0;i<n;i++{
		retArr = append(retArr, a)
		result=a+b
		a=b
		b=result
	}
	return retArr
}
func MemFib(n int)int  {
	Cache:=make(map[int]int)

	return helper(n,Cache)
}
func helper(n int,Cache map[int]int)int  {
	Cache[1]=0
	Cache[2]=1
	if n<3{
		return Cache[n]
	}
	Cache[n]=helper(n-1,Cache)+helper(n-2,Cache)
	return Cache[n]
}