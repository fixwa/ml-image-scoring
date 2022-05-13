# MachineLearning Image Scoring 🔮

A set of applications (UI, Image Parser, ResNet50 model scoring API) that allow an analysis of images to obtain a "tag" for said image.

It consists of
## UI (webapp) 
A very basic webapp in ReactJS to upload images.
This is a form to submit the file to the FileHandlerAPI and display the results of the scoring analysis.

## FileHandlerAPI (fh_api) 
is a golang application to get the image file and parse (standardize/decode) it to a tensor model for scoring.

## ScoringAPI (scoring_api) 
a deeplearning service that uses the ResNet50 architecture to score a TF model. This is based on the work of [Geert Baeke](https://blog.baeke.info/2019/01/02/creating-a-gpu-container-image-for-scoring-with-azure-machine-learning/)
Consist of a container that uses the [Azure Machine Learning comuting service](https://docs.microsoft.com/en-us/azure/machine-learning/quickstart-create-resources) and a Python script that uses the [Keras Applications - ResNetV2](https://keras.io/api/applications/resnet/#resnet50v2-function) library to receive the 4D array payload and process (classify) the image.

To see all services in action:
```
docker-compose up
```

Use the images provided in the "/example_images" folder or some of your own.
