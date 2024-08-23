# 旋转

您可以通过旋转参数，将存储在OSS内的原图按指定方向旋转。本文介绍旋转图片时所用到的参数和示例。

## 参数说明

操作名称：**rotate**

参数说明如下：

| **参数**    | **描述**                 | **取值范围**                   |
| ----------- | ------------------------ | ------------------------------ |
| **[value]** | 图片按顺时针旋转的角度。 | [0,360]默认值：0，表示不旋转。 |

## 注意事项

- 若图片旋转的角度不是90°、180°、270°、360°时，会导致处理后的图片尺寸变大。
- 旋转功能对图片的尺寸有限制，图片的宽或者高不能超过4096 px。
- GIF图片不支持旋转参数。如果对GIF动图执行旋转操作，GIF动图会变成GIF静态图。

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例](../README.md#comprehensive)

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529611.jpg)

- 将原图按顺时针旋转90°

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/rotate,90

  ![旋转1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529612.jpg)

- 将原图按顺时针旋转70°

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/rotate,70