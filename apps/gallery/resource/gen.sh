#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package resource -name christmasBitmap ./img/christmasBitmap.png > $FILE
$BIN/fyne bundle -package resource -append -name galleryBitmap ./img/gallery.svg >> $FILE


