#!/bin/bash

port=5000
agent="is-deploy-agent"

# kill old agent
kill -9 $(ps -ef | grep $agent | awk '{print $2}')

# delete old agent
rm $agent

# unzip new agent.tar.gz
tar -zxvf $agent.tar.gz

# delete new agent.tar.gz
rm $agent.tar.gz

# run new agent
./$agent $port