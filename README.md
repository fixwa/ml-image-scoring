# Image Tagger 🔮

A set of applications (UI, Image Parser, ResNet50 model scoring API) that allow an analysis of images to obtain a "tag" for said image.

It consists of a UI in ReactJS to upload images.
This is a basic application with a form to submit the file to the FileHandlerAPI.

FileHandlerAPI (fh_api) is a golang application to get the image file and parse (standardize/decode) it to a tensor model for scoring.

ScoringAPI (scoring_api) a deeplearning service that uses the ResNet50 architecture to score a TF model.


```
docker-compose up
```