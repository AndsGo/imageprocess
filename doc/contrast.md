# 对比度

对比度是指一幅图像中明暗区域最亮的白和最暗的黑之间不同亮度层级的测量，即指一幅图像灰度反差的大小。您可以通过对比度参数，调整存储在OSS内原图的对比度。本文介绍调节图片对比度时所用到的参数及示例。

## 参数说明

操作名称：**contrast**

| **参数**    | **描述**           | **取值范围**                                                 |
| ----------- | ------------------ | ------------------------------------------------------------ |
| **[value]** | 指定图片的对比度。 | [-100,100]取值＜0：降低图片对比度。取值=0：维持原图对比度。取值＞0：提高图片对比度。 |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例](../README.md#comprehensive)

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529934.jpg)

- 对比度降低50

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/contrast,-50

  ![对比度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529935.jpg)

- 对比度提高50

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/contrast,50

![对比度2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529938.jpg)