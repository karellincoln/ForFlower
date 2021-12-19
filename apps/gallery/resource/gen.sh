#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package resource -name love0 ./img/love0.jpeg > $FILE
$BIN/fyne bundle -package resource -append -name love1 ./img/love1.jpeg >> $FILE
$BIN/fyne bundle -package resource -append -name love2 ./img/love2.jpeg >> $FILE
$BIN/fyne bundle -package resource -append -name love3 ./img/love3.jpeg >> $FILE

