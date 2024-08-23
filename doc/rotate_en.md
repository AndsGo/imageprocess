# Rotate

## Parameters

Operation name: **rotate**

The following table describes the parameters you can configure.

| Parameter   | Description                                         | Valid value                                                  |
| :---------- | :-------------------------------------------------- | :----------------------------------------------------------- |
| **[value]** | The degree by which the image is rotated clockwise. | [0,360]Default value: 0. A value of 0 indicates that the image is not rotated. |

## Usage notes

- If an image is not rotated by 90°, 180°, 270°, or 360°, the size of the processed image increases.
- An image that you want to rotate cannot exceed 4096 × 4096 pixels.
- If rotation parameters are applied to an animated GIF image, the animated GIF image becomes a static GIF image.

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg

![1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0300040761/p531686.jpg)

- Rotate the source image by 90 degrees clockwise.

  The URL used to rotate the source image by 90 degrees clockwise is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/rotate,90.

  ![2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0300040761/p531687.jpg)

- Rotate the source image by 70 degrees clockwise

  The URL used to rotate the source image by 70 degrees clockwise is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/rotate,70.