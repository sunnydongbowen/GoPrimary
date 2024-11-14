/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-12 14:37:59
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-13 11:25:28
 * @FilePath: /GoPrimary/qimi/day06/errror_test.go
 * @Description:
 *
 */
package main

import (
    "fmt"
    "testing"
    "os"
	"errors"
)

var (
	ErrInvalOp = errors.New("invalid Operation")
	ErrInvalId = errors.New("invalid Id")
)

func TestError(t *testing.T) {
   f,err:=os.Open("xxx.txt")
   if err!=nil{
       fmt.Println(err)
   }
    defer f.Close()
}


type Order struct{}

func queryOrder(id int) (*Order,error) {
	//
	if err:=ConnectDB("jj","123");err!=nil{
		//return nil,err
		// base on err，get new err
		return nil,fmt.Errorf("connect to database error:%w",err)
	}

	if s:="can not connect database";len(s)>0{
		return nil,errors.New(s)
	}

	// can connected to database .but cannot find order
	if false{
		//return nil,errors.New("can not find order")
		return nil,ErrInvalId
	}
	if err := ErrInvalOp;err!=nil{
		return nil,err
	}
	// success
	return &Order{},nil
}




func TestError2(t *testing.T) {
	_,err:=queryOrder(1)
	oErr:=errors.Unwrap(err)
	fmt.Println(oErr,oErr==ErrConnectDB)
	
	if ok:=errors.Is(err,ErrInvalId);ok{
		fmt.Println("is invalid id")
	}
	if ok:=errors.Is(err,ErrConnectDB);ok{
		fmt.Println("is connect db error")
	}

    //
	var nErr *DBError
	if ok:=errors.As(err,&nErr);ok{
		fmt.Println(nErr)
	}

	if err!=nil{
		fmt.Println(err)
	}
	
    
}

