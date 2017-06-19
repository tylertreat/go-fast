FROM golang

RUN apt-get -y update
RUN apt-get -y install cmake
RUN apt-get -y install numactl
RUN apt-get -y install r-base
RUN apt-get -y install ruby-full
RUN gem install specific_install
RUN gem specific_install -l https://github.com/jonhoo/experiment
RUN go get github.com/jonhoo/drwmutex/drwmutex-bench
ADD experiment.json /go/
