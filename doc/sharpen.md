# 锐化

您可以通过锐化参数，提高存储在OSS内原图的清晰度。本文介绍对图片进行锐化时所用到的参数及示例。

## 参数说明

操作名称：**sharpen**

参数说明如下：

| **参数**    | **描述**             | **取值范围**                                                 |
| ----------- | -------------------- | ------------------------------------------------------------ |
| **[value]** | 设置锐化效果的强度。 | [50,399]取值越大，图片越清晰，但过大的值可能会导致图片失真。为达到较优效果，推荐取值为100。 |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例](../README.md#comprehensive)

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529928.jpg)



对原图进行锐化处理，锐化参数为100。图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/sharpen,100![锐化1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529929.jpg)