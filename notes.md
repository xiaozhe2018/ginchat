## 踩坑记录
### 1. 安装swagger之后执行swag init 报zsh: command not found: swag 查看bin目录下面有swag vi ~/.zshrc 加入 export PNPM_HOME="/Users/xiaozhe/go/bin:$PATH"  保存退出 ，source ~/.zshrc 再次执行swag init即可