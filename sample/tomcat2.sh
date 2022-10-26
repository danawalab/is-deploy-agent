#!/bin/bash

echo "Tomcat1 Shell Script Start"
echo $PWD

echo "Move Tomcat2 bin dir"
cd ~/tomcat2/bin
echo $PWD

echo "Tomcat2 stop"
sh ./shutdown.sh

echo "Move New WAR dir"
cd /home/is-deploy-agent/sample
echo $PWD

echo "Copy Tomcat2.war to Tomcat2 webapp"
cp tomcat2.war ~/tomcat2/webapps

echo "Move Tomcat1 bin dir"
cd ~/tomcat2/bin
echo $PWD

echo "Tomcat2 start"
sh ./startup.sh