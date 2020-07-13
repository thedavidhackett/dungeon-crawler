import React from "react";
import { Container, Button } from "reactstrap";

export default function Login() {
  return (
    <Container>
      <h4>Welcome to Dungeon Crawler</h4>
      <Button
        outline
        color="primary"
        href="http://localhost:8000/auth/google/login"
      >
        Sign in with Google
      </Button>
    </Container>
  );
}
