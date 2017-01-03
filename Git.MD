## Git 本地提交到远端

出现这样的提示，主要原因是本地可能是通过Git Init生成出来的和GitHub上的Git分支没有关联。
fatal: refusing to merge unrelated histories

解决方式：假如我们的源是origin，分支是master，git pull origin master --allow-unrelated-histories
需要知道，我们的源可以是本地的路径。

git remote add origin git@github.com:OperatorWorld/ShareCode.git

git push -u origin master

git branch --set-upstream-to=<remote>/<branch> master

## 第一次提交是问题

The authenticity of host 'github.com (192.30.253.113)' can't be established.
RSA key fingerprint is SHA256:nThbg6kXUpJWGl7E1IGOCspRomTxdCARLviKw6E5SY8.
Are you sure you want to continue connecting (yes/no)? Y
Please type 'yes' or 'no': yes
Warning: Permanently added 'github.com,192.30.253.113' (RSA) to the list of known hosts.
Permission denied (publickey).
fatal: Could not read from remote repository.

出现这样错误原因，SSH的Key没有的仓库内建立。

git config --global user.name "yourname"
git config --global user.email "youremail"

ssh-keygen -t rsa -C ""
Your identification has been saved /.ssh/id_rsa.
Your public key has been saved in /.ssh/id_rsa.pub