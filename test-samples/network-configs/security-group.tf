resource "aws_security_group" "test-sg" {
  vpc_id      = "${aws_vpc.test-vpc.id}"
  name        = "test-sg"
  description = "This security group is for Terraform Test"
  tags { Name = "test-sg" }
}