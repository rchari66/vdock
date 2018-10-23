# This will create couple of custom bash functions for hugo and git
sourceUpdate() {
    for i in `find ${SERVER_PATH}/scripts -name '*.sh'`; do source $i ; done
}

# source all *.sh files in $SERVER_PATH/scripts
sourceUpdate
