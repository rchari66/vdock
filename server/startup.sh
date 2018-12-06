#!/bin/bash

# Download vdockPlugin for adding custom snippets
git clone https://github.com/rchari66/vdock-plugin-cloud9.git ${C9_PATH}vdockPlugin
cp -r cp -r ${C9_PATH}vdockPlugin /root/.c9/plugins/

# start the cloud9
if [ -z "${AUTH}" ]
then
    # without auth
    echo "Starting c9 in no_authentication mode" >> ${LOG_FILE}
    node /work/c9sdk/server.js -p 8282 -w ${C9_PATH} >> /dev/null &
else
    # with auth
    echo "Starting c9 in authentication mode" >> ${LOG_FILE}
    node /work/c9sdk/server.js -p 8282 -w ${C9_PATH} --auth ${AUTH} >> /dev/null &
fi

# start main server
cd /server/src && /server/src/main
