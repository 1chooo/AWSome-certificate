# SAA Exam Practice

**Table of Contents:**

- [SAA Exam Practice](#saa-exam-practice)
  - [Highly available](#highly-available)
  - [Unpredictable Workloads](#unpredictable-workloads)
  - [Improve Performance](#improve-performance)
  - [Cost-effective](#cost-effective)
  - [Millions of Views](#millions-of-views)

## Highly available

A company has created a VPC with multiple private subnets in multiple Availability Zones (AZs) and one public subnet in one of the AZs. The public subnet is used to launch a NAT gateway. There are instances in the private subnets that use a NAT gateway to connect to the internet. In case of an AZ failure, the company wants to ensure that the instances are not all experiencing internet connectivity issues and that there is a backup plan ready.  
Which solution should a solutions architect recommend that is MOST highly available?
- [ ] Create a new public subnet with a NAT gateway in the same AZ. Distribute the traffic between the two NAT gateways.
- [ ] Create an Amazon EC2 NAT instance in a new public subnet. Distribute the traffic between the NAT gateway and the NAT instance.
- [x] Create public subnets in each AZ and launch a NAT gateway in each subnet. Configure the traffic from the private subnets in each AZ to the respective NAT gateway.
- [ ] Create an Amazon EC2 NAT instance in the same public subnet. Replace the NAT gateway with the NAT instance and associate the instance with an Auto Scaling group with an appropriate scaling policy.


## Unpredictable Workloads

A retail website has intermittent, sporadic, and unpredictable transactional workloads throughout the day that are hard to predict.  
The website is currently hosted on-premises and is slated to be migrated to AWS.   
A new relational database is needed that autoscales capacity to meet the needs of the application's peak load and scales back down when the surge of activity is over.   
Which of the following option is the MOST cost-effective and suitable database setup in this scenario?  
- [x] Launch an Amazon Aurora Serverless DB cluster then set the minimum anœœd maximum capacity for the cluster.
- [ ] ​Launch a DynamoDB Global table with Auto Scaling enabled.
- [ ] ​Launch an Amazon Aurora Provisioned DB cluster with burstable performance DB instance class types.
- [ ] ​Launch an Amazon Redshift data warehouse cluster with Concurrency Scaling.

A solutions architect is designing a new service behind Amazon API Gateway. The request patterns for the service will be unpredictable and can change suddenly from 0 requests to over 500 per second. The total size of the data that needs to be persisted in a backend database is currently less than 1 GB with unpredictable future growth. Data can be queried using simple key-value requests.  
Which combination of AWS services would meet these requirements? **(Choose two.)**
- [ ] AWS Fargate
- [x] AWS Lambda
- [x] Amazon DynamoDB
- [ ] Amazon EC2 Auto Scaling
- [ ] MySQL-compatible Amazon Aurora

> [!NOTE]
> unpredictable -> link to the serverless and the distinguish of the Lambda Function



## Improve Performance

An ecommerce company is running a multi-tier application on AWS. The front-end and backend tiers both run on Amazon EC2, and the database runs on Amazon  
RDS for MySQL. The backend tier communicates with the RDS instance. There are frequent calls to return identical datasets from the database that are causing performance slowdowns.  
Which action should be taken to improve the performance of the backend?
- [ ] Implement Amazon SNS to store the database calls.
- [x] Implement Amazon ElastiCache to cache the large datasets.
- [ ] Implement an RDS for MySQL read replica to cache database calls.
- [ ] Implement Amazon Kinesis Data Firehose to stream the calls to the database.

> [!IMPORTANT]
> `replica` cannot cache the data.

## Cost-effective

A company hosted a web application in an auto scaling group of ec2 instances.  
The IT manager is concerned about the over-provisioning of the resources that can cause higher operating costs.  
A Solutions Architect has been instructed to create a cost-effective solution without affecting the performance of the application.  
Which dynamic scaling policy should be used to satisfy this requirement?

- [x] Use target tracking scaling.


## Millions of Views

Organizers for a global event want to put daily reports online as static HTML pages. The pages are expected to generate millions of views from users around the world. The files are stored in an Amazon S3 bucket. A solutions architect has been asked to design an efficient and effective solution.

Which action should the solutions architect take to accomplish this?
- [ ] Generate presigned URLs for the files.
- [ ] Use cross-Region replication to all Regions.
- [ ] Use the geoproximity feature of Amazon Route 53.
- [x] Use Amazon CloudFront with the S3 bucket as its origin.



















