#! /bin/bash

##  getLogin - Login to git;
_hops_gitLogin() {
    # git config credential.helper store
    if [ "$#" -ne 2 ]; then
	    echo "Expecting args: <git-user-id> <password>"
	    return
    fi
    expect ../expect/githubLogin.exp $1 $2
}

##  getGitRepos - get repository list for given github user;
_hops_getGitRepos() {
    if [ "$#" -ne 1 ]; then
	    echo "Expecting args: <githubUserId>"
	    return ""
    fi
    repos=$(curl https://api.github.com/users/rchari66/repos | grep html_url | grep -v "rchari66\",$" | cut -d":" -f2,3 )
    return $repos
}
