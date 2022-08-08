# wheretogo
go practice

# 项目创建步骤
1.创建项目文件夹
2.项目文件夹目录下创建thrift文件,具体名字自定义,用于自动生成项目
.idl/hello.thrift
3.执行以下命令
hz new -idl idl/hello.thrift -mod namespace
4.整理 & 拉取依赖
go mod tidy
5.编译