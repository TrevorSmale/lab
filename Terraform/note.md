# Terraform

# Terraform File Breakdown

A typical Terraform configuration file can consist of the following components:

## 1. **Provider Block**
Defines the provider (e.g., AWS, Azure, GCP) to interact with cloud services.

```hcl
provider "aws" {
  region  = "us-west-2"
  version = "~> 4.0"
}
```

- provider: Specifies the cloud provider.

- region: Specifies the region to deploy resources.

- version: Specifies the version of the provider.

## 2. Resource Block

Defines the resources to be created.

```hcl

resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"

  tags = {
    Name = "ExampleInstance"
  }
}

```

- resource: Specifies the resource type and a name for referencing it (e.g., aws_instance.example).

- ami: The Amazon Machine Image (AMI) ID.

- instance_type: Specifies the instance size (e.g., t2.micro).

- tags: Optional metadata for the resource.

## 3. Variable Block

Defines input variables for parameterizing the configuration.

```hcl

variable "instance_type" {
  description = "Instance type for the EC2 instance"
  type        = string
  default     = "t2.micro"
}

```

- variable: Declares a variable.

- description: Describes the variable’s purpose.

- type: Specifies the variable type (string, number, list, etc.).

- default: Provides a default value.

### 4. Output Block

Defines outputs to display after Terraform applies the configuration.

```hcl

output "instance_ip" {
  description = "The public IP of the instance"
  value       = aws_instance.example.public_ip
}

```

- output: Declares an output.

- description: Describes the output’s purpose.

- value: Specifies the value to display.

## 5. Data Block

Fetches existing data from the cloud provider for use in the configuration.

```hcl

data "aws_ami" "example" {
  most_recent = true
  owners      = ["self"]

  filter {
    name   = "name"
    values = ["example-ami-*"]
  }
}

```

- data: Declares a data source.

- most_recent: Fetches the most recent AMI.

- filter: Filters the data based on specified criteria.

## 6. Terraform Block

Configures global settings for Terraform.

```hcl

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  required_version = ">= 1.0.0"
}

```

- required_providers: Specifies required providers and versions.

- required_version: Ensures the Terraform version is compatible.

## 7. Modules

Reusable blocks of Terraform configuration.

```hcl

module "vpc" {
  source  = "./modules/vpc"
  cidr    = "10.0.0.0/16"
  subnets = ["10.0.1.0/24", "10.0.2.0/24"]
}

```

- module: Refers to an external or local module. 

- source: Path or URL to the module’s location

- variables: Passes inputs to the module.

## 8. Lifecycle Block

Controls resource lifecycle behavior.

```hcl

resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"

  lifecycle {
    prevent_destroy = true
  }
}

```

- prevent_destroy: Prevents accidental resource destruction.

## 9. Provisioners

Runs scripts or commands on resources after they’re created.

```hcl

resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"

  provisioner "local-exec" {
    command = "echo Instance created!"
  }
}

```

- provisioner: Defines an action to take after resource creation.

- local-exec: Executes commands locally on the machine running Terraform.

## 10. Variables File

Separate variables into a .tfvars file for cleaner code.

variables.tf

```hcl

variable "region" {
  description = "AWS region"
  type        = string
}

```

terraform.tfvars

```hcl

region = "us-west-2"

```

## 11. Backend Configuration

Specifies where Terraform stores the state file.


```hcl

terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "path/to/state"
    region         = "us-west-2"
    encrypt        = true
  }
}

```

- backend: Configures the remote state storage.

- bucket: The S3 bucket to store the state file.

- key: The key (path) within the bucket.

- region: Region for the bucket.



