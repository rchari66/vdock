#!/bin/bash

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
