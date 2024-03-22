# SAA Exam Practice

A company has created a VPC with multiple private subnets in multiple Availability Zones (AZs) and one public subnet in one of the AZs. The public subnet is used to launch a NAT gateway. There are instances in the private subnets that use a NAT gateway to connect to the internet. In case of an AZ failure, the company wants to ensure that the instances are not all experiencing internet connectivity issues and that there is a backup plan ready.  

Which solution should a solutions architect recommend that is MOST highly available?

##### Keywords: `Highly available`, `EC2`, `VPC`

- [ ] Create a new public subnet with a NAT gateway in the same AZ. Distribute the traffic between the two NAT gateways.
- [ ] Create an Amazon EC2 NAT instance in a new public subnet. Distribute the traffic between the NAT gateway and the NAT instance.
- [x] Create public subnets in each AZ and launch a NAT gateway in each subnet. Configure the traffic from the private subnets in each AZ to the respective NAT gateway.
- [ ] Create an Amazon EC2 NAT instance in the same public subnet. Replace the NAT gateway with the NAT instance and associate the instance with an Auto Scaling group with an appropriate scaling policy.

---

A retail website has intermittent, sporadic, and unpredictable transactional workloads throughout the day that are hard to predict.  
The website is currently hosted on-premises and is slated to be migrated to AWS.   
A new relational database is needed that autoscales capacity to meet the needs of the application's peak load and scales back down when the surge of activity is over.   

Which of the following option is the MOST **cost-effective and suitable database** setup in this scenario?  

##### Keywords: `Unpredictable Workloads`, `cost-effective`, `DB`

- [x] Launch an Amazon Aurora Serverless DB cluster then set the minimum and maximum capacity for the cluster.
- [ ] ​Launch a DynamoDB Global table with Auto Scaling enabled.
- [ ] ​Launch an Amazon Aurora Provisioned DB cluster with burstable performance DB instance class types.
- [ ] ​Launch an Amazon Redshift data warehouse cluster with Concurrency Scaling.

---

A solutions architect is designing a new service behind Amazon API Gateway. The request patterns for the service will be unpredictable and can change suddenly from 0 requests to over 500 per second. The total size of the data that needs to be persisted in a backend database is currently less than 1 GB with unpredictable future growth. Data can be queried using simple key-value requests.  

Which combination of AWS services would meet these requirements? **(Choose two.)**

##### Keywords: `Serverless`, `Unpredictable`, `DB`

- [ ] AWS Fargate
- [x] AWS Lambda
- [x] Amazon DynamoDB
- [ ] Amazon EC2 Auto Scaling
- [ ] MySQL-compatible Amazon Aurora

> [!TIP]
> unpredictable -> link to the serverless and the distinguish of the Lambda Function

---

An ecommerce company is running a multi-tier application on AWS. The front-end and backend tiers both run on Amazon EC2, and the database runs on Amazon RDS for MySQL. The backend tier communicates with the RDS instance. There are frequent calls to return identical datasets from the database that are causing performance slowdowns.  

Which action should be taken to improve the performance of the backend?

##### Keywords: `Improve Performance`, `Cache`, `replica`

- [ ] Implement Amazon SNS to store the database calls.
- [x] Implement Amazon ElastiCache to cache the large datasets.
- [ ] Implement an RDS for MySQL read replica to cache database calls.
- [ ] Implement Amazon Kinesis Data Firehose to stream the calls to the database.

> [!IMPORTANT]
> `replica` cannot cache the data.

---

A company hosted a web application in an auto scaling group of ec2 instances.  
The IT manager is concerned about the over-provisioning of the resources that can cause higher operating costs.  
A Solutions Architect has been instructed to create a **cost-effective** solution without affecting the performance of the application. 

Which dynamic scaling policy should be used to satisfy this requirement?

##### Keywords: `Cost-effective`, `Auto Scaling`, `Dynamic Scaling Policy`

- [ ] Use simple scaling.
- [ ] Use scheduled scaling.
- [ ] Use suspend and resume scaling.
- [x] Use target tracking scaling.

---

Organizers for a global event want to put daily reports online as static HTML pages. The pages are expected to generate millions of views from users around the world. The files are stored in an Amazon S3 bucket. A solutions architect has been asked to design an efficient and effective solution.

Which action should the solutions architect take to accomplish this?

##### Keywords: `s3 bucket`, `Amazon S3`, `effectient`, `effective`

- [ ] Generate presigned URLs for the files.
- [ ] Use cross-Region replication to all Regions.
- [ ] Use the geoproximity feature of Amazon Route 53.
- [x] Use Amazon CloudFront with the S3 bucket as its origin.

