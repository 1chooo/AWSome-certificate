What are the three fundamentals drivers of cost with AWS
Compute: Charged per hour/second + varies by instance type

Storage: Charged typically per GB

Data Transfer: Outbound traffic is aggregated and charged, and typically charged per GB. 



You can save money by using Reserve Instances (RIs). What are the three payment models RIs?

- All Upfront Reserved Instance (or AURI)

- Partial Upfront Reserved Instance (or PURI)

- No Upfront Payments Reserved Instance (or NURI)

When you buy Reserved Instances, you receive a greater discount when you make a larger upfront payment.



True or Talse: Volume-based discounts exist for AWS storage solutions.

True!

With AWS, you can get volume-based discounts and realize important savings as your usage increases. For services like Amazon Simple Storage Service (Amazon S3), pricing is tiered, which means that you pay less per GB when you use more.



What are the five services in AWS that can be used free-of-charge?
- Amazon Virtual Private Cloud (Amazon VPC)

- AWS Identity and Access Management (IAM)

- AWS Elastic Beanstalk

- AWS CloudFormation

- Automatic Scaling

- AWS OpsWorks

NOTE: Though there is no charge for these services, there might be charges associated with other

AWS services used with these services



What is Total Cost of Ownership?
Total cost of ownership (TCO) is the financial estimate to help identify direct and indirect costs of a system. TCO includes the cost of a service, plus all the costs that are associated with owning the service.

AWS also offers a TCO calculator.


What is the AWS Simply Monthly Calculator?
The AWS Simple Monthly Calculator helps estimate a monthly AWS bill. You can use this tool to add, modify, and remove services, and it recalculates the estimated monthly charges automatically.


What is AWS Organizations?
AWS Organizations is a free account management service that enables you to consolidate multiple AWS accounts into an organization that you create and centrally manage. AWS Organizations include consolidated billing and account management capabilities that help you to better meet the budgetary, security, and compliance needs of your business


What is the major terminology/roles that make up an AWS Organization?
Root: The man or admin account. This user has full control over all child accounts

Organizational Unit (OU): a container for accounts within a root. An OU can also contain other OUs.

Service Control Policy: Rules that govern a given account. When you attach a policy to one of the nodes in the hierarchy, it flows down and it affects all the branches



What is the recommended four-step process of creating an organization?
Step #1: Create an organization

Step #2: Create organizational units (OUs)

Step #3: Create service control policies

Step #4: Test those restrictions



What are the four technical support plans offered by AWS?
- Basic support plan
- Developer support plan
- Business support plan
- Enterprise support plan


What major features come with the Basic Support Plan?

- 24/7 access to customer service, documentation, whitepapers and support forums.
- Access to six core Trusted Advisor checks.
- Access to Personal Health Dashboard.


What major features come with the Developer Support Plan?

- Want access to guidance and technical support.
- Are exploring how to quickly put AWS to work.
- Use AWS for non-production workloads or applications.


What major features come with the Business Support Plan?

- Run one or more applications in production environments.
- Have multiple services activated, or use key services extensively.
- Depend on their business solutions to be available, scalable, and secure.


What major features come with the Enterprise Support Plan?

- Focus on proactive management to increase efficiency and availability.
- Build and operate workloads that follow AWS best practices.
- Use AWS expertise to support launches and migrations.
- Use a Technical Account Manager (TAM), who provides technical expertise for the full range of AWS services and obtains a detailed understanding of your use case and technology architecture.


Sample Exam Question: 

Which AWS service provides infrastructure security optimization recommendations? 

Answer: AWS Trusted Advisor