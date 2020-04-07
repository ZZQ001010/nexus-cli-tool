mvn -s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml \
-Dmaven.repo.local=/Users/mac/Desktop/resp  \
-DskipTests=true deploy:deploy-file \
-DgroupId=cn.caijiajia \
-DartifactId=acm-web \
-Dversion=0.0.1-SNAPSHOT \
-Dpackaging=jar  \
-Dfile=aliyun-service-core-0.0.1-SNAPSHOT.jar \
-Durl=http://172.19.9.94:10000/repository/maven-snapshots/ \
-DpomFile=aliyun-service-core-0.0.1-SNAPSHOT.pom \
-DrepositoryId=nexus-snapshots \




mvn -s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml 
-Dmaven.repo.local=/Users/mac/Desktop/resp 
-DskipTests=true deploy:deploy-file  
-DgroupId=cn.sunline.acm 
-DartifactId=acm-surface-impl 
-Dversion=6.0.0-SNAPSHOT 
-Dpackaging=jar 
-Dfile=/Applications/apache-maven-3.3.9/respository/cn/sunline/acm/acm-surface-impl/6.0.0-SNAPSHOT/acm-surface-impl-6.0.0-SNAPSHOT.jar 
-Durl=http://172.19.9.94:10000/repository/maven-snapshots/ 
-DrepositoryId=n	exus-snapshots 


/Applications/apache-maven-3.3.9/bin/mvn -s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml 
-Dmaven.repo.local=/Users/mac/Desktop/resp 
-DskipTests=true deploy:deploy-file 
-DgroupId=cn.sunline.acm 
-DartifactId=acm-batch 
-Dversion=5.5.0-SNAPSHOT 
-Dpackaging=jar 
-Dfile=/Applications/apache-maven-3.3.9/respository/cn/sunline/acm/acm-batch/5.5.0-SNAPSHOT/acm-batch-5.5.0-SNAPSHOT.jar 
-Durl=http://172.19.9.94:10000/repository/maven-snapshots/ 
-DrepositoryId=nexus-snapshots


/Applications/apache-maven-3.3.9/bin/mvn -s /Applications/apache-maven-3.3.9/conf/settings_sh_sunline.xml -Dmaven.repo.local=/Users/mac/Desktop/resp -DskipTests=true deploy:deploy-file -DgroupId=cn.caijiajia -DartifactId=aliyun-service-core -Dversion=0.0.1-SNAPSHOT -Dpackaging=jar -Dfile=/Users/mac/Documents/项目/数禾零售贷款核心/shuhejar/jars-target//aliyun-service-core-0.0.1-SNAPSHOT.jar -Durl=http://172.19.9.94:10000/repository/maven-snapshots/ -DrepositoryId=nexus-snapshots





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