---

A company hosted an e-commerce website on an Auto Scaling group of EC2 instances behind an Application Load Balancer. The Solutions Architect noticed that the website is receiving a large number of illegitimate external requests from multiple systems with IP addresses that constantly change. To resolve the performance issues, the Solutions Architect must implement a solution that would block the illegitimate requests with minimal impact on legitimate traffic.

Which of the following options fulfills this requirement?

##### Keywords: `DDoS attacks`, `WAF`

- [x] Create a rate-based rule in AWS WAF and associate the web ACL to an Application Load Balancer

> [!TIP]
> large number of illegitimate external requests from multiple systems with IP addresses that constantly change (DDoS attacks) -> AWS WAF

---

A company operates an ecommerce website on Amazon EC2 instances behind an Application Load Balancer (ALB) in an Auto Scaling group. The site is experiencing performance issues related to a high request rate from illegitimate external systems with changing IP addresses. The security team is worried about potential DDoS attacks against the website. The company must block the illegitimate incoming requests in a way that has a minimal impact on legitimate users.

What should a solutions architect recommend? 

##### Keywords: `DDoS attacks`, `AWS WAF`, `rate-based rules`

- [ ] Deploy Amazon Inspector and associate it with the ALB.
- [x] Deploy AWS WAF, associate it with the ALB, and configure a rate-limiting rule.
- [ ] Deploy rules to the network ACLs associated with the ALB to block the incoming traffic.
- [ ] Deploy Amazon GuardDuty and enable rate-limiting protection when configuring GuardDuty.

> [!TIP]
> About the possibility of DDoS assaults on the website -> AWS WAF
> DDOS works by sending high amounts of SYN packets to the IP node. limiting it to a few requests will keep more connections available. networking question.

> [!NOTE]
> - Rate limit: For a rate-based rule, enter the maximum number of requests to allow in any five-minute period from an IP address that matches the rule's conditions. The rate limit must be at least 100.
> - You can specify a rate limit alone, or a rate limit and conditions. If you specify only a rate limit, AWS WAF places the limit on all IP addresses. If you specify a rate limit and conditions, AWS WAF places the limit on IP addresses that match the conditions.
> - When an IP address reaches the rate limit threshold, AWS WAF applies the assigned action (block or count) as quickly as possible, usually within 30 seconds.
> - Once the action is in place, if five minutes pass with no requests from the IP address, AWS WAF resets the counter to zero.

---

A company is moving its legacy workload to the AWS Cloud. The workload files will be shared,
appended, and frequently accessed through Amazon EC2 instances when they are first created The
files will be accessed occasionally as they age

What should a solutions architect recommend?

##### Keywords: `Legacy Workload`, `Amazon EC2`, `frequently accessed`

- [ ] Store the data using Amazon EC2 instances with attached Amazon Elastic Block Store (Amazon EBS) data volumes
- [ ] Store the data using AWS Storage Gateway volume gateway and export rarely accessed data to Amazon S3 storage 
- [x] Store the data using Amazon Elastic File System (Amazon EFS) with lifecycle management enabled for rarely accessed data
- [ ] Store the data using Amazon S3 with an S3 lifecycle policy enabled to move data to S3 Standard-Infrequent Access (S3 Standard-IA)

> [!TIP]
> The workload files are shared, which makes **Amazon Elastic File System (EFS)** the appropriate choice. 
> 
> EFS supports lifecycle management, allowing rarely accessed data to be moved to **cheaper storage based on the number of days the file has been idle**. It is important to ask questions in the appropriate context to ensure a focused discussion.

---

A company collects temperature, humidity, and atmospheric pressure data in cities across multiple continents. The average volume of data collected per site each day is 500 GB. Each site has a high-speed internet connection. The company's weather forecasting applications are based in a single Region and analyze the data daily.

What is the FASTEST way to aggregate data from all of these global sites?

##### Keywords: `High-speed internet connection`, `FASTEST`, `edge location`

- [x] Enable Amazon S3 Transfer Acceleration on the destination bucket. Use multipart uploads to directly upload site data to the destination bucket.
- [ ] Upload site data to an Amazon S3 bucket in the closest AWS Region. Use S3 cross-Region replication to copy objects to the destination bucket.
- [ ] Schedule AWS Snowball jobs daily to transfer data to the closest AWS Region. Use S3 cross-Region replication to copy objects to the destination bucket.
- [ ] Upload the data to an Amazon EC2 instance in the closest Region. Store the data in an Amazon Elastic Block Store (Amazon EBS) volume. Once a day take an EBS snapshot and copy it to the centralized Region. Restore the EBS volume in the centralized Region and run an analysis on the data daily.

