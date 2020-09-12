FROM golang:1.15.1-buster

RUN groupadd -g 800 jenkins \
  && useradd -m -g 800 -u 800 jenkins

RUN apt-get update \
    && apt-get install --no-install-recommends -y \
        zip \
    && rm -rf /var/lib/apt/lists/*

USER jenkins
