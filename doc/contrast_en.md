# Contrast

Updated at: 2022-12-30 16:02

Contrast refers to the measurement of different brightness levels between the brightest white and the darkest black of an image, that is, the grayscale contrast of an image. You can use contrast parameter to adjust the contrast of the source images stored in OSS. This topic describes the parameters and examples to adjust the contrast for an image.

## Parameters

Operation name: **contrast**

| Parameter   | Description                | Valid value                                                  |
| :---------- | :------------------------- | :----------------------------------------------------------- |
| **[value]** | The contrast of the image. | [-100,100]A value smaller than 0: reduces the contrast.A value of 0: maintains the contrast.A value greater than 0: increases the contrast. |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg
![1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0370040761/p531719.jpg)

- Reduce the contrast by 50

  The URL used to process the image is in the following format: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/contrast,-50

  ![2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0370040761/p531721.jpg)

- Increase the contrast by 50

  The URL used to process the image is in the following format: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/contrast,50

  ![3](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0370040761/p531722.jpg)