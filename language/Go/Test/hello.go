package main

import(
	"fmt"
)
func qsort(values []int,left int,right int){
	temp :=values[left]
	p:=left
    i,j:=left,right

    for i<=j{
        for j>=p && values[j]>=temp{
            j--
        }
        if j>=p{
            values[p]=values[j]
            p=j
        }

        for i<=p && values[i]<=temp{
            i++
        }
        if i<=p{
            values[p]=values[i]
            p=i
		}
	}
	values[p]=temp

    if p-left>1{
        qsort(values,left,p-1)
    }
    if right-p>1{
        qsort(values,p+1,right)
    }
}
func qqsort(values []int){
	qsort(values,0,len(values)-1)
}
func main() {
	a:=[]int{4,2,5,7,1,1}
	qqsort(a)
	fmt.Println(a)
}

