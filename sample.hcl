resource "aws_iam_role" "example" {
  name               = "example"
  assume_role_policy = "..."
}
resource "aws_iam_instance_profile" "example" {
  role = aws_iam_role.example.name
}
resource "aws_iam_role_policy" "example" {
  name = "example"
  role = aws_iam_role.example.name
}
resource "aws_instance" "example" {
  ami = "ami-a1b2c3d4"
  depends_on = [
    aws_iam_role_policy.example,
  ]
}

resource "aws_instance" "server" {
  count = 4 # create four similar EC2 instances
  ami   = "ami-a1b2c3d4"
  tags = {
    Name = "Server ${count.index}"
  }
  policy = jsonencode({
    "Statement" = [{
      # This policy allows software running on the EC2 instance to
      # access the S3 API.
      "Action" = "s3:*",
      "Effect" = "Allow",
    }],
  })
}

resource "aws_security_group" "test-sg" {
  vpc_id      = "${aws_vpc.test-vpc.id}"
  name        = "test-sg"
  description = "This security group is for Terraform Test"
  tags { Name = "test-sg" }
}
