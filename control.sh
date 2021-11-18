#!/bin/bash

workspace=$(cd $(dirname $0); pwd)
cd $workspace

mkdir -p logs

app=wax-engine
conf=conf/cfg.example.yaml
pidfile=logs/app.pid
logfile=logs/app.log
sleep_time=3 # unit: second

function check_pid() {
  if [ -f $pidfile ]; then
    pid=`cat $pidfile`
    if [ -n $pid ]; then
      running=`ps -p $pid | grep -v "PID TTY" | wc -l`
      return $running
    fi
  fi
  return 0
}

function build() {
  go build -o $app
  if [ $? -ne 0 ]; then
    exit $?
  fi
  echo "succeed to build app as $app"
  ./$app -v
}

function pack() {
  build
  version=`./$app -v`
  file_list="control.sh conf $app"
  tar czf $app-$version.tar.gz $file_list
}

function start() {
  check_pid
  running=$?
  if [ $running -gt 0 ]; then
    echo -n "$app is running already, pid="
    cat $pidfile
    return 1
  fi
  echo "use the config file: $conf"
  nohup $workspace/$app -c $conf > $logfile 2>&1 &
  echo $! > $pidfile

  # check process after started
  echo "wait $sleep_time seconds to check process..."
  sleep $sleep_time
  check_pid
  running=$?
  if [ $running -le 0 ]; then
    echo -n "failed to start $app, see $logfile"
    echo
  else
    echo "succeed to start $app, pid=$!"
  fi
}

function stop() {
  check_pid
  running=$?
  if [ $running -le 0 ]; then
    echo "$app is already stopped"
    return
  fi

  pid=`cat $pidfile`
  kill -9 $pid
  echo "$app stopped..."
}

function restart() {
  stop
  sleep 1
  echo "try to restart $app..."
  start
}

function status() {
  check_pid
  running=$?
  if [ $running -gt 0 ]; then
    echo -n "$app is running, pid="
    cat $pidfile
  else
    echo "$app is stopped"
  fi
}

function help() {
  echo "$0 build|pack|start|stop|restart|status"
}

if [ "$1" == "" ]; then
  help
elif [ "$1" == "stop" ]; then
  stop
elif [ "$1" == "build" ]; then
  build
elif [ "$1" == "pack" ]; then
  pack
elif [ "$1" == "start" ]; then
  start
elif [ "$1" == "stop" ]; then
  stop
elif [ "$1" == "restart" ]; then
  restart
elif [ "$1" == "status" ]; then
  status
else
  help
fi