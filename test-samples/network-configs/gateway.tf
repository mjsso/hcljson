resource "aws_internet_gateway" "test-gateway" {
  vpc_id = "${aws_vpc.test-vpc.id}"
  tags {
    Name = "test-gateway"
  }
}