> [!TIP]
> The question asks for the **FASTEST** way to aggregate data from all of these global sites. -> Transfer Acceleration
>
> - Transfer Acceleration will first get the data into the nearest AWS data center OR nearest Edge Locations.





> [!NOTE]
> Upload site data to an Amazon S3 bucket in the closest AWS Region and use S3 cross-Region replication to copy objects to the destination bucket. -> doesn't use edge location as the first destination. 
> 
> ### Why use Transfer Acceleration?
> 
> You might want to use Transfer Acceleration on a bucket for various reasons:
> 
> - Your customers upload to a **centralized bucket** from all over the world.
> - You transfer **gigabytes to terabytes** of data on a regular basis across continents.
> - You can't use **all of your available bandwidth** over the internet when uploading to Amazon S3.
> 
> Reference: [Configuring fast, secure file transfers using Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html)

---

A solutions architect is designing a solution where users will be directed to a backup static error page if the primary website is unavailable. The primary website's DNS records are hosted in Amazon Route 53 where their domain is pointing to an Application Load Balancer (ALB).

Which configuration should the solutions architect use to meet the company's needs while minimizing changes and infrastructure overhead?

##### Keywords: `Route 53`, `Backup Static Error Page`, `ALB`, `cost-effective`

- [ ] Point a Route 53 alias record to an Amazon CloudFront distribution with the ALB as one of its origins. Then, create custom error pages for the distribution.
- [x] Set up a Route 53 active-passive failover configuration. Direct traffic to a static error page hosted within an Amazon S3 bucket when Route 53 health checks determine that the ALB endpoint is unhealthy.
- [ ] Update the Route 53 record to use a latency-based routing policy. Add the backup static error page hosted within an Amazon S3 bucket to the record so the traffic is sent to the most responsive endpoints.
- [ ] Set up a Route 53 active-active configuration with the ALB and an Amazon EC2 instance hosting a static error page as endpoints. Route 53 will only send requests to the instance if the health checks fail for the ALB.

<!-- Active-passive failover -
Use an active-passive failover configuration when you want a primary resource or group of resources to be available the majority of the time and you want a secondary resource or group of resources to be on standby in case all the primary resources become unavailable. When responding to queries, Route 53 includes only the healthy primary resources. If all the primary resources are unhealthy, Route 53 begins to include only the healthy secondary resources in response to DNS queries.
To create an active-passive failover configuration with one primary record and one secondary record, you just create the records and specify Failover for the routing policy. When the primary resource is healthy, Route 53 responds to DNS queries using the primary record. When the primary resource is unhealthy, Route
53 responds to DNS queries using the secondary record.
How Amazon Route 53 averts cascading failures
As a first defense against cascading failures, each request routing algorithm (such as weighted and failover) has a mode of last resort. In this special mode, when all records are considered unhealthy, the Route 53 algorithm reverts to considering all records healthy.
For example, if all instances of an application, on several hosts, are rejecting health check requests, Route 53 DNS servers will choose an answer anyway and return it rather than returning no DNS answer or returning an NXDOMAIN (non-existent domain) response. An application can respond to users but still fail health checks, so this provides some protection against misconfiguration.
Similarly, if an application is overloaded, and one out of three endpoints fails its health checks, so that it's excluded from Route 53 DNS responses, Route 53 distributes responses between the two remaining endpoints. If the remaining endpoints are unable to handle the additional load and they fail, Route 53 reverts to distributing requests to all three endpoints.
Reference:
https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-types.html https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-problems.html -->


---

## NO ANSWER YET

A popular social media website uses a CloudFront web distribution to serve their static contents to their millions of users around the globe. They are receiving a number of complaints recently that their users take a lot of time to log into their website. There are also occasions when their users are getting HTTP 504 errors. You are instructed by your manager to significantly reduce the user's login time to further optimize the system. 

Which of the following options should you use together to set up a cost-effective solution that can improve your application's performance? (Select TWO.) 

