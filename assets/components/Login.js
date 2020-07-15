import React from "react";
import { Container, Row, Col, Button } from "reactstrap";

export default function Login() {
  return (
    <Container>
      <Row>
        <Col xs={12} md={{ size: 4, offset: 4 }}>
          <div className="login">
            <img className="login-image" src="/images/dungeon-outline.png" />
            <h4>Welcome to Dungeon Crawler</h4>
            <Button
              outline
              color="primary"
              href="http://localhost:8000/auth/google/login"
            >
              <img src="/images/google_icon.svg" />
              Sign in with Google
            </Button>
          </div>
        </Col>
      </Row>
    </Container>
  );
}
