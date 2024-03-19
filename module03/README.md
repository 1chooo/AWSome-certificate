# Module 3 - Adding a Storage Layer

### Q1
Amazon Simple Storage Service (Amazon S3) provide a good solution for which of the following use cases?

- [ ] A data warehouse for business intelligence
- [x] An internet accessible storage location for video files that an external website accesses
- [ ] Hourly storage of frequently accessed temporary files
- [ ] A cluster for traditional Apache Spark and Apache Hadoop installations to process big data


### Q2
A company is interested in using Amazon Simple Storage Service (Amazon S3) alone to host their website, instead of a traditional web server. Which types of content does Amazon 53 support for static web hosting? (Select THREE.)

- [x] HTML files and image files
- [x] Client-side scripts
- [ ] Server-side scripts
- [ ] Dynamic HTML files
- [x] Video and sound files


### Q3
Which scenarios represent a good use for Amazon Simple Storage Service (Amazon S3)? (Select TWO.)

- [ ] Housing the root volume of a live operating system
- [ ] Providing a mountable file system for Linux-based workloads
- [x] Backing up critical data
- [ ] Exposing a virtual tape library to on-premises backup systems
- [x] Storing computation and analytics data


### Q4
A company wants to use an S3 bucket to store sensitive data. Which actions can they take to protect their data? (Select TWO.)

- [ ] Uploading unencrypted files to Amazon S3 because Amazon S3 encrypts the files by default
- [x] Enabling server-side encryption on the 53 bucket before uploading sensitive data
- [ ] Enabling server-side encryption on the 53 bucket after uploading sensitive data
- [x] Using client-side encryption to protect data in transit
- [ ] Using Secure File Transfer Protocol (SFTP) to connect directly to Amazon 53


### Q5
A company must create a common place to store shared files. Which requirements does Amazon Simple Storage Service (Amazon 53) support? (Select TWO.)

- [x] Recover deleted files
- [x] Maintain different versions of files
- [ ] Lock a file so that only one person at a time can edit it
- [ ] Attach comments to files
- [ ] Compare file contents between files


### Q6
A customer service team accesses case data daily for up to 30 days. Cases can be reopened and require immediate access for 1 year after they are closed. Reopened cases require 2 days to process. Which solution meets the requirements and is the most cost-efficient?

- [ ] Store all case data in 53 Standard so that it is available whenever needed.
- [x] Store case data in S3 Standard. Use a lifecycle policy to move the data into 53
- [ ] Standard-Infrequent Access (53 Standard-IA) after 30 days.
- [ ] Store case data in S3 Standard. Use a lifecycle policy to move the data into Amazon S3 Glacier after 30 days.
- [ ] Store case data in 53 Intelligent-Tiering to automatically move data between tiers based on access frequency.


### Q7
Which Amazon Simple Storage Service (Amazon S3) unaccelerated data transfers have an associated cost? (Select TWO.)

- [ ] IN from the internet
- [x] OUT to the internet
- [x] OUT to other AWS Regions
- [ ] OUT to other AWS services in the same AWS Region
- [ ] OUT to Amazon CloudFront


### Q8
A company is migrating 100 terabytes (TB) of data from their on-premises data center to Amazon Simple Storage Service (Amazon S3). The company connects to Amazon Web Services (AWS) by using a single 155 megabits per second (Mbps) internet connection. Which data transfer option is the fastest and most cost-effective?

- [ ] AWS Management Console
- [ ] Amazon S3 multipart uploads
- [x] AWS Snowball
- [ ] AWS Snowmobile


### Q9
A video producer must regularly transfer several video files to Amazon Simple Storage Service (Amazon 53). The files range from 100-700 MB. The internet connection has been unreliable, causing some uploads to fail. Which solution provides the fastest, most reliable, and most cost-effective way to transfer these files to Amazon S3?

- [ ] AWS Snowmobile
- [ ] AWS Management Console
- [ ] AWS Snowball
- [x] Amazon S3 multipart uploads


### Q10
Which qualities vary by AWS Region? (Select TWO.)

- [x] Cost-effectiveness of workloads
- [ ] Data privacy
- [ ] High availability of workloads
- [x] Service and feature availability
- [ ] Capacity for more customers
