
FROM techtaste/dock-c9

LABEL maintainer rchari<techtaste.me>

ENV C9_PATH /ws/
# SERVER_ROOT_DIR
ENV SERVER_PATH /server
ENV HUGO_BIN ${SERVER_PATH}/scripts/hugo_bin
ENV BLOG_PATH=/ws/blog-source
ENV SITE_PATH=/ws/mySite

ENV LOG_FILE=/tmp/server.log

# Add hugo_bin to PATH
ENV PATH=${PATH}:${HUGO_BIN}

ADD server/ /server/
ADD bashrc.sh /tmp/bashrc.sh

RUN mkdir /ws && apt-get update -y \
    && apt-get install hugo -y \
    && apt-get install iproute2 -y \
    && DEBIAN_FRONTEND=noninteractive apt-get install expect -y \
    && cat /tmp/bashrc.sh >> /root/.bashrc && rm /tmp/bashrc.sh \
    && rm /work/proxy* && rm /work/startup.sh \
    && chmod +x $SERVER_PATH/scripts/hugo_bin/* \
    && mkdir -p /root/.c9/plugins/vdockPlugin


EXPOSE 8288 8286

WORKDIR /server/

CMD ["/bin/bash", "/server/startup.sh"]
