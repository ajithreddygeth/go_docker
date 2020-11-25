Learn go docker 



DOCKER:

Build:

docker build -t mywebsite:<version>

RUN:

docker run -p 8081:8081 mywebsite:<version>


JENKINS:

dockr build -t  ajithreddydocker/jenkinsallinone:1.0 .

docker run -d --name=myJenkins2 -v myVol1:/var/jenkins_home -p 8181:8080 -p 50001:50000 ajithreddydocker/jenkinsallinone:1.0