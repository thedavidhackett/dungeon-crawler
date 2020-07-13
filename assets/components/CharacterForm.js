import React from "react";
import {
  Button,
  TextField,
  InputLabel,
  Select,
  MenuItem,
} from "@material-ui/core";

export default function CharacterForm({ onSubmit, label }) {
  return (
    <form onSubmit={onSubmit}>
      <TextField
        name="name"
        id="name"
        label="Name"
        variant="outlined"
        required
      ></TextField>
      <InputLabel id="image-label">Image</InputLabel>
      <Select labelId="image-label" id="image" name="image" required>
        <MenuItem value={"red-circle.png"}>Red</MenuItem>
        <MenuItem value={"blue-circle.png"}>Blue</MenuItem>
        <MenuItem value={"green-cirle.png"}>Green</MenuItem>
        <MenuItem value={"yellow-cirle.png"}>Yellow</MenuItem>
        <MenuItem value={"purple-cirle.png"}>Purple</MenuItem>
        <MenuItem value={"black-cirle.png"}>Black</MenuItem>
      </Select>
      <Button type="submit" variant="contained" color="primary">
        {label}
      </Button>
    </form>
  );
}
