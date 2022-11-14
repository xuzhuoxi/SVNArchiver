# SVNArchiver

## Release Notes

+ 提供使用xml配置进行批量归档后，增加归档信息文件的输出功能。

### Improvements

+ 增加归档信息文件的输出。

+ 归档信息文件中会记录每一个归档任务生成的归档文件对应的文件名称、特征码和详细路径。

+ 支持json和xml两种格式的归档信息文件。

+ 支持md5和sha1两种算法的特征码。

### Changes

+ xml配置增加<log>标签，用于配置归档信息文件的输出。

+ xml配置中标签<task>增加属性参数id，用于标识每一个归档任务。
