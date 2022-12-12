#!/bin/bash

agent="is-deploy-agent"
port=$1
version=$2
backup_dir=backup-dir

# download agent
wget https://github.com/danawalab/$agent/releases/download/$version/$agent-$version.tar.gz
# kill old agent
kill -9 $(ps -ef | grep $agent | awk '{print $2}')
# delete old agent
rm $agent
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
