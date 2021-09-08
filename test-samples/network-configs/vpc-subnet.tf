resource "aws_vpc" "test-vpc" {
  cidr_block = "10.0.0.0/16"
  tags {
    Name = "test-vpc"
  }
}

resource "aws_subnet" "test-subnet-a" {
  vpc_id            = "${aws_vpc.test-vpc.id}"
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-2a"
}