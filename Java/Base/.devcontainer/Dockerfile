# Specifies what version of Java to use
ARG VERSION
# Specifies whether to install Maven
ARG INSTALL_MAVEN 
# Specifies whether to install Gradle
ARG INSTALL_GRADLE

# Full list of tags can be found at https://github.com/docker-library/docs/blob/master/openjdk/README.md#supported-tags-and-respective-dockerfile-links
FROM openjdk:${VERSION}

# Specify whether container comes with build tools (optional)
RUN if [ "${INSTALL_MAVEN}" = "true" ]; then su vscode -c "source /usr/local/sdkman/bin/sdkman-init.sh && sdk install maven \"${MAVEN_VERSION}\""; fi \
    && if [ "${INSTALL_GRADLE}" = "true" ]; then su vscode -c "source /usr/local/sdkman/bin/sdkman-init.sh && sdk install gradle \"${GRADLE_VERSION}\""; fi
