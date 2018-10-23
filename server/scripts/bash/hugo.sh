#!/bin/bash

hops() {
    option=$1
    prefix="_"
     if [[ $option == "" ]]; then
        ${prefix}hops_helpinfo
     else
        ${prefix}hops_${option} "${@:2}" || _hops_helpinfo
     fi
}

##  configure - configure/reconfigure hops with stack details;
_hops_configure() {
    echo "configure hops settings.."
}

_hops_helpinfo() {
    echo
    echo "hops [option] "
    echo
    for i in `find ${SERVER_PATH}/scripts -name '*.sh'`;
    do
      awk '/^##/ {
		print substr( $0, 3, 500 )
		while (getline line) {
			   if ( line !~ /^##/ ) {
        			break
    			}
    		   print substr( line, 3, 500 )
        }
      }' $i
    done
    echo
}

_getwords() {
    for i in `find ${SERVER_PATH}/scripts -name '*.sh'`;
    do
        words=$(grep -h "^_hops_"  $i | grep "()"  | awk -F[_\(] '{print $3}');
        for word in $words;
            do
                echo $word
            done
    done
}

_hopsAutoComplete() {
     _allwords=$(_getwords)
    local cur=${COMP_WORDS[COMP_CWORD]}
    COMPREPLY=( $(compgen -W "${_allwords}" -- $cur) )
}

complete -F _hopsAutoComplete hops
