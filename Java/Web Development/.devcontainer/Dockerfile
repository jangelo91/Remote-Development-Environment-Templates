ARG VERSION="11-buster"
# ARG HOME=""



FROM openjdk:${VERSION}

ARG INSTALL_MAVEN="false"
ARG MAVEN_VERSION=3.6.3
ARG INSTALL_GRADLE="false"
ARG GRADLE_VERSION=5.4.1
# RUN if [ "${INSTALL_MAVEN}" = "true" ]; then vscode -c "source /usr/local/sdkman/bin/sdkman-init.sh && sdk install maven \"${MAVEN_VERSION}\""; fi \
    # && if [ "${INSTALL_GRADLE}" = "true" ]; then vscode -c "source /usr/local/sdkman/bin/sdkman-init.sh && sdk install gradle \"${GRADLE_VERSION}\""; fi

# FROM springio/spring-ci-base

# RUN apt-get update
#RUN apt-get install -y openjdk-11-jdk 



ARG USERNAME=vscode

# COPY . .

WORKDIR /