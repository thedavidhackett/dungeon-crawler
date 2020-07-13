import React, { useState } from "react";
import { Redirect } from "react-router-dom";
import {
  Button,
  Form,
  FormGroup,
  Label,
  Input,
  Container,
  Row,
  Col,
} from "reactstrap";

export default function CreateDungeon() {
  const [id, setId] = useState(null);
  const classes = useStyles();
  const handleSubmit = (e) => {
    e.preventDefault();
    fetch("/api/create-dungeon", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        name: e.target.name.value,
        height: e.target.height.value,
        width: e.target.width.value,
      }),
    }).then((response) => {
      if (response.status == 200) {
        response.json().then((data) => {
          setId(data.InsertedID);
        });
      } else {
        console.log("failure");
      }
    });
  };

  return id ? (
    <Redirect to={"/dungeon/" + id} />
  ) : (
    <Container>
      <h4>Create Dungeon</h4>
      <Form onSubmit={handleSubmit}>
        <Row>
          <Col xs={12}>
            <FormGroup>
              <Label for="name">Name</Label>
              <Input type="text" name="name" id="name" />
            </FormGroup>
          </Col>
          <Col xs={12} sm={6}>
            <FormGroup>
              <Label for="height">Height</Label>
              <Input type="number" name="height" id="height" />
            </FormGroup>
          </Col>
          <Col xs={12} sm={6}>
            <FormGroup>
              <Label for="width">Width</Label>
              <Input type="number" name="width" id="width" />
            </FormGroup>
          </Col>
          <Col xs={12}>
            <Button>Create</Button>
          </Col>
        </Row>
      </Form>
    </Container>
  );
}
