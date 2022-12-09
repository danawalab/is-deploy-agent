#!/bin/bash

export agent="is-deploy-agent"
echo "${agent}"
export port=$1
echo "${port}"
export version=$2
echo "${version}"
export backup_dir=backup-dir

function execute_agent() {
  # unzip new agent.tar.gz
  echo "tar -zxvf ${agent}-${version}.tar.gz"
  tar -zxvf ${agent}-${version}.tar.gz
  # create backup directory
  if [ ! -e $backup_dir ]; then
    mkdir $backup_dir
  fi
  # move new agent.tar.gz  backup-dir directory
  echo "mv ${agent}-${version}.tar.gz $backup_dir"
  mv ${agent}-${version}.tar.gz $backup_dir
  # run new agent
  echo "/home/isdeploy/$1 $2"
  /home/isdeploy/$1 $2
}

# check port parameter
if [ -z "$1" ]; then
  echo "${port} is empty"
else
  # check agent version parameter
  if [ -z "$2" ]; then
    echo "${agent} ${version} is empty"
  else
    # check agent
    if [ ! -e ${agent} ]; then
      # download agent
      wget https://github.com/danawalab/${agent}/releases/download/${version}/${agent}-${version}.tar.gz
    else
      # download agent
      wget https://github.com/danawalab/${agent}/releases/download/${version}/${agent}-${version}.tar.gz
      # kill old agent
      kill -9 $(ps -ef | grep ${agent} | awk '{print $2}')
      # delete old agent
      rm ${agent}
    fi
    execute_agent ${agent} ${port}
  fi
fi
