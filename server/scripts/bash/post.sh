#!/bin/bash

##  newpost - create new post: e.g. $ hops newpost <name>
_hops_newpost() {
    if [ "$#" -ne 1 ]; then
	    echo "Expecting args: <post-name>"
	    return 1
    fi
    
    # goto hugo src dir and create post
    cd ${BLOG_PATH}
    hugo new post/$1 || (cd - && return 1)
    cd -
    return 0
}