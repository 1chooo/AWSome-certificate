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


