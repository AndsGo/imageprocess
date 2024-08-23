# Adjust image quality

Updated at: 2022-12-30 15:54

The quality adjustment operation uses the format of a source image to compress the image. You can use the quality adjustment parameters to modify the quality of source images stored in Object Storage Service (OSS). This topic describes the parameters and examples for image quality adjustment.

Quality adjustment applies only to JPG and WebP images.

## Parameters

Operation name: **quality**

The following table describes the parameters that you can configure when you adjust the quality of an image.

| Parameter | Description                                                  | Valid value |
| :-------- | :----------------------------------------------------------- | :---------- |
| **q**     | Specifies the relative quality of the image and compresses the source image based on percentage.If the source image quality is 100%, you can obtain an image whose quality value is 90% after you add the `quality,q_90` parameter. If the quality value of the source image is 80%, you can obtain an image whose quality value is 72% after you add the `quality,q_90` parameter.**Note**The q parameter applies only to source images in the JPG format to specify the relative quality of the images. If a source image is in the WebP format, this parameter works the same as **Q**. The absolute quality is specified for the image. | [1,100]     |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg

![1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/7024640761/p532019.jpg)

- Adjust the relative quality of an image

  Configure the parameters based on the following requirements:

  - Resize the image to a width of 100 pixels: `resize,w_100`
  - Set the relative quality value of the image to 80%: `quality,q_80`

  The URL used to process the image is in the following format: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_100/quality,q_80

  ![2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/7024640761/p532021.jpg)