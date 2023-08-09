package main

import "fmt"
type node struct{
	prev *node
	val int
	 next *node

}

var head *node=nil;//these three are global variables
var tail *node=nil;
var count =0
var a int =0

func addfront(num int){
	count++
	front:=&node{ nil, num, nil}
	if tail==nil{
		head=front
		tail=front
		return
	}
    front.next=head
	head.prev=front
	head=front
}

func insert(num,pos int){
	var th *node =head
	if(pos<=0){
		addfront(num)
		return
	}
	if(pos>count){
		addback(num)
		return
	}
	count++

	for ;pos>1; {
        th=th.next
		pos--
	}
    temp:=&node{nil,num,nil}
	temp.prev=th
	temp.next=th.next
	th.next=temp
	temp.next.prev=temp//if we use like this we can able to change temp variables next address previous variable
}

func addback(num int){
	count++
	back:=&node{nil,num,nil}
	if tail==nil{
		head=back
		tail=back
		return
	}
	back.prev=tail
	tail.next=back
	tail=back
}

func deletefront(){
   count--
   if(head==nil){
	fmt.Println("the list is empty")
	return
   }
   if(head.next==nil){
	head=nil
	tail=nil
	return
   }
   head=head.next
   head.prev=nil
}

func deleteback(){
	count--
	if(head==nil){
		fmt.Println("the list is empty")
		return
	   }
	   if(head.next==nil){
		head=nil
		tail=nil
		return
	   }
	   tail=tail.prev
	   tail.next=nil
}

func deleterandom(pos int){
	if(pos==0){
		deletefront()
		return
	   }
	   if(pos==count){
		deleteback()
		return
	   }
	count--
	if(head==nil){
		fmt.Println("the list is empty")
		return
	   }
	   if(head.next==nil){
		head=nil
		tail=nil
		return
	   }
	   var th *node=head
	   
	   for ;pos>0; {
             th=th.next
			 pos--
	   }
	   th.prev.next = th.next
	   th.next.prev=th.prev
}
 func fronttraversal(){
	if(head==nil){
		fmt.Println("list is empty")
		return
	}
       var th *node=head
	   for i:=0;i<count;i++{
          fmt.Println(th.val)
		  th=th.next
	   }
 }


 func backtraversal(){
	if(head==nil){
		fmt.Println("list is empty")
		return
	}
	var th *node=tail
	for i:=0;i<count;i++{
	   fmt.Println(th.val)
	   th=th.prev
	}
 }

func main(){
for true{

 fmt.Println("choose from the options given below\n")
 //fmt.Printf("\n1.Insert in begining\n2.Insert at last\n3.Insert at any random location\n4.Delete from Beginning\n  
 //5.Delete from last\n6.Delete the node after the given data\n7.front traversal\n8.back travesal\n9.Exit\n")
 var choice int
 fmt.Scanln(&choice)
 switch(choice){
 case 1:{
	var num int
	fmt.Println("enter the value to be added in the first")
	fmt.Scanln(&num)
    addfront(num)
 }
case 2:{
	var num int
	fmt.Println("enter the value to be added in the last")
	fmt.Scanln(&num)
    addback(num)
}

case 3:{
	var pos int
	var num int
	fmt.Println("enter the value to be added in the random position")
	fmt.Scanln(&num)
	fmt.Scanln(&pos)
    insert(num,pos)
}
case 4:{
	
    deletefront()
}
case 5:{
	
    deleteback()
}
case 6:{
	var pos int
	fmt.Println("enter the value to be delete in the random")
	fmt.Scanln(&pos)
    deleterandom(pos)
}
case 7:{
	
    fronttraversal()
}
case 8:{
	
    backtraversal()
}
case 9:{
	fmt.Println(count)
}
default:{
	a=1
}
 }
 if(a==1){
	break;
 }
}
}
