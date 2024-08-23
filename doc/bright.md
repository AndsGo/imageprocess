# 亮度

您可以通过亮度参数，调节存储在OSS内的原图亮度。本文介绍调节图片亮度时所用到的参数及示例。

## 参数说明

操作名称：**bright**

参数说明如下：

| **参数**    | **描述**         | **取值范围**                                                 |
| ----------- | ---------------- | ------------------------------------------------------------ |
| **[value]** | 指定图片的亮度。 | [-100, 100]取值＜0：降低图片亮度。取值=0：不调整图片亮度。取值＞0：提高图片亮度。 |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例](../README.md#comprehensive)

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg
![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8532220761/p529913.jpg)

- 将图片亮度提高50

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/bright,50![亮度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7532220761/p529917.jpg)

- 将图片亮度降低50

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/bright,-50