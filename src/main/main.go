package main

import (
	"archive/zip"
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

/**
mvn
	-s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml
	-Dmaven.repo.local=/Users/mac/Desktop/resp
	-DskipTests=true deploy:deploy-file
	-DgroupId=cn.sunline.acm
	-DartifactId=acm-web
	-Dversion=5.5.0-SNAPSHOT
	-Dpackaging=war
	-Dfile=acm-web-5.5.0-SNAPSHOT.war
	-Durl=http://172.19.9.94:10000/repository/maven-snapshots/
	-DrepositoryId=nexus-snapshots
*/
type MvnDeploy struct {
	repoLocal, artifactId, version, groupId, packaging, file, url, repositoryId string
}

type PomDoc struct {
	XMLName    xml.Name  `xml:"project"`
	GroupId    string    `xml:"groupId"`
	ArtifactId string    `xml:"artifactId"`
	Version    string    `xml:"version"`
	Parent     ParentDom `xml:"parent"`
}

type ParentDom struct {
	XMLName xml.Name `xml:"parent"`
	GroupId string   `xml:"groupId"`
	Version string   `xml:"version"`
}

// 创建一个数组
var war_file_arr, jar_file_arr, pom_file_arr = make([]string, 0), make([]string, 0), make([]string, 0)

var url = "http://172.19.9.94:10000/repository/maven-snapshots/"

//var   urlR = "http://172.19.9.94:10000/repository/maven-releases/"
var repoLocal = "/Users/mac/Desktop/resp"
var repositoryId = "nexus-snapshots"

//var  repositoryIdR = "nexus-releases"
var settingConfigPath = "/Applications/apache-maven-3.3.9/conf/xxx.xml"
var targetDir = ""
var pomVersion = ""

var conf = "conf.properties"
var second_cmd = ""

func init() {
	initProperties()
	// 解析配置文件
	parsing_conf()
}

/**
解析配置文件
*/
func parsing_conf() {
	f, err := os.Open(conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// 创建map
	prop := make(map[string]string)

	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(a)

		line_arr := strings.Split(line, "=")
		prop[strings.Trim(line_arr[0], " ")] = strings.Trim(line_arr[1], " ")
	}

	// 给参数赋值
	assignment(prop)

}

/**
为全局变量赋值
*/
func assignment(prop map[string]string) {
	log.Println(prop)
	url = prop["url"]
	repositoryId = prop["repositoryId"]
	if pomVersion == "releases" {
		url = prop["urlR"]
		repositoryId = prop["repositoryIdR"]
	}

	repoLocal = prop["repoLocal"]
	settingConfigPath = prop["settingConfigPath"]
	targetDir = prop["targetDir"]
}

/**
初始化配置文件
*/
func initProperties() {
	if len(os.Args) <= 1 {
		log.Fatal("请输入子命令: pom|jar|war|version|help")
	}

	flag.StringVar(&conf, "c", "127.0.0.1", "请指定配置文件的地址")
	flag.StringVar(&pomVersion, "v", "snapshots", "请输入包的版本")

	flag.CommandLine.Parse(os.Args[2:])

	second_cmd = os.Args[1]

	log.Println("配置文件地址>>" + conf)
}

func main() {
	//要遍历的文件夹

	//遍历的文件夹
	//参数：要遍历的文件夹，层级（默认：0）
	findDir(targetDir, 0)
	if second_cmd == "jar" {
		jarAction()
	} else if second_cmd == "war" {
		warAction()
	} else if second_cmd == "pom" {
		pomAction()
	} else {
		log.Fatalln("不能识别的文件类型")
	}
	log.Println("程序执行完毕自动退出...")
}

func jarAction() {
	println("==========jar_file===========")
	for index, value := range jar_file_arr {
		println(index, "  ", value)
		cmd := zipList(value, "jar")
		fmt.Println(cmd)
		CmdExecutor(cmd)
	}
}

func warAction() {
	println("==========war_file===========")
	for index, value := range war_file_arr {
		println(index, "   ", value)
		cmd := zipList(value, "war")
		fmt.Println(cmd)
		CmdExecutor(cmd)

	}
}

func pomAction() {
	fmt.Println("==========pom_file===========")
	for index, value := range pom_file_arr {
		fmt.Println(index, "  ", value)
		cmd := pasePomFile(value)
		fmt.Println(cmd)
		CmdExecutor(cmd)
	}
}

/**
****************************** 遍历目录扫描所有的pom jar war *************************
 */
func findDir(dir string, num int) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	// 遍历这个文件夹
	for _, fi := range fileinfo {

		// 重复输出制表符，模拟层级结构
		fmt.Println(strings.Repeat("\t", num))

		// 判断是不是目录
		if fi.IsDir() {
			//println(`目录：`, fi.Name())
			findDir(dir+`/`+fi.Name(), num+1)
		} else {
			//println(`文件：`,dir+ fi.Name())
			if strings.HasSuffix(fi.Name(), ".war") {
				war_file_arr = append(war_file_arr, dir+string(os.PathSeparator)+fi.Name())
			} else if strings.HasSuffix(fi.Name(), ".jar") {
				if strings.HasSuffix(fi.Name(), "-sources.jar") {
					log.Println(fi.Name() + " >>>是源码包自动排除掉")
					continue
				}
				jar_file_arr = append(jar_file_arr, dir+string(os.PathSeparator)+fi.Name())
			} else if strings.HasSuffix(fi.Name(), ".pom") {
				pom_file_arr = append(pom_file_arr, dir+string(os.PathSeparator)+fi.Name())
			}
		}
	}

	return
}

