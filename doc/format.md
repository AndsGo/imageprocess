# 格式转换

通过格式转换参数，您无需将图片下载到本地进行转换，只需指定URL转换存储在OSS内原图的格式。本文介绍对图片进行格式转换时所用到的参数及示例。

## **使用场景**

- 适应不同设备和平台：不同的浏览器、操作系统或移动设备可能支持不同的图片格式。例如，WebP格式在浏览器上能提供更好的压缩效率，通过OSS图片格式转换，可以将上传的图片转换为多种格式，确保在各种终端上的兼容性和最佳显示效果。
- 节省存储成本：某些图片格式（如WebP）在保证视觉质量的前提下，文件大小通常小于其他传统格式。通过格式转换功能，可以在不牺牲图像质量的情况下减少存储空间占用，从而降低存储成本。
- 统一资源管理：在电商、社交、媒体等行业中，需要对大量用户上传的图片进行标准化处理。您可以将上传的图片统一转换成指定格式，便于后续的一致管理和分发。

## 注意事项

- 图片处理包含缩放操作时，建议将格式转换参数放到处理参数的最后。

  例如`image/resize,w_100/format,jpg`

- 图片处理包含缩放和水印操作时，建议将格式转换参数添加在缩放参数之后。

  例如`image/resize,w_100/format,jpg/watermark,...`

- 如果原图没有透明通道，转换成PNG、Web、BMP等存在透明通道的格式，默认会把透明填充成白色。

- OSS不支持将透明色填充为黑色。

## 参数说明

操作名称：**format**

参数说明如下：

| **取值范围** | **描述**                                                     |
| ------------ | ------------------------------------------------------------ |
| **jpg**      | 将原图保存为JPG格式。**重要**不支持将存在透明通道的HEIC格式的图片保存为JPG格式。 |
| **png**      | 将原图保存为PNG格式。                                        |
| **webp**     | 将原图保存为WebP格式。                                       |
| **bmp**      | 将原图保存为BMP格式。                                        |
| **gif**      | 原图为GIF图片则继续保存为GIF格式；原图不是GIF图片，则按原图格式保存。 |
| **tiff**     | 将原图保存为TIFF格式。                                       |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.gif，

该测试基于 

[综合示例](../README.md#comprehensive)

图片访问地址为：

http://127.0.0.1:8080/file/example.gif

![gif](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139212.png)



- 将原图转换为PNG格式

  图片处理URL为：http://127.0.0.1:8080/file/example.gif?x-oss-process=image/format,png![png](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139213.png)

- 将原图转换成JPG格式，并支持渐进显示

  需求及处理参数如下：

  - 图片设置为渐进显示：`interlace,1`
  - 图片转换为JPG格式：`format,jpg`

  图片处理URL为：http://127.0.0.1:8080/file/example.gif?x-oss-process=image/format,jpg

  ![img](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0548459951/p2555.jpg)

- 将原图缩放为宽200 px，并转换为WebP格式

  需求及处理参数如下：

  - 图片缩放为宽200 px：`resize,w_200`
  - 图片转换为WebP格式：`format,webp`

  图片处理URL为：http://127.0.0.1:8080/file/example.gif?x-oss-process=image/resize,w_200,h_100/format,webp

目前转换为webp还不支持动图。