package models

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Post post model
type Post struct {
	gorm.Model
	Authors         []*User   `gorm:"many2many:user_posts"`       // 所有作者, 包括第一作者
	Tags            []*Tag    `gorm:"many2many:post_tags"`        // 标签
	PrimaryAuthor   *User     `gorm:"foreignkey:PrimaryAuthorID"` // 第一作者
	UUID            uuid.UUID `gorm:"type:uuid;unique_index"`     // UUID
	Slug            string    `gorm:"unique_index;not null"`      // Slug 对应post的实际url地址
	PrimaryAuthorID uint
	Title           string
	Excerpt         string
	Markdown        string
	HTML            string
	Image           string
	Featured        bool
	Paged           bool
	Status          string
	Language        string
	MetaTitle       string
	MetaDescription string
	PublishedAt     time.Time
	PublishedBy     uint
	CreateBy        uint
	UpdateBy        uint
}

// BeforeCreate 初始化uuid
func (v *Post) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// IsNode IsNode
func (v *Post) IsNode() {}

// GetID GetID
func (v Post) GetID() interface{} {
	return v.ID
}

// GetCreateAt GetCreateAt
func (v Post) GetCreateAt() interface{} {
	return v.CreatedAt
}

// GetUpdateAt GetUpdateAt
func (v Post) GetUpdateAt() interface{} {
	return v.UpdatedAt
}

func newDemoPost(i int) (p *Post) {
	p = &Post{
		Authors:         PreDefinedUsers,
		Tags:            PreDefinedTags,
		PrimaryAuthor:   PreDefinedUsers[0],
		PrimaryAuthorID: PreDefinedUsers[0].ID,
		Title:           "demo-blog-title" + strconv.Itoa(i),
		Slug:            "demo-blog-slug" + strconv.Itoa(i),
		Excerpt:         "demo-blog-excerpt",
		Featured:        false,
		Paged:           false,
		Status:          "published",
		HTML:            "HTML CONTENT",
	}
	return
}

func init() {
	for i := 0; i < 20; i++ {
		PreDefinedPosts = append(PreDefinedPosts, newDemoPost(i))
	}
}

