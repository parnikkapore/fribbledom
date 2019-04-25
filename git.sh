# Convenience functions for Git

function git_save(){
    git add -A && git commit -a
}

function git_save_upload(){
    git add -A && git commit -a && git push -u origin master
}
