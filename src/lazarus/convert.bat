@echo off
echo %gopath%
set path=%path%;D:\Language\gopath\pkg\mod\github.com\ying32\govcl@v2.0.6+incompatible\Tools\res2go\src
res2go.exe -path .\ -outpath ..\gui -pkgname gui -outmain false