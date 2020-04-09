# SQL-DSL转换器

ElasticSearch是一个很好用的全文搜索引擎, 但是他的搜索语言(DSL)与我们常用的sql存在很大的差异. 

本项目实现了将sql语句转换为dsl语句的功能, 能够模拟大部分的dsl查询. 

sql转换模块基于 [elasticsql](https://github.com/cch123/elasticsql) 开发, 再其基础上添加了一些dsl的语法(例如exist).

目前有两个版本:

- 控制台版本: 命令行执行, master分支
- web版本: 提供了一个web页面, web_v1.0.1分支

## 1. 使用

### 1.1. Command

建议将编译的二进制文件拷贝到`$PATH/bin`, 可以在任意目录使用

```bash
dslConvert "select * from myindex where (params='test' or address='Asia') and gmt_time between '2020.03.20 14:30:40.234' and '2020.03.22 14:30:40.234'" [-p]
```

可选参数: 

- `-p` : 格式化输出

### 1.2. Web

默认使用`3232`端口, 支持功能:

- 自动补全, Tab键唤出
- 高亮显示
- sql语句格式化

页面展示:

![dslConvert.Web](https://i.loli.net/2020/03/20/euDcBOknydNptHE.jpg)

## 2. 例子

查看example


## 3. 说明

elasticsql提醒:

>To use this tool, you need to understand the term query and match phrase query of elasticsearch.
>
>Setting a field to analyzed or not analyzed will get different results.

虽然这个工具能实现一定程度上的sql -> dsl转换, 但由于dsl的写法较sql复杂了很多, 所以如果想要最好效果的dsl语句还是要自己去学习他的语法.

本工具只是提供了一种简单方法来供一些想要使用es但还未深入学习的人, 亦或者是es初学者理解sql和dsl的联系.

有什么问题或者建议欢迎提Issue or PR.



## License

MIT