/**
	解析pom file

mvn -s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml \
-Dmaven.repo.local=/Users/mac/Desktop/resp  \
-DskipTests=true deploy:deploy-file \
-DgroupId=cn.caijiajia \
-DartifactId=flowplus-parent \
-Dversion=1.0.0.RELEASE \
-Dfile=flowplus-parent-1.0.0.RELEASE.pom \
-Dpackaging=pom  \
-Durl=http://172.19.9.94:10000/repository/maven-releases/ \
-DpomFile=flowplus-parent-1.0.0.RELEASE.pom \
-DrepositoryId=nexus-releases

*/
func pasePomFile(filePath string) (cmd *exec.Cmd) {
	//解析xml拿到 artifactId version groupId
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := PomDoc{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	version := v.Version
	groupId := v.GroupId
	artifactId := v.ArtifactId

	if len(version) == 0 {
		version = v.Parent.Version
	}
	if len(groupId) == 0 {
		groupId = v.Parent.GroupId
	}

	//fmt.Println("groupId>>>"+groupId)
	//fmt.Println("artifactId>>>"+artifactId)
	//fmt.Println("version>>>"+version)

	cmd = exec.Command("mvn",
		"-s", settingConfigPath,
		"-Dmaven.repo.local="+repoLocal,
		"-DskipTests=true",
		"deploy:deploy-file",
		"-DgroupId="+groupId,
		"-DartifactId="+artifactId,
		"-Dversion="+version,
		"-Dpackaging=pom",
		"-Dfile="+filePath,
		"-Durl="+url,
		"-DpomFile="+filePath[0:len(filePath)-3]+"pom",
		"-DrepositoryId="+repositoryId)
	return
}

/**
获取zip 文件列表中pom.properties
*/
func zipList(filePath string, packaging string) (cmd *exec.Cmd) {
	m := make(map[string]string)

	cf, err := zip.OpenReader(filePath) //读取zip文件
	if err != nil {
		fmt.Println(err)
	}
	defer cf.Close()
	for _, file := range cf.File {
		//println(file.Name)
		if strings.HasSuffix(file.Name, "pom.properties") {
			rc, err := file.Open()
			if err != nil {
				println(err)
			}
			br := bufio.NewReader(rc)
			for {
				line, _, c := br.ReadLine()
				if c == io.EOF {
					break
				}
				lineStr := string(line) //一行一行文件
				if strings.Contains(lineStr, "=") && !strings.Contains(lineStr, "#") {
					arr := strings.Split(lineStr, "=")
					m[arr[0]] = arr[1]
				}

			}
		}
	}
	//封装成commond 对象
	cmd = exec.Command("mvn",
		"-s", settingConfigPath,
		"-Dmaven.repo.local="+repoLocal,
		"-DskipTests=true",
		"deploy:deploy-file",
		"-DgroupId="+m["groupId"],
		"-DartifactId="+m["artifactId"],
		"-Dversion="+m["version"],
		"-Dpackaging="+packaging,
		"-Dfile="+filePath,
		"-Durl="+url,
		"-DpomFile="+filePath[0:len(filePath)-3]+"pom",
		"-DrepositoryId="+repositoryId)
	return
}

/**
拼装命令
	mvn
		-s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml
		-Dmaven.repo.local=/Users/mac/Desktop/resp
		-DskipTests=true deploy:deploy-file
		-DgroupId=cn.sunline.acm
		-DartifactId=acm-web
		-Dversion=5.5.0-SNAPSHOT
		-Dpackaging=war
		-Dfile=acm-web-5.5.0-SNAPSHOT.war
		-Durl=http://172.19.9.94:10000/repository/maven-snapshots/
		-DrepositoryId=nexus-snapshots
*/

// 命令执行者
func CmdExecutor(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	//获取输出对象，可以从该对象中读取输出结果
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}
