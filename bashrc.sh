# This will create couple of custom bash functions for hugo and git
sourceUpdate() {
    for i in `find ${SERVER_PATH}/scripts -name '*.sh'`; do source $i ; done
}

# source all *.sh files in $SERVER_PATH/scripts
sourceUpdate

# Alias for git commands
setupGitAliases() {
	alias gd='git diff'
	alias gb='git branch --sort=creatordate'
	alias gcm="git checkout master"
	alias gp='git pull'
	alias gc='git checkout $1'
	alias gdc='git diff --color-words'
	alias gds='git diff --staged --color-words'  
	alias gl='git log --oneline -n 20'
	alias gs='git status'
}