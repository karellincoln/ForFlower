#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
TTFFILE=bundleTff.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE $TTFFILE


$BIN/fyne bundle -package resource -name songTtf ./ttf/ShangShouJianSongXianXiTi-2.ttf > $TTFFILE
$BIN/fyne bundle -package resource -append -name boldSongTtf ./ttf/No.16-ShangShouRuiYuanTi-2.ttf >> $TTFFILE


$BIN/fyne bundle -package resource -name iconBitmap ./img/icon.png > $FILE
$BIN/fyne bundle -package resource -append -name galleryBitmap ./img/gallery.svg >> $FILE
$BIN/fyne bundle -package resource -append -name answerBitmap ./img/answer.jpeg >> $FILE
$BIN/fyne bundle -package resource -append -name infoBitmap ./img/info.jpeg >> $FILE

