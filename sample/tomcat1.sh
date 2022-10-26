#!/bin/bash

echo "Tomcat1 Shell Script Start"
echo $PWD

echo "Move Tomcat1 bin dir"
cd ~/tomcat1/bin
echo $PWD

echo "Tomcat1 stop"
sh ./shutdown.sh

echo "Move New WAR dir"
cd /home/is-deploy-agent/sample
echo $PWD

echo "Copy Tomcat1.war to Tomcat1 webapp"
cp tomcat1.war ~/tomcat1/webapps

echo "Move Tomcat1 bin dir"
cd ~/tomcat1/bin
echo $PWD

echo "Tomcat1 start"
sh ./startup.sh