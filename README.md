**nexus-cli  是针对nexus 私服上传下载依赖的一个工具**

# 快速开始
目前需要自己手动改一些变量， 然后配置go语言运行环境，运行，修改的变量如下
```go

// 私服snapshots group
const url = "http://172.19.9.94:10000/repository/maven-snapshots/"
// 私服releases group
const urlR = "http://172.19.9.94:10000/repository/maven-releases/"
// 本地仓库路径
const repoLocal = "/Users/mac/Desktop/resp"
// 远程仓库snapshots group 的id
const repositoryId = "nexus-snapshots"
// 远程仓库releases group 的id
const repositoryIdR = "nexus-releases"
//apache-maven conf 的配置文件地址
const settingConfigPath = "/Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml"
// 需要上传jar war  pom 文件的目录
dir := `/Users/mac/Documents/项目/xxx/xxx/temp/pom/`
```
# 例子
以传jar包为例
1. 把jar包和pom放在一个目录下
![avatar](https://images.gitee.com/uploads/images/2020/0511/175012_e4b76aee_1894834.jpeg)
2.setting.xml 需要配置私服的用户名密码
3. 配置配置文件
4. 执行命令
```shell
./nexus-cli-mac  jar -c /Users/mac/Documents/gowork/src/nexus-cli/src/resources/conf.properties -v snapshots
```
5. 效果
![avatar](https://images.gitee.com/uploads/images/2020/0511/175158_5c1c04b9_1894834.jpeg)
# TODO 
-   把参数抽取成命令行参数，用户无需修改代码
-   打包成linux MacOs window 三种版本二进制可执行文件
-   支持更多关于nuxus 问题