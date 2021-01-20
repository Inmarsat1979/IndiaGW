#!/bin/sh

#Find the Process ID for syncapp running instance

PID=`ps -ef | grep sac_controller 'awk {print $2}'`

if [[ -z "$PID" ]] then
Kill -9 PID
fi
