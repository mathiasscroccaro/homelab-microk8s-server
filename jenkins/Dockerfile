FROM jenkins/jenkins:lts
USER root
RUN apt-get update -qq \
    && apt-get install -qqy curl
# RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
# RUN add-apt-repository \
#    "deb [arch=amd64] https://download.docker.com/linux/debian \
#    $(lsb_release -cs) \
#    stable"
# RUN apt-get update  -qq \
#     && apt-get -y install docker-ce
# RUN usermod -aG docker jenkins

# RUN curl https://get.docker.com/ > dockerinstall && chmod 777 dockerinstall && ./dockerinstall

RUN curl -fsSL https://get.docker.com -o get-docker.sh

RUN sh get-docker.sh

# RUN update-alternatives --set iptables /usr/sbin/iptables-legacy
# RUN update-alternatives --set ip6tables /usr/sbin/ip6tables-legacy

RUN usermod -aG docker jenkins

RUN systemctl enable docker
