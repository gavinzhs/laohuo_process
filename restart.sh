#!/usr/bin/env bash
host=123.56.135.128
slow_log=150

#an xin jie sms is 0, ronglian sms is 1
SMS=0

ulimit -n 20000
fuser -k dizhi
MARTINI_ENV=production GOMAXPROCS=4 nohup ./dizhi -p ":80" -tok ":7243" -h $host -slow_log $slow_log -sms $SMS &