- [ ] Configure your origin to add a Cache-Control max-age directive to your objects, and specify the longest practical value for max-age to increase the cache hit ratio of your CloudFront distribution. 
- [ ] Use multiple and geographically disperse VPCs to various AWS regions then create a transit VPC to connect all of your resources. In order to handle the requests faster, set up Lambda functions in each region using the AWS Serverless Application Model (SAM) service. 
- [ ] Customize the content that the CloudFront web distribution delivers to your users using Lambda@Edge, which allows your Lambda functions to execute the authentication process in AWS locations closer to the users. 
- [ ] Set up an origin failover by creating an origin group with two origins. Specify one as the primary origin and the other as the second origin which CloudFront automatically switches to when the primary origin returns specific HTTP status code failure responses.
- [ ] Deploy your application to multiple AWS regions to accommodate your users around the world. Set up a Route 53 record with latency routing policy to route incoming traffic to the region that provides the best latency to the user.

---

An online medical system hosted in AWS stores sensitive Personally Identifiable Information (PII) of the users in an Amazon S3 bucket. 

Both the master keys and the unencrypted data should never be sent to AWS to comply with the strict compliance and regulatory requirements of the company.

Which S3 encryption technique should the Architect use?

- [ ] Use S3 client-side encryption with a KMS-managed customer master key.
- [x] Use S3 client-side encryption with a client-side master key.
- [ ] Use S3 server-side encryption with a KMS-managed managed key.
- [ ] Use S3 server-side encryption with customer-provided key.


---

A Solutions Architect is hosting a website in an Amazon S3 bucket named tutorialsdojo. 
The users load the website using the following URL: `http://tutorialsdojo.s3-website-us-east-1.amazonaws.com` and there is a new requirement to add a JavaScript on the webpages in order to make authenticated HTTP  GET requests against the same bucket by using the Amazon S3 API endpoint (tutorialsdojo.s3.amazonaws.com). 

Upon testing, you noticed that the web browser blocks JavaScript from allowing those requests.

Which of the following options is the MOST suitable solution that you should implement for this scenario?

- [ ] Enable cross-account access.
- [ ] Enable Cross-Zone Load Balancing.
- [x] Enable Cross-origin resource sharing (CORS) configuration in the bucket.
- [ ] Enable Cross-Region Replication (CRR).

https://tutorialsdojo.com/aws-certified-solutions-architect-associate-saa-c03-sample-exam-questions/


---

A company is in the process of migrating their applications to AWS. One of their systems requires a database that can scale globally and handle frequent scheme changes. The application should not have any downtime or performance issue whenwever there is a scheme change in the database.

It should also provide a low latency response to high-traffic queries, Which is the most sutiable database solution to use to achieve this requirement?

- [ ] An Amazon RDS instane in Multi-AZ Deployment Configuration
- [x] Amazon DynamoDB
- [ ] An Amazon Aurora database with Read Replicas
- [ ] Redshift


---

A tech company has a CRM application hosted on an Auto Scaling group of On-Demand EC2 instances with different instance types and sizes. The application is extensively used during office hours from 9 in the morning to 5 in the afternoon. Their users are complaining that the performance of the application is slow during the start of the day but then works normally after a couple of hours.

Which of the following is the MOST operationally efficient solution to implement to ensure the application works properly at the beginning of the day?

- [ ] Configure a Dynamic scaling policy for the Auto Scaling group to launch new instances based on the CPU utilization.
- [ ] Configure a Dynamic scaling policy for the Auto Scaling group to launch new instances based on the Memory utilization.
- [ ] Configure a Scheduled scaling policy for the Auto Scaling group to launch new instances before the start of the day.
- [ ] Configure a Predictive scaling policy for the Auto Scaling group to automatically adjust the number of Amazon EC2 instances


---

A financial application is composed of an Auto Scaling group of EC2 instances, an Application Load Balancer, and a MySQL RDS instance in a Multi-AZ Deployments configuration. To protect the confidential data of your customers, you have to ensure that your RDS database can only be accessed using the profile credentials specific to your EC2 instances via an authentication token.

As the Solutions Architect of the company, which of the following should you do to meet the above requirement?

- [ ] Enable the IAM DB Authentication.
- [ ] Configure SSL in your application to encrypt the database connection to RDS.
- [ ] Create an IAM Role and assign it to your EC2 instances which will grant exclusive access to your RDS instance.
- [ ] Use a combination of IAM and STS to restrict access to your RDS instance via a temporary token.


---


A company hosted a web application in an Auto Scaling group of EC2 instances. The IT manager is concerned about the over-provisioning of the resources that can cause higher operating costs. A Solutions Architect has been instructed to create a cost-effective solution without affecting the performance of the application.

Which dynamic scaling policy should be used to satisfy this requirement?

- [ ] Use simple scaling.
- [ ] Use scheduled scaling.
- [ ] Use suspend and resume scaling.
- [ ] Use target tracking scaling.


---

