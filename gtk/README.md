## mac install
###1.
brew install gtk+
或者brew install gtk+3
###2.
pkg-config --libs --cflags gtk+-2.0
或者pkg-config --libs --cflags gtk+-3.0
###3.
clang `pkg-config --libs --cflags gtk+-3.0` -O3 -g0 -s -o gtk gtk.c

## 后备
clang 'pkg-config --libs --cflags gtk+-3.0' main.c

##glade
brew install glade

http://www.baidu.com/link?url=q3_pV11Y945VVONi-j0b1-xkKgTNReVb5oV8YfAWCc4BmEveD2hVkhkBqPvXqpADl9GxDJx-hvMjEJAA6a0WW_23WMwSO-7G2A0J-gzbn8O&wd=&eqid=f565bd1000158160000000035eb79fd0
