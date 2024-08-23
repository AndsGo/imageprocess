# Blur

## **Scenarios**

- Privacy protection: Before you publish an image that contains sensitive information, you can blur part of the image.
- Better visual experience for an image composed of multiple layers: Properly blurring the image can smooth the edges between different layers to provide a more comfortable visual experience.
- Better display of a low-resolution image: When the resolution of an image is low and cannot meet the requirements of a high-definition display, you can blur the image.

## Parameters

Operation: **blur**

Blur generates a blurred version of the image using a Gaussian function.

The parameter must be a positive number, indicating the blurriness of the image.

The following table describes the parameters.

| **Parameter** | **Required** | **Description**                   | **Valid value**                                   |
| ------------- | ------------ | --------------------------------- | ------------------------------------------------- |
| [value]       | Yes          | Set Gaussian function parameters. | [1,50]A greater value specifies a blurrier image. |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/8850665071/p674595.jpg)

- Blur an image

  Configure the parameters based on the following requirements: If you want to set the blur radius to 10 and the standard deviation of a normal distribution to 10, add `5` to the URL of the image.

  The following URL is used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/blur,5![模糊1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/8850665071/p674663.jpg)

# 