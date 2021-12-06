FROM ubuntu:18.04

RUN apt-get update; apt-get install -y curl

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin/kubectl

RUN useradd -ms /bin/bash user
USER user
WORKDIR /home/user

ADD ./automatt /usr/local/bin/automatt
ADD ./BE-reset.sh /home/user/BE-reset.sh

CMD [ "automatt" ]