resource "aws_route_table" "test-route-table" {
  vpc_id = "${aws_vpc.test-vpc.id}"
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.test-gateway.id}"
  }
  tags {
    Name = "test-route-table"
  }
}

resource "aws_route_table_association" "my-subnet-association" {
  subnet_id      = "${aws_subnet.test-subnet-a.id}"
  route_table_id = "${aws_route_table.test-route-table.id}"
}