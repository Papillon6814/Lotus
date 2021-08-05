FROM teracy/ubuntu:16.04-dind-latest

USER root

RUN apt update
RUN apt -y upgrade
RUN apt -y install sudo
RUN apt-get -y install fish git openssl curl wget nodejs npm unzip
RUN sudo npm install n -g
#RUN sudo n stable
RUN sudo apt purge -y nodejs npm
RUN sudo apt -y install vim protobuf-compiler

RUN useradd -ms /bin/bash newuser

RUN curl https://get.starport.network/starport | bash
RUN sudo mv starport /usr/local/bin/

RUN wget https://dl.google.com/go/go1.15.8.linux-amd64.tar.gz
RUN sudo tar -xvf go1.15.8.linux-amd64.tar.gz   
RUN sudo mv go /usr/local  

ENV GOROOT /usr/local/go 
ENV GOPATH /workspaces/cosmos_tutorial
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

# Install Proto3
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
RUN unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
RUN sudo mv protoc3/bin/* /usr/local/bin/
RUN sudo mv protoc3/include/* /usr/local/include/
RUN sudo rm -rf protoc3

EXPOSE 1317
EXPOSE 12345
EXPOSE 22657
EXPOSE 8080
EXPOSE 4500

RUN sudo service docker start
#docker run -it 66c48aa7e20b --port "8080:8080"