// PreDefinedPosts PreDefinedPosts
var (
	PreDefinedPosts = []*Post{
		{
			Authors:         PreDefinedUsers,
			Tags:            PreDefinedTags,
			PrimaryAuthor:   PreDefinedUsers[0],
			PrimaryAuthorID: PreDefinedUsers[0].ID,
			Title:           "ubuntu18.04环境下docker设置指南",
			Slug:            "ubuntu-docker-setup",
			Excerpt:         "在ubuntu18.04下安装docker和docker-compose",
			Featured:        false,
			Paged:           false,
			Status:          "published",
			HTML:            "<!--kg-card-begin: markdown--><h2 id=\"docker\">安装docker</h2>\n<p>依照官方文档的介绍，目前有两种主流的方式安装docker,其中脚本安装方式应该是对手动命令安装的一种封装，总之两种安装方式结果都一致</p>\n<h3 id=\"dockerinstall\">使用<code>docker-install</code>脚本进行安装</h3>\n<blockquote>\n<p>详情参阅<a href=\"https://github.com/docker/docker-install\">https://github.com/docker/docker-install</a></p>\n</blockquote>\n<p>这个安装脚本的目的是为了在受支持的linux发行版（主流发行版如ubuntu,centos等）中快速安装最新docker-ce。建议不要在生产系统中依赖此脚本进行部署工作。<br>\n<em>经笔者测试，此脚本在<code>AWS</code>的<code>Amazon Linux</code>中并不能正常工作。</em></p>\n<p>命令如下：</p>\n<pre><code class=\"language-bash\">curl -fsSL https://get.docker.com -o get-docker.sh\nsh get-docker.sh\n</code></pre>\n<h3 id=\"\">依照官方文档进行安装</h3>\n<blockquote>\n<p>详情参阅<a href=\"https://docs.docker.com/install/\">https://docs.docker.com/install/</a></p>\n</blockquote>\n<p>命令如下：</p>\n<pre><code class=\"language-bash\"># 卸载旧版本的docker, 如果是新系统，可跳过此步骤\nsudo apt-get remove docker docker-engine docker.io containerd runc\n\n# 更新软件源\nsudo apt-get update\n\n# 允许apt使用https协议\nsudo apt-get install \\\n    apt-transport-https \\\n    ca-certificates \\\n    curl \\\n    gnupg-agent \\\n    software-properties-common\n\n# 添加Docker的官方GPG秘钥\n# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -\n\n# 验证秘钥\nsudo apt-key fingerprint 0EBFCD88\n\n# 设置软件仓库， 此命令设置为软件仓库为稳定版 stable\nsudo add-apt-repository \\\n&quot;deb [arch=amd64] https://download.docker.com/linux/ubuntu \\\n$(lsb_release -cs) \\\nstable&quot;\n\n# 更新软件源\nsudo apt-get update\n\n# 正式安装docker\nsudo apt-get install docker-ce docker-ce-cli containerd.io\n</code></pre>\n<p>新建<code>docker</code>用户组，并将当前用户添加到<code>docker</code>组中，这样可以避免当前用户使用docker命令时必须加sudo才能</p>\n<pre><code class=\"language-bash\">sudo groupadd docker\nsudo usermod -aG docker $USER\n# 随后重启或者注销，以使用户组设置生效\n</code></pre>\n<p>设置docker开机自启</p>\n<pre><code class=\"language-bash\">sudo systemctl enable docker\n</code></pre>\n<p>disable docker start on boot</p>\n<pre><code class=\"language-bash\">sudo systemctl disable docker\n</code></pre>\n<h2 id=\"docker\">设置docker国内镜像</h2>\n<p>从以下镜像中选择一个镜像，笔者使用阿里云镜像</p>\n<ul>\n<li><a href=\"https://www.docker-cn.com/registry-mirror\">Docker 中国官方镜像加速</a></li>\n<li><a href=\"https://cr.console.aliyun.com/cn-hangzhou/mirrors\">阿里云加速(需注册账号)</a></li>\n<li><a href=\"https://www.daocloud.io/mirror\">DaoCloud 加速器</a></li>\n<li><a href=\"https://kirk-enterprise.github.io/hub-docs/#/user-guide/mirror\">七牛云镜像加速</a></li>\n</ul>\n<p>将下列命令的<code>https://replace.with.your.mirror.com</code>替换成你选择的镜像地址，按顺序执行下列命令</p>\n<pre><code class=\"language-bash\">sudo mkdir -p /etc/docker\nsudo tee /etc/docker/daemon.json &lt;&lt;-'EOF'\n{\n  &quot;registry-mirrors&quot;: [&quot;https://replace.with.your.mirror.com&quot;]\n}\nEOF\n\nsudo systemctl daemon-reload\nsudo systemctl restart docker\n</code></pre>\n<h2 id=\"dockercompose\">安装docker-compose</h2>\n<p>依照官方文档的介绍，目前主流的方式是直接下载对应平台的compose二进制文件<br>\n命令如下：</p>\n<pre><code class=\"language-bash\">curl -L https://github.com/docker/compose/releases/download/1.24.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose\nchmod +x /usr/local/bin/docker-compose\n</code></pre>\n<h2 id=\"dockerdockercompose\">测试docker与docker-compose</h2>\n<p>运行下列命令：验证docker与docker-compose 是否安装成功</p>\n<pre><code class=\"language-bash\">docker --version\ndocker-compose --version\n</code></pre>\n<p>如正确返回版本号，docker安装到此结束。</p>\n<!--kg-card-end: markdown-->",
			Markdown:        "\\n[^_^]: # (url:ubuntu-docker-setup)\\n[^_^]: # (tag:#tech,note,setup,docker)\\n[^_^]: # (excerpt:在ubuntu18.04下安装docker和docker-compose)\\n\\n## 安装docker\\n\\n依照官方文档的介绍，目前有两种主流的方式安装docker,其中脚本安装方式应该是对手动命令安装的一种封装，总之两种安装方式结果都一致\\n\\n### 使用`docker-install`脚本进行安装\\n\\n> 详情参阅[https://github.com/docker/docker-install](https://github.com/docker/docker-install)\\n\\n这个安装脚本的目的是为了在受支持的linux发行版（主流发行版如ubuntu,centos等）中快速安装最新docker-ce。建议不要在生产系统中依赖此脚本进行部署工作。  \\n*经笔者测试，此脚本在`AWS`的`Amazon Linux`中并不能正常工作。*\\n\\n命令如下：\\n\\n```bash\\ncurl -fsSL https://get.docker.com -o get-docker.sh\\nsh get-docker.sh\\n```\\n\\n### 依照官方文档进行安装\\n\\n> 详情参阅[https://docs.docker.com/install/](https://docs.docker.com/install/)\\n\\n命令如下：\\n\\n```bash\\n# 卸载旧版本的docker, 如果是新系统，可跳过此步骤\\nsudo apt-get remove docker docker-engine docker.io containerd runc\\n\\n# 更新软件源\\nsudo apt-get update\\n\\n# 允许apt使用https协议\\nsudo apt-get install \\\\\\n    apt-transport-https \\\\\\n    ca-certificates \\\\\\n    curl \\\\\\n    gnupg-agent \\\\\\n    software-properties-common\\n\\n# 添加Docker的官方GPG秘钥\\n# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -\\n\\n# 验证秘钥\\nsudo apt-key fingerprint 0EBFCD88\\n\\n# 设置软件仓库， 此命令设置为软件仓库为稳定版 stable\\nsudo add-apt-repository \\\\\\n\\\"deb [arch=amd64] https://download.docker.com/linux/ubuntu \\\\\\n$(lsb_release -cs) \\\\\\nstable\\\"\\n\\n# 更新软件源\\nsudo apt-get update\\n\\n# 正式安装docker\\nsudo apt-get install docker-ce docker-ce-cli containerd.io\\n```\\n\\n新建`docker`用户组，并将当前用户添加到`docker`组中，这样可以避免当前用户使用docker命令时必须加sudo才能\\n\\n```bash\\nsudo groupadd docker\\nsudo usermod -aG docker $USER\\n# 随后重启或者注销，以使用户组设置生效\\n```\\n\\n设置docker开机自启\\n\\n```bash\\nsudo systemctl enable docker\\n```\\n\\ndisable docker start on boot\\n\\n```bash\\nsudo systemctl disable docker\\n```\\n\\n## 设置docker国内镜像\\n\\n从以下镜像中选择一个镜像，笔者使用阿里云镜像\\n\\n- [Docker 中国官方镜像加速](https://www.docker-cn.com/registry-mirror)\\n- [阿里云加速(需注册账号)](https://cr.console.aliyun.com/cn-hangzhou/mirrors)\\n- [DaoCloud 加速器](https://www.daocloud.io/mirror)\\n- [七牛云镜像加速](https://kirk-enterprise.github.io/hub-docs/#/user-guide/mirror)\\n\\n将下列命令的`https://replace.with.your.mirror.com`替换成你选择的镜像地址，按顺序执行下列命令\\n\\n```bash\\nsudo mkdir -p /etc/docker\\nsudo tee /etc/docker/daemon.json <<-'EOF'\\n{\\n  \\\"registry-mirrors\\\": [\\\"https://replace.with.your.mirror.com\\\"]\\n}\\nEOF\\n\\nsudo systemctl daemon-reload\\nsudo systemctl restart docker\\n```\\n\\n## 安装docker-compose\\n\\n依照官方文档的介绍，目前主流的方式是直接下载对应平台的compose二进制文件\\n命令如下：\\n\\n```bash\\ncurl -L https://github.com/docker/compose/releases/download/1.24.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose\\nchmod +x /usr/local/bin/docker-compose\\n```\\n\\n## 测试docker与docker-compose\\n\\n运行下列命令：验证docker与docker-compose 是否安装成功\\n\\n```bash\\ndocker --version\\ndocker-compose --version\\n```\\n\\n如正确返回版本号，docker安装到此结束。\\n",
		},
		{
			Authors:         PreDefinedUsers,
			Tags:            PreDefinedTags,
			PrimaryAuthor:   PreDefinedUsers[0],
			PrimaryAuthorID: PreDefinedUsers[0].ID,
			Title:           "Windows 环境 JDK 配置",
			Slug:            "jdk-windows-setup",
			Excerpt:         "Windows 下配置多版本 JDK 环境",
			Featured:        false,
			Paged:           false,
			Status:          "published",
			HTML:            "<!--kg-card-begin: markdown--><blockquote>\n<p>Tested Environment<br>\nOperating System: Microsoft Windows Version 1903 (OS Build 18865.1000)<br>\n谨以此文献给笔者活在虚拟机的 Windows</p>\n</blockquote>\n<h2 id=\"\">概览</h2>\n<p><code>JDK</code>的设置主要有两点：</p>\n<ul>\n<li>下载正确的版本\n<ul>\n<li>对应平台（Windows/Linux）</li>\n<li>对应架构（64bit/32bit）</li>\n<li>对应需求版本（1.8/11）</li>\n</ul>\n</li>\n<li>配置正确的环境变量\n<ul>\n<li>JAVA_HOME</li>\n<li>Path</li>\n</ul>\n</li>\n</ul>\n<h2 id=\"jdk\">下载 JDK</h2>\n<p>建议同时下载<code>JDK8</code>和<code>JDK11</code>，无特殊需求建议使用<code>OpenJDK</code>，建议下载压缩包格式的（非<code>exe</code>或<code>msi</code>后缀）<br>\n为方便读者查阅，以下几个网址是笔者推荐的<code>JDK</code>获取途径：</p>\n<ul>\n<li><a href=\"https://www.oracle.com/technetwork/es/java/javasebusiness/downloads/index.html\">Oracle JDK</a></li>\n<li><a href=\"https://openjdk.java.net/\">OpenJDK</a></li>\n<li><a href=\"https://github.com/ojdkbuild/ojdkbuild\">ojdkbuild</a></li>\n</ul>\n<p>笔者使用的是 ojdkbuild 提供的 JDK，直接到<a href=\"https://github.com/ojdkbuild/ojdkbuild\">ojdkbuild</a>下载对应平台和版本的<code>JDK</code>，不再赘述。</p>\n<p><img src=\"/content/images/2019/04/ojdkbuild-download.png\" alt=\"ojdkbuild-download\"></p>\n<h2 id=\"jdk\">解压 JDK</h2>\n<ol>\n<li>\n<p>在<code>c盘</code>新建<code>Dev</code>文件夹，将下载的<code>JDK</code>解压到该文件夹下</p>\n</li>\n<li>\n<p>随后删除空层文件夹并更改<code>JDK</code>文件夹的名字，使其简短些</p>\n</li>\n<li>\n<p>注意文件夹层次结构（注意第二步的提示），理想情况下应该有如图所示有效路径：</p>\n<p><img src=\"/content/images/2019/04/jdk-folder.png\" alt=\"jdk-folder\"></p>\n</li>\n</ol>\n<h2 id=\"\">设置环境变量</h2>\n<p>为了后续的脚本切换<code>JDK</code>,此处使用的是用户环境变量，<em>非</em>系统环境变量，如非特殊说明，环境变量一律指代用户环境变量</p>\n<ol>\n<li>\n<p>新建 <code>JAVA_HOME_{version}</code> 变量,其中 version 为 JDK 版本,有多少个版本的 JDK 就建立多少个对应的环境变量,环境变量的值分别为对应的<code>JDK</code>的路径<br>\n例如: 本人使用两个版本的<code>JDK</code>,分别是 JDK8 和 11,则新建两个环境变量,分别是</p>\n<table>\n<thead>\n<tr>\n<th>Variable</th>\n<th>Value</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>JAVA_HOME_8</td>\n<td>C:\\Dev\\Java\\OpenJDK1.8.0</td>\n</tr>\n<tr>\n<td>JAVA_HOME_11</td>\n<td>C:\\Dev\\Java\\OpenJDK11.0.2</td>\n</tr>\n</tbody>\n</table>\n<p><img src=\"/content/images/2019/04/env-java-home-8.png\" alt=\"env-java-home-8\"><br>\n<img src=\"/content/images/2019/04/env-java-home-11.png\" alt=\"env-java-home-11\"></p>\n</li>\n<li>\n<p>新建 <code>JAVA_HOME</code> 变量,其值为上一步所新建的变量,格式为<code>%</code>包裹,见示例:</p>\n<table>\n<thead>\n<tr>\n<th>Variable</th>\n<th>Value</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>JAVA_HOME</td>\n<td>%JAVA_HOME_8%</td>\n</tr>\n</tbody>\n</table>\n<p><img src=\"/content/images/2019/04/env-java-home.png\" alt=\"env-java-home\"></p>\n</li>\n<li>\n<p><code>Path</code>变量中新增一条记录,设置为<code>%JAVA_HOME%\\bin</code></p>\n<p><img src=\"/content/images/2019/04/env-path.png\" alt=\"env-path\"></p>\n</li>\n</ol>\n<h2 id=\"\">检查配置是否成功</h2>\n<p>命令行运行<code>java -version</code>，应该输出<code>Java</code>的版本信息</p>\n<p><img src=\"/content/images/2019/04/cmd-java-version.png\" alt=\"cmd-java-version\"></p>\n<!--kg-card-end: markdown--><!--kg-card-begin: markdown--><h2 id=\"\">编写版本切换脚本</h2>\n<p>如需更改当前使用的<code>Java</code>版本，可以手动更改<code>JAVA_HOME</code>的变量值。如果更改<code>Java</code>版本的频次高，可以使用脚本来简化这一操作。<br>\n笔者技拙，尝试过使用<code>bat脚本</code>更改系统环境变量，但这样脚本就需要管理员权限，且无法保证<code>JDK</code>的切换效果，遂更换思路，设置用户环境变量来达成目的。<br>\n下列脚本一律不需要管理员权限，且只更改用户环境变量，在<code>Windows 10,v1903</code>测试通过。如果你的<code>Java</code>环境变量跟上文提到的设置方式一致，那么应该可以直接使用下面脚本。<br>\n新建一个文件夹，如<code>C:\\Scripts</code>，然后在该文件夹下新建这三个文件：</p>\n<ol>\n<li>主脚本 <code>jdk_switch.bat</code></li>\n<li>切换脚本 <code>switch_to_openjdk_8.bat</code></li>\n<li>切换脚本 <code>switch_to_openjdk_11.bat</code></li>\n</ol>\n<p><img src=\"/content/images/2019/04/scripts-folder.png\" alt=\"scripts-folder\"></p>\n<h3 id=\"jdk_switchbat\">主脚本 jdk_switch.bat</h3>\n<p>(代码托管于Github Gist,只有聪明人才看得见#逃)</p>\n<script src=\"https://gist.github.com/nediiii/03f7daf70f6989c671e953eeeabab9fd.js\"></script><!--kg-card-end: markdown--><!--kg-card-begin: markdown--><h3 id=\"switch_to_openjdk_8bat\">切换脚本 switch_to_openjdk_8.bat</h3>\n<pre><code class=\"language-batch\">SETX JAVA_HOME &quot;%%JAVA_HOME_8%%&quot;\n</code></pre>\n<h3 id=\"switch_to_openjdk_11bat\">切换脚本 switch_to_openjdk_11.bat</h3>\n<pre><code class=\"language-batch\">SETX JAVA_HOME &quot;%%JAVA_HOME_11%%&quot;\n</code></pre>\n<h3 id=\"\">脚本效果</h3>\n<p>运行脚本的效果如图所示：</p>\n<p><img src=\"/content/images/2019/04/script-view.png\" alt=\"script-view\"></p>\n<h3 id=\"\">添加环境变量</h3>\n<p>为了在命令行中能够直接使用脚本，需要把脚本目录添加进 Path 变量里。过程不再赘述。如图：</p>\n<p><img src=\"/content/images/2019/04/env-path.png\" alt=\"env-path\"></p>\n<p>此后可在命令行中直接运行脚本：</p>\n<p><img src=\"/content/images/2019/04/cmd-script.png\" alt=\"cmd-script\"></p>\n<h3 id=\"\">设置索引</h3>\n<p>为了使用 Windows 搜索时能直接搜索到脚本，需要把脚本目录添加进索引目录里。<br>\n要进入<code>索引设置</code>，有两个方式（下列第 1 步和第 2 步最终效果相同，皆是为了进入第 3 步）</p>\n<ol>\n<li>\n<p>按<code>Windows键</code>然后键入<code>change how windows</code>，如图所示，点击图中标示的选项（如有问题，可从第 2 步开始）：</p>\n<p><img src=\"/content/images/2019/04/windows-search.png\" alt=\"windows-search\"></p>\n</li>\n<li>\n<p>从<code>设置</code>-&gt;<code>搜索</code>里进入，如图：</p>\n<p><img src=\"/content/images/2019/04/setting-search-1.png\" alt=\"setting-search-1\"><br>\n<img src=\"/content/images/2019/04/setting-search-2.png\" alt=\"setting-search-2\"></p>\n</li>\n<li>\n<p>进入到了索引设置，点击修改，如图：</p>\n<p><img src=\"/content/images/2019/04/setting-search-3.png\" alt=\"setting-search-3\"></p>\n</li>\n<li>\n<p>勾选上脚本目录后保存，如图：</p>\n<p><img src=\"/content/images/2019/04/setting-search-4.png\" alt=\"setting-search-4\"></p>\n</li>\n<li>\n<p>测试索引效果，<code>Windows</code>能直接搜索到脚本</p>\n<p><img src=\"/content/images/2019/04/script-result-view.png\" alt=\"script-result-view\"></p>\n</li>\n</ol>\n<h2 id=\"\">结语</h2>\n<p>此篇笔记主要目的是留个记录，为了以后遇到 JDK 的配置能够有个自己的参考，也为了将自己的一些思考和经验分享出去。<br>\n笔者写的东西还少，实在难以把握内容的深浅度，写详细了显得繁琐，写简洁了又稍显简陋。<br>\n关于文中内容有任何问题或者建议，欢迎留言探讨。</p>\n<!--kg-card-end: markdown-->",
			Markdown:        "[^_^]: # (url:jdk-windows-setup)\\n[^_^]: # (tag:note,setup,jdk,java,windows,#tech)\\n[^_^]: # (excerpt:Windows 下配置多版本 JDK 环境)\\n\\n> Tested Environment  \\n> Operating System: Microsoft Windows Version 1903 (OS Build 18865.1000)  \\n> 谨以此文献给笔者活在虚拟机的 Windows\\n\\n## 概览\\n\\n`JDK`的设置主要有两点：\\n\\n- 下载正确的版本\\n  - 对应平台（Windows/Linux）\\n  - 对应架构（64bit/32bit）\\n  - 对应需求版本（1.8/11）\\n- 配置正确的环境变量\\n  - JAVA_HOME\\n  - Path\\n\\n## 下载 JDK\\n\\n建议同时下载`JDK8`和`JDK11`，无特殊需求建议使用`OpenJDK`，建议下载压缩包格式的（非`exe`或`msi`后缀）  \\n为方便读者查阅，以下几个网址是笔者推荐的`JDK`获取途径：\\n\\n- [Oracle JDK](https://www.oracle.com/technetwork/es/java/javasebusiness/downloads/index.html)\\n- [OpenJDK](https://openjdk.java.net/)\\n- [ojdkbuild](https://github.com/ojdkbuild/ojdkbuild)\\n\\n笔者使用的是 ojdkbuild 提供的 JDK，直接到[ojdkbuild](https://github.com/ojdkbuild/ojdkbuild)下载对应平台和版本的`JDK`，不再赘述。\\n\\n![ojdkbuild-download](/content/images/2019/04/ojdkbuild-download.png)\\n\\n## 解压 JDK\\n\\n1. 在`c盘`新建`Dev`文件夹，将下载的`JDK`解压到该文件夹下\\n2. 随后删除空层文件夹并更改`JDK`文件夹的名字，使其简短些\\n3. 注意文件夹层次结构（注意第二步的提示），理想情况下应该有如图所示有效路径：\\n\\n   ![jdk-folder](/content/images/2019/04/jdk-folder.png)\\n\\n## 设置环境变量\\n\\n为了后续的脚本切换`JDK`,此处使用的是用户环境变量，*非*系统环境变量，如非特殊说明，环境变量一律指代用户环境变量\\n\\n1. 新建 `JAVA_HOME_{version}` 变量,其中 version 为 JDK 版本,有多少个版本的 JDK 就建立多少个对应的环境变量,环境变量的值分别为对应的`JDK`的路径  \\n   例如: 本人使用两个版本的`JDK`,分别是 JDK8 和 11,则新建两个环境变量,分别是\\n\\n   | Variable     | Value                     |\\n   | ------------ | ------------------------- |\\n   | JAVA_HOME_8  | C:\\\\Dev\\\\Java\\\\OpenJDK1.8.0  |\\n   | JAVA_HOME_11 | C:\\\\Dev\\\\Java\\\\OpenJDK11.0.2 |\\n\\n   ![env-java-home-8](/content/images/2019/04/env-java-home-8.png)\\n   ![env-java-home-11](/content/images/2019/04/env-java-home-11.png)\\n\\n2. 新建 `JAVA_HOME` 变量,其值为上一步所新建的变量,格式为`%`包裹,见示例:\\n\\n   | Variable  | Value         |\\n   | --------- | ------------- |\\n   | JAVA_HOME | %JAVA_HOME_8% |\\n\\n   ![env-java-home](/content/images/2019/04/env-java-home.png)\\n\\n3. `Path`变量中新增一条记录,设置为`%JAVA_HOME%\\\\bin`\\n\\n   ![env-path](/content/images/2019/04/env-path.png)\\n\\n## 检查配置是否成功\\n\\n命令行运行`java -version`，应该输出`Java`的版本信息\\n\\n![cmd-java-version](/content/images/2019/04/cmd-java-version.png)\\n\\n\"}],[\"markdown\",{\"markdown\":\"## 编写版本切换脚本\\n\\n如需更改当前使用的`Java`版本，可以手动更改`JAVA_HOME`的变量值。如果更改`Java`版本的频次高，可以使用脚本来简化这一操作。  \\n笔者技拙，尝试过使用`bat脚本`更改系统环境变量，但这样脚本就需要管理员权限，且无法保证`JDK`的切换效果，遂更换思路，设置用户环境变量来达成目的。  \\n下列脚本一律不需要管理员权限，且只更改用户环境变量，在`Windows 10,v1903`测试通过。如果你的`Java`环境变量跟上文提到的设置方式一致，那么应该可以直接使用下面脚本。\\n新建一个文件夹，如`C:\\\\Scripts`，然后在该文件夹下新建这三个文件：\\n\\n1. 主脚本 `jdk_switch.bat`\\n2. 切换脚本 `switch_to_openjdk_8.bat`\\n3. 切换脚本 `switch_to_openjdk_11.bat`\\n\\n![scripts-folder](/content/images/2019/04/scripts-folder.png)\\n\\n### 主脚本 jdk_switch.bat\\n\\n(代码托管于Github Gist,只有聪明人才看得见#逃)\\n<script src=\\\"https://gist.github.com/nediiii/03f7daf70f6989c671e953eeeabab9fd.js\\\"></script>\"}],[\"markdown\",{\"markdown\":\"\\n\\n### 切换脚本 switch_to_openjdk_8.bat\\n\\n```batch\\nSETX JAVA_HOME \\\"%%JAVA_HOME_8%%\\\"\\n```\\n\\n### 切换脚本 switch_to_openjdk_11.bat\\n\\n```batch\\nSETX JAVA_HOME \\\"%%JAVA_HOME_11%%\\\"\\n```\\n\\n### 脚本效果\\n\\n运行脚本的效果如图所示：\\n\\n![script-view](/content/images/2019/04/script-view.png)\\n\\n### 添加环境变量\\n\\n为了在命令行中能够直接使用脚本，需要把脚本目录添加进 Path 变量里。过程不再赘述。如图：\\n\\n![env-path](/content/images/2019/04/env-path.png)\\n\\n此后可在命令行中直接运行脚本：\\n\\n![cmd-script](/content/images/2019/04/cmd-script.png)\\n\\n### 设置索引\\n\\n为了使用 Windows 搜索时能直接搜索到脚本，需要把脚本目录添加进索引目录里。\\n要进入`索引设置`，有两个方式（下列第 1 步和第 2 步最终效果相同，皆是为了进入第 3 步）\\n\\n1. 按`Windows键`然后键入`change how windows`，如图所示，点击图中标示的选项（如有问题，可从第 2 步开始）：\\n\\n   ![windows-search](/content/images/2019/04/windows-search.png)\\n\\n2. 从`设置`->`搜索`里进入，如图：\\n\\n   ![setting-search-1](/content/images/2019/04/setting-search-1.png)\\n   ![setting-search-2](/content/images/2019/04/setting-search-2.png)\\n\\n3. 进入到了索引设置，点击修改，如图：\\n\\n   ![setting-search-3](/content/images/2019/04/setting-search-3.png)\\n\\n4. 勾选上脚本目录后保存，如图：\\n\\n   ![setting-search-4](/content/images/2019/04/setting-search-4.png)\\n\\n5. 测试索引效果，`Windows`能直接搜索到脚本\\n\\n   ![script-result-view](/content/images/2019/04/script-result-view.png)\\n\\n## 结语\\n\\n此篇笔记主要目的是留个记录，为了以后遇到 JDK 的配置能够有个自己的参考，也为了将自己的一些思考和经验分享出去。  \\n笔者写的东西还少，实在难以把握内容的深浅度，写详细了显得繁琐，写简洁了又稍显简陋。  \\n关于文中内容有任何问题或者建议，欢迎留言探讨。",
		},
		{
			Authors:         PreDefinedUsers,
			Tags:            PreDefinedTags,
			PrimaryAuthor:   PreDefinedUsers[0],
			PrimaryAuthorID: PreDefinedUsers[0].ID,
			Title:           "用 portainer 优雅管理 docker 容器",
			Slug:            "docker-gui-app",
			Excerpt:         "厌倦了命令行？这里有漂亮的 GUI 应用助你管理 Docker",
			Featured:        false,
			Paged:           false,
			Status:          "published",
			HTML:            "<!--kg-card-begin: markdown--><blockquote>\n<p>笔者自从接触了 Docker 后就深深沉迷于它，然当跑的应用多了，众多的容器难免眼花缭乱。便去搜寻相关的解决方案，在此将信息整合分享。<br>\n如果你也厌倦了命令行，希望尝试一下 GUI 应用，那么这篇文章值得一看。</p>\n</blockquote>\n<h2 id=\"dockerui\">主流的 Docker ui</h2>\n<ul>\n<li><a href=\"https://www.portainer.io\">Portainer</a></li>\n<li><a href=\"https://kitematic.com/\">Kitematic</a></li>\n<li><a href=\"https://dockstation.io/\">DockStation</a></li>\n</ul>\n<p><em>你可能会奇怪为什么<code>Shipyard</code>没有在列表上，这里明确告诉你，<code>Shipyard</code>项目已经停止开发，所以笔者将其移除了</em></p>\n<h2 id=\"\">对比</h2>\n<table>\n<thead>\n<tr>\n<th>功能</th>\n<th>DockerStation</th>\n<th>Kitematic</th>\n<th>Portainer</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>平台支持</td>\n<td>全平台</td>\n<td>不支持 linux</td>\n<td>全平台</td>\n</tr>\n<tr>\n<td>Docker Compose 支持</td>\n<td>+</td>\n<td>-</td>\n<td>-</td>\n</tr>\n<tr>\n<td>Docker Machine 支持</td>\n<td>+</td>\n<td>+</td>\n<td>+</td>\n</tr>\n<tr>\n<td>任意容器配置</td>\n<td>-</td>\n<td>+</td>\n<td>+</td>\n</tr>\n<tr>\n<td>容器启动/停止/重启</td>\n<td>+</td>\n<td>+</td>\n<td>+</td>\n</tr>\n<tr>\n<td>查看容器日志</td>\n<td>+</td>\n<td>+</td>\n<td>+</td>\n</tr>\n<tr>\n<td>日志全文搜索</td>\n<td>+</td>\n<td>-</td>\n<td>-</td>\n</tr>\n<tr>\n<td>容器分组和搜索</td>\n<td>+</td>\n<td>-</td>\n<td>-</td>\n</tr>\n<tr>\n<td>资源监控</td>\n<td>+</td>\n<td>-</td>\n<td>+</td>\n</tr>\n<tr>\n<td>远程节点支持</td>\n<td>+</td>\n<td>-</td>\n<td>+</td>\n</tr>\n<tr>\n<td>应用模板</td>\n<td>-</td>\n<td>-</td>\n<td>+</td>\n</tr>\n<tr>\n<td>自定义 hubs</td>\n<td>-</td>\n<td>-</td>\n<td>+</td>\n</tr>\n</tbody>\n</table>\n<h2 id=\"portainer\">Portainer 安装</h2>\n<blockquote>\n<p>详情参阅<a href=\"https://www.portainer.io/installation/\">https://www.portainer.io/installation/</a></p>\n</blockquote>\n<p>因<code>Portainer</code>是个 Web 应用，其自然可以轻松以<code>docker</code>方式运行。<br>\nlinux 下执行下列命令（其他平台请查看官方文档，在此不赘述）：</p>\n<pre><code class=\"language-bash\"># 新建 volume\ndocker volume create portainer_data\n# 运行容器\ndocker run -d -p 9000:9000 \\\n-v /var/run/docker.sock:/var/run/docker.sock \\\n-v portainer_data:/data \\\nportainer/portainer\n</code></pre>\n<h2 id=\"portainer\">Portainer 使用</h2>\n<p>如<code>Portainer</code>在本地运行，则登录 <a href=\"localhost:9000\">localhost:9000</a> 即可访问<code>Portainer</code>应用，否则请输入对应的<code>ip:port</code>访问，笔者将其运行在本地做演示。</p>\n<h3 id=\"\">注册</h3>\n<p>登录<a href=\"localhost:9000\">localhost:9000</a>后，会跳转到初始化页面，要求您创建一个管理员账号，如图：</p>\n<p><img src=\"/content/images/2019/04/portainer-register.png\" alt=\"portainer-register\"></p>\n<p>填写好账号信息后，点击左下角创建用户（Create user)，注册工作就完成了。</p>\n<h3 id=\"\">连接</h3>\n<p>注册账号之后进入Docker环境选择界面，选中远端环境（Remote）还需要填写额外的地址和认证信息，这里笔者选择本地环境（<code>Local</code>）做演示（此选项后期可修改，亦可添加多个Docker，在此不需要过多纠结），如图：<br>\n<img src=\"/content/images/2019/04/portainer-connect.png\" alt=\"portainer-connect\"></p>\n<h3 id=\"\">主界面</h3>\n<p>成功连接Docker之后便进入到应用的主界面了，展示了关于Docker的概览信息，如图：</p>\n<p><img src=\"/content/images/2019/04/portainer-main-page.png\" alt=\"portainer-main-page\"></p>\n<p>更多功能诸如多用户，多端点，多注册器等功能就等读者自己去体验了。</p>\n<h2 id=\"dockstation\">DockStation 安装</h2>\n<blockquote>\n<p>详情参阅<a href=\"https://dockstation.io/\">https://dockstation.io/</a></p>\n</blockquote>\n<p>与<code>Portainer</code>不同，<code>DockStation</code>是一个GUI应用，也就不太适合用<code>Docker</code>来运行。<br>\n这里采用的安装其应用包的方式将其当成一个应用运行在宿主机上。<br>\n如笔者的<code>ubuntu18</code>，只需下载对应的<a href=\"https://github.com/DockStation/dockstation/releases/download/v1.5.0/dockstation_1.5.0_amd64.deb\">deb包</a>，然后安装即可。</p>\n<h2 id=\"dockstation\">DockStation 使用</h2>\n<p>如图所示，应用UI风格偏MAC，功能不算复杂，熟悉<code>Docker</code>的用户很容易能理解它。有兴趣的朋友可以入手玩一玩。</p>\n<p><img src=\"/content/images/2019/04/dockerstation.png\" alt=\"dockerstation\"></p>\n<!--kg-card-end: markdown-->",
			Markdown:        "[^_^]: # (url:docker-gui-app)\\n[^_^]: # (tag:docker,#tech)\\n[^_^]: # (excerpt:厌倦了命令行？这里有漂亮的 GUI 应用助你管理 Docker)\\n\\n> 笔者自从接触了 Docker 后就深深沉迷于它，然当跑的应用多了，众多的容器难免眼花缭乱。便去搜寻相关的解决方案，在此将信息整合分享。  \\n> 如果你也厌倦了命令行，希望尝试一下 GUI 应用，那么这篇文章值得一看。\\n\\n## 主流的 Docker ui\\n\\n- [Portainer](https://www.portainer.io)\\n- [Kitematic](https://kitematic.com/)\\n- [DockStation](https://dockstation.io/)\\n\\n_你可能会奇怪为什么`Shipyard`没有在列表上，这里明确告诉你，`Shipyard`项目已经停止开发，所以笔者将其移除了_\\n\\n## 对比\\n\\n| 功能                | DockerStation | Kitematic    | Portainer |\\n| ------------------- | ------------- | ------------ | --------- |\\n| 平台支持            | 全平台        | 不支持 linux | 全平台    |\\n| Docker Compose 支持 | +             | -            | -         |\\n| Docker Machine 支持 | +             | +            | +         |\\n| 任意容器配置        | -             | +            | +         |\\n| 容器启动/停止/重启  | +             | +            | +         |\\n| 查看容器日志        | +             | +            | +         |\\n| 日志全文搜索        | +             | -            | -         |\\n| 容器分组和搜索      | +             | -            | -         |\\n| 资源监控            | +             | -            | +         |\\n| 远程节点支持        | +             | -            | +         |\\n| 应用模板            | -             | -            | +         |\\n| 自定义 hubs         | -             | -            | +         |\\n\\n## Portainer 安装\\n\\n> 详情参阅[https://www.portainer.io/installation/](https://www.portainer.io/installation/)\\n\\n因`Portainer`是个 Web 应用，其自然可以轻松以`docker`方式运行。  \\nlinux 下执行下列命令（其他平台请查看官方文档，在此不赘述）：\\n\\n```bash\\n# 新建 volume\\ndocker volume create portainer_data\\n# 运行容器\\ndocker run -d -p 9000:9000 \\\\\\n-v /var/run/docker.sock:/var/run/docker.sock \\\\\\n-v portainer_data:/data \\\\\\nportainer/portainer\\n```\\n\\n## Portainer 使用\\n\\n如`Portainer`在本地运行，则登录 [localhost:9000](localhost:9000) 即可访问`Portainer`应用，否则请输入对应的`ip:port`访问，笔者将其运行在本地做演示。\\n\\n### 注册\\n\\n登录[localhost:9000](localhost:9000)后，会跳转到初始化页面，要求您创建一个管理员账号，如图：\\n\\n![portainer-register](/content/images/2019/04/portainer-register.png)\\n\\n填写好账号信息后，点击左下角创建用户（Create user)，注册工作就完成了。\\n\\n### 连接\\n\\n注册账号之后进入Docker环境选择界面，选中远端环境（Remote）还需要填写额外的地址和认证信息，这里笔者选择本地环境（`Local`）做演示（此选项后期可修改，亦可添加多个Docker，在此不需要过多纠结），如图：\\n![portainer-connect](/content/images/2019/04/portainer-connect.png)\\n\\n### 主界面\\n\\n成功连接Docker之后便进入到应用的主界面了，展示了关于Docker的概览信息，如图：\\n\\n![portainer-main-page](/content/images/2019/04/portainer-main-page.png)\\n\\n更多功能诸如多用户，多端点，多注册器等功能就等读者自己去体验了。\\n\\n## DockStation 安装\\n\\n> 详情参阅[https://dockstation.io/](https://dockstation.io/)\\n\\n与`Portainer`不同，`DockStation`是一个GUI应用，也就不太适合用`Docker`来运行。  \\n这里采用的安装其应用包的方式将其当成一个应用运行在宿主机上。\\n如笔者的`ubuntu18`，只需下载对应的[deb包](https://github.com/DockStation/dockstation/releases/download/v1.5.0/dockstation_1.5.0_amd64.deb)，然后安装即可。\\n\\n## DockStation 使用\\n\\n如图所示，应用UI风格偏MAC，功能不算复杂，熟悉`Docker`的用户很容易能理解它。有兴趣的朋友可以入手玩一玩。\\n\\n![dockerstation](/content/images/2019/04/dockerstation.png)\\n",
		},
		{
			Authors:         PreDefinedUsers,
			Tags:            PreDefinedTags,
			PrimaryAuthor:   PreDefinedUsers[0],
			PrimaryAuthorID: PreDefinedUsers[0].ID,
			Title:           "demo-blog-title",
			Slug:            "demo-blog-slug",
			Excerpt:         "demo-blog-excerpt",
			Featured:        false,
			Paged:           false,
			Status:          "published",
			HTML:            "HTML CONTENT",
		},
	}
)
