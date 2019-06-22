package zex

import (
	"github.com/metaleap/zentient"
)

func OnPreInit() error {
	z.Lang.ID, z.Lang.Title, z.Lang.Enabled =
		"examplelang", "ExampleLang", false // need true when your code has verified that your lang is installed on the user's machine, else also return an err:
	return nil
}

func OnPostInit() {
}

/*
// your executable would then be:
package main

import (
	"github.com/metaleap/zentient"
	"github.com/metaleap/zentient/lang/examplelang"
)

func main() {
	z.InitAndServeOrPanic(zex.OnPreInit, zex.OnPostInit)
}

*/
