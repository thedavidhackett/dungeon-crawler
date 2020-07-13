import React from "react";
import { Button, Form, FormGroup, Label, Input } from "reactstrap";

export default function CharacterForm({ onSubmit, label }) {
  return (
    <Form onSubmit={onSubmit}>
      <FormGroup>
        <Label for="name">Name</Label>
        <Input type="text" name="name" id="name" required />
      </FormGroup>
      <FormGroup>
        <Label for="image">Image</Label>
        <Input type="select" id="image" name="image" required>
          <option value="red-circle.png">Red</option>
          <option value="blue-circle.png">Blue</option>
          <option value="green-cirle.png">Green</option>
          <option value="yellow-cirle.png">Yellow</option>
          <option value="purple-cirle.png">Purple</option>
          <option value="black-cirle.png">Black</option>
        </Input>
      </FormGroup>
      <Button>{label}</Button>
    </Form>
  );
}