A company is designing a banking portal that uses Amazon ElastiCache for Redis as its distributed session management component. Since the other Cloud Engineers in your department have access to your ElastiCache cluster, you have to secure the session data in the portal by requiring them to enter a password before they are granted permission to execute Redis commands.

As the Solutions Architect, which of the following should you do to meet the above requirement?

- [ ] Set up an IAM Policy and MFA which requires the Cloud Engineers to enter their IAM credentials and token before they can access the ElastiCache cluster.
- [ ] Set up a Redis replication group and enable the AtRestEncryptionEnabledparameter.
- [ ] Authenticate the users using Redis AUTH by creating a new Redis Cluster with both the --transit-encryption-enabled and --auth-tokenparameters enabled.
- [ ] Enable the in-transit encryption for Redis replication groups.


---

A company runs an online payments application in an Auto Scaling group of Amazon EC2 instances in multiple Availability Zones. The EC2 instances are all launched in private subnets. An internet-facing Application Load Balancer (ALB) has been provisioned and points to the existing EC2 instances as the target group. The team noticed that the internet traffic was not reaching the Amazon EC2 instances.

What is the MOST operationally efficient solution that meets these requirements?

- [ ] Set up a NAT gateway in a public subnet to allow incoming Internet traffic. Use a Gateway Load Balancer instead of an Application Load Balancer.
- [ ] Move the existing Amazon EC2 instances that are running from the private subnets to public subnets. Allow outbound traffic to 0.0.0.0/0 in the security groups of the EC2 instances.
- [ ] Add a rule to allow outbound traffic to 0.0.0.0/0 Fin the security groups of the EC2 instances. Update the route tables of the existing subnets to send all 0.0.0.0/0 traffic through the internet gateway route.
- [ ] Launch public subnets in each Availability Zone and associate them with the Application Load Balancer. Modify the route tables for the public subnets with a route to the private subnets of the EC2 instances.


---

The company that you are working for has a highly available architecture consisting of an elastic load balancer and several EC2 instances configured with auto-scaling in three Availability Zones. You want to monitor your EC2 instances based on a particular metric, which is not readily available in CloudWatch.

Which of the following is a custom metric in CloudWatch which you have to manually set up?

- [ ] Memory Utilization of an EC2 instance
- [ ] CPU Utilization of an EC2 instance
- [ ] Disk Reads activity of an EC2 instance
- [ ] Network packets out of an EC2 instance


---

A software development company is using serverless computing with AWS Lambda to build and run applications without having to set up or manage servers. They have a Lambda function that connects to a MongoDB Atlas, which is a popular Database as a Service (DBaaS) platform and also uses a third party API to fetch certain data for their application. One of the developers was instructed to create the environment variables for the MongoDB database hostname, username, and password as well as the API credentials that will be used by the Lambda function for DEV, SIT, UAT, and PROD environments.

Considering that the Lambda function is storing sensitive database and API credentials, how can this information be secured to prevent other developers in the team, or anyone, from seeing these credentials in plain text? Select the best option that provides maximum security.

- [ ] There is no need to do anything because, by default, AWS Lambda already encrypts the environment variables using the AWS Key Management Service.
- [ ] Enable SSL encryption that leverages on AWS CloudHSM to store and encrypt the sensitive information.
- [ ] AWS Lambda does not provide encryption for the environment variables. Deploy your code to an EC2 instance instead.
- [ ] Create a new KMS key and use it to enable encryption helpers that leverage on AWS Key Management Service to store and encrypt the sensitive information.


---

There was an incident in your production environment where the user data stored in the S3 bucket has been accidentally deleted by one of the Junior DevOps Engineers. The issue was escalated to your manager and after a few days, you were instructed to improve the security and protection of your AWS resources.

What combination of the following options will protect the S3 objects in your bucket from both accidental deletion and overwriting? (Select TWO.)

- [ ] Enable Versioning
- [ ] Provide access to S3 data strictly through pre-signed URL only
- [ ] Disallow S3 Delete using an IAM bucket policy
- [ ] Enable Amazon S3 Intelligent-Tiering
- [ ] Enable Multi-Factor Authentication Delete


---

A government entity is conducting a population and housing census in the city. 

Each household information uploaded on their online portal is stored in encrypted files in Amazon S3. 

The government assigned its Solutions Architect to set compliance policies that verify sensitive data in a manner that meets their compliance standards. 

They should also be alerted if there are compromised files detected containing personally identifiable information (PII), protected health information (PHI) or intellectual properties (IP). 

Which of the following should the Architect implement to satisfy this requirement? 








































