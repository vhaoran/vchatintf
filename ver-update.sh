#!/bin/bash

up(){
   ver="$1";
   file="$2";
   echo "------------->""$file";

   sed -E -i '' "s/github\.com\/vhaoran\/vchat .*/github\.com\/vhaoran\/vchatintf $ver/g"  "$file";
   echo "-------after update----------------"
   grep "vchatintf"  "$file"


};


VER=$(git ls-remote --heads |awk '{ print $1}')


up "$VER"  "$GOPATH/src/vchatuser/go.mod";
up "$VER"  "$GOPATH/src/vchatutil/go.mod";
up "$VER"  "$GOPATH/src/vchatmsg/go.mod";
up "$VER"  "$GOPATH/src/vchatfinance/go.mod";
up "$VER"  "$GOPATH/src/vchatgw/go.mod";
up "$VER"  "$GOPATH/src/vchatws/go.mod";





echo "$VER"
