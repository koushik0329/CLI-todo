package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func Newcmd() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a todo with title")
	
	flag.IntVar(&cf.Del, "del", -1, "delete a todo")
	flag.StringVar(&cf.Edit, "edit","", "edit a todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "specify a todo to mark completed")
	flag.BoolVar(&cf.List, "list", false, "list all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos){
	switch{
	case cf.List:
		todos.print()
	case cf.Add!="":
		todos.add(cf.Add)
	
	case cf.Edit!="":
		split:=strings.SplitN(cf.Edit,":", 2)
		if(len(split)!=2){
			fmt.Println("Error")
			os.Exit(1)
		}
		index, err:=strconv.Atoi(split[0])
		if err!=nil{
			fmt.Println("Error")
			os.Exit(1)
		}
		todos.edit(index,split[1])
	case cf.Del!=-1:
		todos.delete(cf.Del)
	case cf.Toggle!=-1:
		todos.toggle(cf.Toggle)	
	default:
		fmt.Println("Enter a valid command")
	}
}