## Cloud Services

* **AWS Lambda**: The main script runs on AWS Lambda.

* **AWS Cloudwatch**: Cloudwatch sends weekly triggers to the lambda function. Can be configured to send daily, monthly, or multiple combinations of triggers.

* **AWS Simple Email Service**: The service used to send emails. You can also use the Gmail API with some minor changes in the code. 

## Environment Variables

* `EMAIL_FROM`: Email address used to send emails. 
* `EMAIL_NAME`, `EMAIL_PASS`: AWS SeS authentication.
* `FEEDS_SOURCE`: The feeds are read from a .txt file at this URL. 
* `RSS_TARGET`: Email recipient.
* `TEMPLATE_S3_BUCKET`: URL of the S3 bucket storing the templates.