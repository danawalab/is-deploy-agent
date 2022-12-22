package tests

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestTomcatHealthCheck(t *testing.T) {
	cmd := exec.Command("bash", "-c", "ps -ef | grep /root/tomcat")
	output, _ := cmd.Output()

	fmt.Println(string(output))
}

/*
테스트 결과
=======================================================================================================================
	=== RUN   TestTomcatHealthCheck
	root      6082     1 57 07:32 pts/7    00:00:05 /usr/bin/java -Djava.util.logging.config.file=/root/tomcat1/conf/logging.properties -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager -Djdk.tls.ephemeralDHKeySize=2048 -Djava.protocol.handler.pkgs=org.apache.catalina.webresources -Dorg.apache.catalina.security.SecurityListener.UMASK=0027 -Dignore.endorsed.dirs= -classpath /root/tomcat1/bin/bootstrap.jar:/root/tomcat1/bin/tomcat-juli.jar -Dcatalina.base=/root/tomcat1 -Dcatalina.home=/root/tomcat1 -Djava.io.tmpdir=/root/tomcat1/temp org.apache.catalina.startup.Bootstrap start
	root      6136     1 74 07:32 pts/7    00:00:02 /usr/bin/java -Djava.util.logging.config.file=/root/tomcat2/conf/logging.properties -Djava.util.logging.manager=org.apache.juli.ClassLoaderLogManager -Djdk.tls.ephemeralDHKeySize=2048 -Djava.protocol.handler.pkgs=org.apache.catalina.webresources -Dorg.apache.catalina.security.SecurityListener.UMASK=0027 -Dignore.endorsed.dirs= -classpath /root/tomcat2/bin/bootstrap.jar:/root/tomcat2/bin/tomcat-juli.jar -Dcatalina.base=/root/tomcat2 -Dcatalina.home=/root/tomcat2 -Djava.io.tmpdir=/root/tomcat2/temp org.apache.catalina.startup.Bootstrap start
	root      6259  6254  0 07:32 pts/7    00:00:00 bash -c ps -ef | grep /root/tomcat
	root      6261  6259  0 07:32 pts/7    00:00:00 grep /root/tomcat

	--- PASS: TestTomcatHealthCheck (0.00s)
	PASS
	ok      command-line-arguments  0.005s
=======================================================================================================================
exec.Command는 파이프라인 '|' 이 적응 안됨 그래서 지금위 테스트 코드와 같이 사용하거나 따로 파이프라인을 만들어 줘야함

root      6259  6254  0 07:32 pts/7    00:00:00 bash -c ps -ef | grep /root/tomcat
root      6261  6259  0 07:32 pts/7    00:00:00 grep /root/tomcat
이 두줄을 빼거나 없애는 방법? tomcat 프로세스 여부 확인에 방해 됨
*/
