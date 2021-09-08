resource "tls_private_key" "example" {
  algorithm = "RSA"
  rsa_bits  = 4096

  provisioner "local-exec" {
    command = "echo '${self.private_key_pem}' > ./aws-key.pem"
  }
}

resource "aws_key_pair" "terraform-key" {
  key_name   = "aws-key"
  public_key = "${tls_private_key.example.public_key_openssh}"
}