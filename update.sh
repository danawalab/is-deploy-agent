#!/bin/bash

agent="is-deploy-agent"
port=$1
version=$2
backup_dir=backup-dir

function execute_agent() {
  # unzip new agent.tar.gz
  tar -zxvf $agent-$version.tar.gz
  # create backup directory
  if [ ! -e $backup_dir ]; then
    mkdir $backup_dir
  fi
  # move new agent.tar.gz  backup-dir directory
  mv $agent-$version.tar.gz $backup_dir
  # run new agent
  ./$agent $port
}

# check port parameter
if [ -z "$1" ]; then
  echo "port is empty"
else
  # check agent version parameter
  if [ -z "$2" ]; then
    echo "agent version is empty"
  else
    # check agent
    if [ ! -e $agent ]; then
      # download agent
      wget https://github.com/danawalab/$agent/releases/download/$version/$agent-$version.tar.gz
    else
      ./update2.sh $port $version
    fi
    execute_agent
  fi
fi