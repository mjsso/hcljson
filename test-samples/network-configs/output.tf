output "test-subnet-a-id" {
  value = "${aws_subnet.test-subnet-a.id}"
}

output "test-sg-id" {
  value = "${aws_security_group.test-sg.id}"
}